package utility

import (
	"log"
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
	Offset         Vector
	Normal         Vector
	Distance       float64
	RDistance      float64
	DistanceRatio  float64
	RDistanceRatio float64
}

func NewTraceResult_NoHit(fulldistance float64) TraceResult {
	return TraceResult{
		IsHit:          false,
		Offset:         ZeroVector(),
		Normal:         ZeroVector(),
		Distance:       fulldistance,
		RDistance:      0,
		DistanceRatio:  1,
		RDistanceRatio: 0,
	}
}

func NewTraceResult_Hit(offset Vector, normal Vector, distance float64, fulldistance float64) TraceResult {
	if normal.IsZero() || distance < 0 || fulldistance <= 0 || distance > fulldistance {
		return NewTraceResult_NoHit(0)
	}

	rdist := fulldistance - distance
	return TraceResult{
		IsHit:          true,
		Offset:         offset,
		Normal:         normal.Normalize(),
		Distance:       distance,
		RDistance:      rdist,
		DistanceRatio:  distance / fulldistance,
		RDistanceRatio: rdist / fulldistance,
	}
}

func (l *Level) traceCircle(circle CircleF, offset Vector, except Collider) TraceResult {
	tc := math.Ceil(offset.Length())
	tu := offset.DivF(tc)

	for i := 1.0; i <= tc; i++ {
		nc := NewCircleF(circle.Origin.Add(tu.MulF(i)), circle.Radius)
		for _, c := range l.Colliders {
			if c == except {
				continue
			}

			normal := nc.Intersect(c.GetBounds())
			if !normal.IsZero() {
				vecdiff := tu.MulF(i - 1)
				return NewTraceResult_Hit(vecdiff, normal, vecdiff.Length(), offset.Length())
			}
		}
	}

	return NewTraceResult_NoHit(offset.Length())
}

func (l *Level) Trace(target any, offset Vector, except Collider) TraceResult {
	switch v := target.(type) {
	case CircleF:
		return l.traceCircle(v, offset, except)
	default:
		log.Println("Detected not supported trace target type")
		return NewTraceResult_NoHit(0)
	}
}
