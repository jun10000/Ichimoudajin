package utility

import (
	"math"
	"slices"
)

type Level struct {
	IsLooping      bool
	Drawers        []Drawer
	InputReceivers []InputReceiver
	AITickers      []AITicker
	Tickers        []Ticker
	Colliders      []Collider
}

func NewLevel() *Level {
	return &Level{}
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

func (l *Level) Trace(target Bounder, offset Vector, excepts []Collider, isDebug bool) TraceResult {
	ol := offset.Length()
	on := offset.Normalize()

	for i := 0; i <= int(math.Trunc(ol)+1); i++ {
		v := on.MulF(float64(i))
		t := target.Offset(v)
		r, n := l.Intersect(t, excepts)
		if r {
			if isDebug {
				dc := ColorGreen
				switch i {
				case 0:
					dc = ColorRed
				case 1:
					dc = ColorYellow
				}

				switch dt := target.(type) {
				case CircleF:
					DrawDebugCircle(dt.Origin, dt.Radius, dc)
				default:
					db := target.BoundingBox()
					DrawDebugRectangle(db.Location(), db.Size(), dc)
				}
			}
			if i == 0 {
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
