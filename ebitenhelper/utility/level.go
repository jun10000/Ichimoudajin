package utility

import (
	"errors"
	"io/fs"
	"log"
	"math"
	"math/rand/v2"
	"runtime"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	Name                string
	IsLooping           bool
	AIGridSize          Vector
	AILocationDeviation float64
	AIPathfinding       *AStar

	Colliders      Set[Collider]
	InputReceivers Set[InputReceiver]
	AITickers      Set[AITicker]
	Tickers        Set[Ticker]
	Drawers        []Drawer
	DebugDraws     []func(screen *ebiten.Image)
}

func NewLevel(name string) *Level {
	return &Level{
		Name:                name,
		IsLooping:           false,
		AIGridSize:          NewVector(32, 32),
		AILocationDeviation: 0.5,
		AIPathfinding:       NewAStar(),

		Colliders:      make(Set[Collider]),
		InputReceivers: make(Set[InputReceiver]),
		AITickers:      make(Set[AITicker]),
		Tickers:        make(Set[Ticker]),
		Drawers:        make([]Drawer, 0, InitialDrawerCap),
		DebugDraws:     make([]func(screen *ebiten.Image), 0, DebugInitialDrawsCap),
	}
}

func (l *Level) Add(actor any) {
	if a, ok := actor.(Drawer); ok {
		l.Drawers = append(l.Drawers, a)
	}
	if a, ok := actor.(InputReceiver); ok {
		l.InputReceivers.Add(a)
	}
	if a, ok := actor.(AITicker); ok {
		l.AITickers.Add(a)
	}
	if a, ok := actor.(Ticker); ok {
		l.Tickers.Add(a)
	}
	if a, ok := actor.(Collider); ok {
		l.Colliders.Add(a)
	}
}

func (l *Level) GetColliderBounds(excepts Set[Collider]) func(yield func(Bounder) bool) {
	return func(yield func(Bounder) bool) {
		for c := range l.Colliders.SubRange(excepts) {
			for b := range c.GetColliderBounds() {
				if !yield(b) {
					return
				}
			}
		}
	}
}

func (l *Level) Intersect(target Bounder, excepts Set[Collider]) (result bool, normal Vector) {
	for b := range l.GetColliderBounds(excepts) {
		r, n := Intersect(target, b)
		if r {
			return true, n
		}
	}

	return false, ZeroVector()
}

func (l *Level) Trace(target Bounder, offset Vector, excepts Set[Collider]) (rOffset Vector, rNormal Vector, rIsHit bool) {
	ol, on := offset.Decompose()
	var bo Bounder

	for i := 0; i <= int(math.Trunc(ol)+1); i++ {
		v := on.MulF(float64(i))
		bo = target.Offset(v.X, v.Y, bo)
		r, n := l.Intersect(bo, excepts)
		if r {
			DrawDebugTraceDistance(target, i)
			if i <= TraceSafeDistance {
				return ZeroVector(), n, true
			} else {
				o := on.MulF(float64(i - 1))
				return o, n, true
			}
		}
	}

	return offset, ZeroVector(), false
}

func (l *Level) AIMove(self Mover, target Collider) {
	srl := self.GetMainColliderBounds().BoundingBox().CenterLocation()
	trl := target.GetMainColliderBounds().BoundingBox().CenterLocation()
	spl := l.RealToPFLocation(srl)
	tpl := l.RealToPFLocation(trl)

	if res, ok := l.AIPathfinding.GetResult(spl, tpl); ok {
		switch c := len(res); {
		case c > 2:
			trl1 := l.PFToRealLocation(res[1], true, l.AILocationDeviation)
			trl2 := l.PFToRealLocation(res[2], true, l.AILocationDeviation)
			trlave := trl1.Add(trl2.Sub(trl1).DivF(2))
			self.AddInput(trlave.Sub(srl), 1)
		case c == 2:
			trl1 := l.PFToRealLocation(res[1], true, l.AILocationDeviation)
			self.AddInput(trl1.Sub(srl), 1)
		case c == 1:
			self.AddInput(trl.Sub(srl), 1)
		}

		DrawDebugAIPath(res)
	}
}

func (l *Level) AIIsPFLocationValid(location Point) bool {
	loc := l.PFToRealLocation(location, false, 0)
	s := GetScreenSize()
	if loc.X < 0 || loc.Y < 0 || loc.X >= float64(s.X) || loc.Y >= float64(s.Y) {
		return false
	}

	b := NewRectangleF(
		loc.X+AIValidOffset,
		loc.Y+AIValidOffset,
		loc.X+l.AIGridSize.X-AIValidOffset,
		loc.Y+l.AIGridSize.Y-AIValidOffset)

	excepts := make(Set[Collider])
	for t := range l.AITickers {
		if c, ok := t.(Collider); ok {
			excepts.Add(c)
		}
	}
	for t := range l.InputReceivers {
		if c, ok := t.(Collider); ok {
			excepts.Add(c)
		}
	}

	r, _ := l.Intersect(b, excepts)
	return !r
}

func (l *Level) RealToPFLocation(realLocation Vector) Point {
	return realLocation.Div(l.AIGridSize).Trunc()
}

func (l *Level) PFToRealLocation(pfLocation Point, isCenter bool, deviation float64) Vector {
	rr := l.AIGridSize.MulF((rand.Float64() - 0.5) * deviation)
	r := pfLocation.ToVector().Mul(l.AIGridSize).Add(rr)
	if isCenter {
		r = r.Add(l.AIGridSize.DivF(2))
	}
	return r
}

func (l *Level) GetPFCacheFileName() string {
	return l.Name + ".pfd"
}

func (l *Level) LoadPFCache() error {
	return l.AIPathfinding.LoadCache(l.GetPFCacheFileName())
}

func (l *Level) LoadOrBuildPFCache() error {
	err := l.LoadPFCache()
	if errors.Is(err, fs.ErrNotExist) {
		return l.BuildPFCache()
	}

	return err
}

func (l *Level) BuildPFCache() error {
	pf := l.AIPathfinding
	sz := l.RealToPFLocation(GetScreenSize().SubXY(1, 1).ToVector()).AddXY(1, 1)
	sem := make(chan Empty, runtime.GOMAXPROCS(0)-1)
	wg := sync.WaitGroup{}
	stime := time.Now()

	defer close(sem)
	log.Println("Started building PF cache")
	for sx := 0; sx < sz.X; sx++ {
		for sy := 0; sy < sz.Y; sy++ {
			for gx := 0; gx < sz.X; gx++ {
				for gy := 0; gy < sz.Y; gy++ {
					start := NewPoint(sx, sy)
					goal := NewPoint(gx, gy)
					if _, ok := pf.GetCache(start, goal); !ok {
						sem <- Empty{}
						wg.Add(1)
						go func(start Point, goal Point) {
							defer wg.Done()
							defer func() { <-sem }()
							pf.GetResultForce(start, goal)
						}(start, goal)
					}
				}
			}
			log.Printf("Building PF cache: %.2f%%\n", float32(sx*sz.Y+sy+1)*100/float32(sz.X*sz.Y))
		}
	}

	wg.Wait()
	log.Printf("Completed building PF cache, %.1fs elapsed.\n", time.Since(stime).Seconds())
	return pf.SaveCache(l.GetPFCacheFileName())
}

func (l *Level) AddDebugDraw(event func(*ebiten.Image)) {
	l.DebugDraws = append(l.DebugDraws, event)
}

func (l *Level) ClearDebugDraw() {
	l.DebugDraws = l.DebugDraws[:0]
}
