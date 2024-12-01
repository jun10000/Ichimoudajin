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
	IsHit   bool
	Offset  Vector
	ROffset Vector
	Normal  Vector
}

func NewTraceResult_NoHit(offset Vector) TraceResult {
	return TraceResult{
		IsHit:   false,
		Offset:  offset,
		ROffset: ZeroVector(),
		Normal:  ZeroVector(),
	}
}

func NewTraceResult_Hit(offset Vector, roffset Vector, normal Vector) TraceResult {
	return TraceResult{
		IsHit:   true,
		Offset:  offset,
		ROffset: roffset,
		Normal:  normal.Normalize(),
	}
}

func (l *Level) traceCircle(circle CircleF, offset Vector, except Collider) TraceResult {
	tc := math.Ceil(offset.Length())
	tu := offset.DivF(tc)

	for i := 0.0; i <= tc; i++ {
		to := NewCircleF(circle.Origin.Add(tu.MulF(i)), circle.Radius)
		for _, c := range l.Colliders {
			if c == except {
				continue
			}

			n := to.Intersect(c.GetBounds())
			if !n.IsZero() {
				tro := tu.MulF(i - 2)
				return NewTraceResult_Hit(tro, offset.Sub(tro), n)
			}
		}
	}

	return NewTraceResult_NoHit(offset)
}

func (l *Level) Trace(target any, offset Vector, except Collider) TraceResult {
	switch v := target.(type) {
	case CircleF:
		return l.traceCircle(v, offset, except)
	default:
		log.Println("Detected not supported trace target type")
		return NewTraceResult_NoHit(ZeroVector())
	}
}
