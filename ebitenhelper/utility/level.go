package utility

import (
	"math"
)

type Level struct {
	Drawers      []Drawer
	KeyReceivers []KeyReceiver
	Tickers      []Ticker
	Colliders    []Collider
}

func NewLevel() *Level {
	return &Level{}
}

func (l *Level) Add(actor any) {
	d, ok := actor.(Drawer)
	if ok {
		l.Drawers = append(l.Drawers, d)
	}

	r, ok := actor.(KeyReceiver)
	if ok {
		l.KeyReceivers = append(l.KeyReceivers, r)
	}

	t, ok := actor.(Ticker)
	if ok {
		l.Tickers = append(l.Tickers, t)
	}

	c, ok := actor.(Collider)
	if ok {
		l.Colliders = append(l.Colliders, c)
	}
}

func (l *Level) AddRange(actors []any) {
	for _, a := range actors {
		l.Add(a)
	}
}

type TraceResult struct {
	IsHit         bool
	Location      Vector
	Normal        Vector
	Distance      float64
	DistanceRatio float64
}

func NewTraceResult(distance float64) TraceResult {
	return TraceResult{
		IsHit:         false,
		Location:      ZeroVector(),
		Normal:        ZeroVector(),
		Distance:      distance,
		DistanceRatio: 1,
	}
}

func (l *Level) RectTrace(src Vector, dst Vector, size Vector, except Collider) TraceResult {
	tracevec := dst.Sub(src)
	tracecount := math.Ceil(tracevec.Length())
	tracevec_unit := tracevec.DivF(tracecount)
	result := NewTraceResult(tracevec.Length())

	for i := 1.0; i <= tracecount; i++ {
		rect := NewRectangleF(src.Add(tracevec_unit.MulF(i)), size)
		for _, c := range l.Colliders {
			if c == except {
				continue
			}

			normal := rect.Intersect(c.GetBounds())
			if !normal.IsZero() {
				vecdiff := tracevec_unit.MulF(i - 1)
				dist := vecdiff.Length()
				result.IsHit = true
				result.Location = src.Add(vecdiff)
				result.Normal = normal
				result.Distance = dist
				result.DistanceRatio = dist / tracevec.Length()
				return result
			}
		}
	}

	return result
}
