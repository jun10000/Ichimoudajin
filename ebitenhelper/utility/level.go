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
	IsHit          bool
	Location       Vector
	Normal         Vector
	Distance       float64
	RDistance      float64
	DistanceRatio  float64
	RDistanceRatio float64
}

func NewTraceResult_NoHit(fulldistance float64) TraceResult {
	return TraceResult{
		IsHit:          false,
		Location:       ZeroVector(),
		Normal:         ZeroVector(),
		Distance:       fulldistance,
		RDistance:      0,
		DistanceRatio:  1,
		RDistanceRatio: 0,
	}
}

func NewTraceResult_Hit(location Vector, normal Vector, distance float64, fulldistance float64) TraceResult {
	if normal.IsZero() || distance < 0 || fulldistance <= 0 || distance > fulldistance {
		return NewTraceResult_NoHit(0)
	}

	rdist := fulldistance - distance
	return TraceResult{
		IsHit:          true,
		Location:       location,
		Normal:         normal.Normalize(),
		Distance:       distance,
		RDistance:      rdist,
		DistanceRatio:  distance / fulldistance,
		RDistanceRatio: rdist / fulldistance,
	}
}

func (l *Level) RectTrace(src Vector, dst Vector, size Vector, except Collider) TraceResult {
	tracevec := dst.Sub(src)
	tracecount := math.Ceil(tracevec.Length())
	tracevec_unit := tracevec.DivF(tracecount)

	for i := 1.0; i <= tracecount; i++ {
		rect := NewRectangleF(src.Add(tracevec_unit.MulF(i)), size)
		for _, c := range l.Colliders {
			if c == except {
				continue
			}

			normal := rect.Intersect(c.GetBounds())
			if !normal.IsZero() {
				vecdiff := tracevec_unit.MulF(i - 1)
				return NewTraceResult_Hit(src.Add(vecdiff), normal, vecdiff.Length(), tracevec.Length())
			}
		}
	}

	return NewTraceResult_NoHit(tracevec.Length())
}
