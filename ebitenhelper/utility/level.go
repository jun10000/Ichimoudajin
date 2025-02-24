package utility

import (
	"errors"
	"image/color"
	"io/fs"
	"log"
	"math"
	"math/rand/v2"
	"runtime"
	"slices"
	"sync"
	"time"
)

type Level struct {
	colliders []Collider

	Name                string
	IsLooping           bool
	Drawers             []Drawer
	InputReceivers      []InputReceiver
	AITickers           []AITicker
	Tickers             []Ticker
	AIGridSize          Vector
	AILocationDeviation float64
	AIPathfinding       *AStar
}

func NewLevel(name string) *Level {
	return &Level{
		Name:                name,
		AIGridSize:          NewVector(64, 64),
		AILocationDeviation: 0.5,
		AIPathfinding:       NewAStar(),
	}
}

func (l *Level) Add(actor any) {
	if a, ok := actor.(Drawer); ok {
		l.Drawers = append(l.Drawers, a)
	}
	if a, ok := actor.(InputReceiver); ok {
		l.InputReceivers = append(l.InputReceivers, a)
	}
	if a, ok := actor.(AITicker); ok {
		l.AITickers = append(l.AITickers, a)
	}
	if a, ok := actor.(Ticker); ok {
		l.Tickers = append(l.Tickers, a)
	}
	if a, ok := actor.(Collider); ok {
		l.colliders = append(l.colliders, a)
	}
}

func (l *Level) AddRange(actors []any) {
	for _, a := range actors {
		l.Add(a)
	}
}

type TraceResult struct {
	IsHit      bool
	IsFirstHit bool
	Offset     Vector
	ROffset    Vector
	Normal     Vector
}

func NewTraceResultNoHit(offset Vector) TraceResult {
	return TraceResult{
		IsHit:      false,
		IsFirstHit: false,
		Offset:     offset,
		ROffset:    ZeroVector(),
		Normal:     ZeroVector(),
	}
}

func NewTraceResultHit(offset Vector, roffset Vector, normal Vector, isFirstHit bool) TraceResult {
	return TraceResult{
		IsHit:      true,
		IsFirstHit: isFirstHit,
		Offset:     offset,
		ROffset:    roffset,
		Normal:     normal.Normalize(),
	}
}

func (l *Level) GetAllColliders(excepts []Collider) []Collider {
	r := make([]Collider, 0, len(l.colliders)-len(excepts))
	for _, c := range l.colliders {
		if slices.Contains(excepts, c) {
			continue
		}
		r = append(r, c)
	}

	return r
}

func (l *Level) GetAllColliderBounds(excepts []Collider) []Bounder {
	rCap := len(l.colliders) - len(excepts)
	if l.IsLooping {
		rCap *= 9
	}
	r := make([]Bounder, 0, rCap)
	s := GetGameInstance().ScreenSize.ToVector()

	for _, c := range l.GetAllColliders(excepts) {
		b := c.GetColliderBounds()
		r = append(r, b)
		if !l.IsLooping {
			continue
		}

		r = append(r,
			b.Offset(-s.X, -s.Y),
			b.Offset(0, -s.Y),
			b.Offset(s.X, -s.Y),
			b.Offset(-s.X, 0),
			b.Offset(s.X, 0),
			b.Offset(-s.X, s.Y),
			b.Offset(0, s.Y),
			b.Offset(s.X, s.Y))
	}

	return r
}

func (l *Level) Intersect(target Bounder, excepts []Collider) (result bool, normal Vector) {
	for _, b := range l.GetAllColliderBounds(excepts) {
		r, n := Intersect(target, b)
		if r {
			return true, n
		}
	}

	return false, ZeroVector()
}

func (l *Level) Trace(target Bounder, offset Vector, excepts []Collider) TraceResult {
	ol := offset.Length()
	on := offset.Normalize()

	for i := 0; i <= int(math.Trunc(ol)+1); i++ {
		v := on.MulF(float64(i))
		t := target.Offset(v.X, v.Y)
		r, n := l.Intersect(t, excepts)
		if r {
			if IsShowDebugTraceDistance {
				var dc color.RGBA
				switch i {
				case 0:
					dc = ColorRed
				case 1:
					dc = ColorYellow
				case 2:
					dc = ColorGreen
				case 3:
					dc = ColorBlue
				default:
					dc = ColorGray
				}

				switch dt := target.(type) {
				case CircleF:
					DrawDebugCircle(NewVector(dt.OrgX, dt.OrgY), dt.Radius, dc)
				default:
					db := target.BoundingBox()
					DrawDebugRectangle(db.Location(), db.Size(), dc)
				}
			}
			if i <= TraceSafeDistance {
				return NewTraceResultHit(ZeroVector(), offset, n, true)
			} else {
				o := on.MulF(float64(i - 1))
				ro := offset.Sub(o)
				return NewTraceResultHit(o, ro, n, false)
			}
		}
	}

	return NewTraceResultNoHit(offset)
}

func (l *Level) AIMove(self Mover, target Collider) {
	srl := self.GetColliderBounds().BoundingBox().CenterLocation()
	trl := target.GetColliderBounds().BoundingBox().CenterLocation()
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

		if IsShowDebugAIPath {
			for _, p := range res {
				DrawDebugRectangle(l.PFToRealLocation(p, false, 0), l.AIGridSize, ColorGreen)
			}
		}
	}
}

func (l *Level) AIIsPFLocationValid(location Point) bool {
	s := GetGameInstance().ScreenSize
	loc := l.PFToRealLocation(location, false, 0)
	if loc.X < 0 || loc.Y < 0 || loc.X >= float64(s.X) || loc.Y >= float64(s.Y) {
		return false
	}

	b := NewRectangleF(
		loc.X+AIValidOffset,
		loc.Y+AIValidOffset,
		loc.X+l.AIGridSize.X-AIValidOffset,
		loc.Y+l.AIGridSize.Y-AIValidOffset)

	excepts := make([]Collider, 0, len(l.AITickers)+len(l.InputReceivers))
	for _, t := range l.AITickers {
		if c, ok := t.(Collider); ok {
			excepts = append(excepts, c)
		}
	}
	for _, t := range l.InputReceivers {
		if c, ok := t.(Collider); ok {
			excepts = append(excepts, c)
		}
	}

	r, _ := l.Intersect(b, excepts)
	return !r
}

func (l *Level) RealToPFLocation(realLocation Vector) Point {
	return realLocation.Div(l.AIGridSize).Floor()
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
	sz := l.RealToPFLocation(GetGameInstance().ScreenSize.SubXY(1, 1).ToVector()).AddXY(1, 1)
	sem := make(chan struct{}, runtime.GOMAXPROCS(0)-1)
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
						sem <- struct{}{}
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
