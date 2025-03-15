package utility

import (
	"errors"
	"io/fs"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	pfValidCache *Smap[Point, bool]

	Name          string
	IsLooping     bool
	AIGridSize    Vector
	AIPathfinding *AStar

	Colliders        *Smap[Collider, [9]Bounder]
	StaticColliders  *Smap[Collider, [9]Bounder]
	MovableColliders *Smap[MovableCollider, [9]Bounder]
	InputReceivers   []InputReceiver
	Players          []Player
	AITickers        []AITicker
	Tickers          []Ticker
	Drawers          [][]Drawer
	DebugDraws       []func(screen *ebiten.Image)
}

func NewLevel(name string, isLooping bool) *Level {
	return &Level{
		pfValidCache: NewSmap[Point, bool](),

		Name:          name,
		IsLooping:     isLooping,
		AIGridSize:    NewVector(32, 32),
		AIPathfinding: NewAStar(),

		Colliders:        NewSmap[Collider, [9]Bounder](),
		StaticColliders:  NewSmap[Collider, [9]Bounder](),
		MovableColliders: NewSmap[MovableCollider, [9]Bounder](),
		InputReceivers:   make([]InputReceiver, 0, InitialInputReceiverCap),
		Players:          make([]Player, 0, InitialInputReceiverCap),
		AITickers:        make([]AITicker, 0, InitialAITickerCap),
		Tickers:          make([]Ticker, 0, InitialTickerCap),
		Drawers:          make([][]Drawer, 0, ZOrderMax+1),
		DebugDraws:       make([]func(screen *ebiten.Image), 0, DebugInitialDrawsCap),
	}
}

func (l *Level) Add(actor any) {
	if a, ok := actor.(Collider); ok {
		l.Colliders.Store(a, a.GetColliderBounds())
		if m, ok := a.(MovableCollider); ok {
			l.MovableColliders.Store(m, m.GetColliderBounds())
		} else {
			l.StaticColliders.Store(a, a.GetColliderBounds())
		}
	}
	if a, ok := actor.(InputReceiver); ok {
		l.InputReceivers = append(l.InputReceivers, a)
	}
	if a, ok := actor.(Player); ok {
		l.Players = append(l.Players, a)
	}
	if a, ok := actor.(AITicker); ok {
		l.AITickers = append(l.AITickers, a)
	}
	if a, ok := actor.(Ticker); ok {
		l.Tickers = append(l.Tickers, a)
	}
	if a, ok := actor.(Drawer); ok {
		z := ZOrderDefault
		if az, ok := a.(ZHolder); ok {
			z = az.ZOrder()
		}

		for range z - len(l.Drawers) + 1 {
			l.Drawers = append(l.Drawers, make([]Drawer, 0, InitialDrawerCap))
		}

		l.Drawers[z] = append(l.Drawers[z], a)
	}

	if a, ok := actor.(Parenter); ok {
		for _, ac := range a.Children() {
			l.Add(ac)
		}
	}
}

func (l *Level) Remove(actor any) {
	if a, ok := actor.(Collider); ok {
		l.Colliders.Delete(a)
		if m, ok := a.(MovableCollider); ok {
			l.MovableColliders.Delete(m)
		} else {
			l.StaticColliders.Delete(a)
		}
	}
	if a, ok := actor.(InputReceiver); ok {
		l.InputReceivers = RemoveSliceItem(l.InputReceivers, a)
	}
	if a, ok := actor.(Player); ok {
		l.Players = RemoveSliceItem(l.Players, a)
	}
	if a, ok := actor.(AITicker); ok {
		l.AITickers = RemoveSliceItem(l.AITickers, a)
	}
	if a, ok := actor.(Ticker); ok {
		l.Tickers = RemoveSliceItem(l.Tickers, a)
	}
	if a, ok := actor.(Drawer); ok {
		z := ZOrderDefault
		if az, ok := a.(ZHolder); ok {
			z = az.ZOrder()
		}

		l.Drawers[z] = RemoveSliceItem(l.Drawers[z], a)
	}

	if a, ok := actor.(Parenter); ok {
		for _, ac := range a.Children() {
			l.Remove(ac)
		}
	}
}

func (l *Level) AIMove(self MovableCollider, target Collider) {
	srl := self.GetMainColliderBounds().CenterLocation()
	trl := target.GetMainColliderBounds().CenterLocation()
	spl := l.RealToPFLocation(srl)
	tpl := l.RealToPFLocation(trl)

	if res, ok := l.AIPathfinding.GetResult(spl, tpl); ok {
		switch c := len(res); {
		case c > 2:
			trl1 := l.PFToRealLocation(res[1], true)
			trl2 := l.PFToRealLocation(res[2], true)
			trlave := trl1.Add(trl2.Sub(trl1).DivF(2))
			self.AddInput(trlave.Sub(srl), 1)
		case c == 2:
			trl1 := l.PFToRealLocation(res[1], true)
			self.AddInput(trl1.Sub(srl), 1)
		case c == 1:
			self.AddInput(trl.Sub(srl), 1)
		}

		DrawDebugAIPath(res)
	}
}

func (l *Level) AIIsPFLocationValid(location Point) bool {
	if r, ok := l.pfValidCache.Load(location); ok {
		return r
	}

	loc := l.PFToRealLocation(location, false)
	s := GetScreenSize()
	if loc.X < 0 || loc.Y < 0 || loc.X >= float64(s.X) || loc.Y >= float64(s.Y) {
		return false
	}

	b := NewRectangleF(
		loc.X+AIValidOffset,
		loc.Y+AIValidOffset,
		loc.X+l.AIGridSize.X-AIValidOffset,
		loc.Y+l.AIGridSize.Y-AIValidOffset)
	r, _, _ := Intersect(l.StaticColliders, b, nil)

	l.pfValidCache.Store(location, !r)
	return !r
}

func (l *Level) RealToPFLocation(realLocation Vector) Point {
	return realLocation.Div(l.AIGridSize).Trunc()
}

func (l *Level) PFToRealLocation(pfLocation Point, isCenter bool) Vector {
	r := pfLocation.ToVector().Mul(l.AIGridSize)
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
	sz := l.RealToPFLocation(GetScreenSize().ToVector())
	sem := make(chan Empty, runtime.GOMAXPROCS(0)-1)
	wg := sync.WaitGroup{}
	stime := time.Now()

	defer close(sem)
	log.Println("Started building PF cache")
	for sx := range sz.X {
		for sy := range sz.Y {
			for gx := range sz.X {
				for gy := range sz.Y {
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
		}
		log.Printf("Building PF cache: %d%%\n", (sx+1)*100/sz.X)
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
