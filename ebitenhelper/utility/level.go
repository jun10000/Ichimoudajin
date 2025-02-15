package utility

import (
	"image/color"
	"math"
	"math/rand/v2"
	"slices"
)

type Level struct {
	IsLooping           bool
	Drawers             []Drawer
	InputReceivers      []InputReceiver
	AITickers           []AITicker
	Tickers             []Ticker
	Colliders           []Collider
	AIGridSize          Vector
	AILocationDeviation float64
	AIPathfinding       *AStar
}

func NewLevel() *Level {
	return &Level{
		AIGridSize:          NewVector(32, 32),
		AILocationDeviation: 0.5,
		AIPathfinding:       StartAStar(),
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
		l.Colliders = append(l.Colliders, a)
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

func (l *Level) GetAllBounds(excepts []Collider) []Bounder {
	r := []Bounder{}
	for _, c := range l.Colliders {
		if slices.Contains(excepts, c) {
			continue
		}

		b := c.GetColliderBounds()
		r = append(r, b)
		if l.IsLooping {
			s := GetGameInstance().ScreenSize.ToVector()
			r = append(r,
				b.Offset(s.Mul(NewVector(-1, -1))),
				b.Offset(s.Mul(NewVector(0, -1))),
				b.Offset(s.Mul(NewVector(1, -1))),
				b.Offset(s.Mul(NewVector(-1, 0))),
				b.Offset(s.Mul(NewVector(1, 0))),
				b.Offset(s.Mul(NewVector(-1, 1))),
				b.Offset(s.Mul(NewVector(0, 1))),
				b.Offset(s.Mul(NewVector(1, 1))),
			)
		}
	}
	return r
}

func (l *Level) Intersect(target Bounder, excepts []Collider) (result bool, normal Vector) {
	for _, b := range l.GetAllBounds(excepts) {
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
		t := target.Offset(v)
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
					DrawDebugCircle(dt.Origin, dt.Radius, dc)
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
	sl := self.GetColliderBounds().BoundingBox().CenterLocation()
	tl := target.GetColliderBounds().BoundingBox().CenterLocation()

	res, ok := l.AIPathfinding.Run(l.RealToPFLocation(sl), l.RealToPFLocation(tl))
	switch ok {
	case AStarResultReasonResponsed:
		switch c := len(res); {
		case c > 2:
			dl1 := l.PFToRealLocation(res[1], true, l.AILocationDeviation)
			dl2 := l.PFToRealLocation(res[2], true, l.AILocationDeviation)
			tl = dl1.Add(dl2.Sub(dl1).DivF(2))
			self.AddInput(tl.Sub(sl), 1)
		case c == 2:
			tl = l.PFToRealLocation(res[1], true, l.AILocationDeviation)
			self.AddInput(tl.Sub(sl), 1)
		case c == 1:
			self.AddInput(tl.Sub(sl), 1)
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

	b := NewRectangleF(loc.AddF(AIValidOffset), l.AIGridSize.SubF(AIValidOffset*2))

	var excepts []Collider
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
