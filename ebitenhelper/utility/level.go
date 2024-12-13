package utility

import (
	"log"
	"math"
)

type Level struct {
	IsLooping    bool
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

func (l *Level) GetColliderBounds(except Collider) []Bounder {
	ret := []Bounder{}
	for _, c := range l.Colliders {
		if c == except {
			continue
		}

		b := c.GetBounds()
		ret = append(ret, b)
		if l.IsLooping {
			ss := GetGameInstance().ScreenSize.ToVector()
			ret = append(ret,
				b.Offset(ss.Mul(NewVector(-1, -1))),
				b.Offset(ss.Mul(NewVector(0, -1))),
				b.Offset(ss.Mul(NewVector(1, -1))),
				b.Offset(ss.Mul(NewVector(-1, 0))),
				b.Offset(ss.Mul(NewVector(1, 0))),
				b.Offset(ss.Mul(NewVector(-1, 1))),
				b.Offset(ss.Mul(NewVector(0, 1))),
				b.Offset(ss.Mul(NewVector(1, 1))),
			)
		}
	}
	return ret
}

func (l *Level) traceCircle(circle CircleF, offset Vector, except Collider) TraceResult {
	cnt := math.Ceil(offset.Length())
	uni := offset.DivF(cnt)
	bs := l.GetColliderBounds(except)

	for i := 0.0; i <= cnt; i++ {
		obj := NewCircleF(circle.Origin.Add(uni.MulF(i)), circle.Radius)
		for _, b := range bs {
			tr := obj.Intersect(b)
			if !tr.IsZero() {
				res := uni.MulF(i - 2)
				return NewTraceResult_Hit(res, offset.Sub(res), tr)
			}
		}
	}

	return NewTraceResult_NoHit(offset)
}

func (l *Level) Trace(target Bounder, offset Vector, except Collider) TraceResult {
	switch v := target.(type) {
	case CircleF:
		return l.traceCircle(v, offset, except)
	default:
		log.Println("Detected not supported trace target type")
		return NewTraceResult_NoHit(ZeroVector())
	}
}
