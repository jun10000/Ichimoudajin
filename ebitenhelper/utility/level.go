package utility

import (
	"math"
)

type TraceResult struct {
	IsHit       bool
	HitLocation Vector
	HitNormal   Vector
}

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

func (l *Level) RectTrace(src Vector, dst Vector, size Vector, except Collider) TraceResult {
	vecdiff := dst.Sub(src)
	tracecount := math.Ceil(vecdiff.Length())
	tracediff := vecdiff.DivF(tracecount)

	for i := 1.0; i <= tracecount; i++ {
		tracerect := NewRectangleF(src.Add(tracediff.MulF(i)), size)
		for _, c := range l.Colliders {
			if c == except {
				continue
			}

			n := tracerect.Intersect(c.GetBounds())
			if !n.IsZero() {
				return TraceResult{
					IsHit:       true,
					HitLocation: src.Add(tracediff.MulF(i - 1)),
					HitNormal:   n,
				}
			}
		}
	}

	return TraceResult{}
}
