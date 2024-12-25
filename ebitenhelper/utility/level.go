package utility

import (
	"math"
)

type Level struct {
	IsLooping      bool
	Drawers        []Drawer
	InputReceivers []InputReceiver
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
	IsHit   bool
	Offset  Vector
	ROffset Vector
	Normal  Vector
}

func NewTraceResultNoHit(offset Vector) TraceResult {
	return TraceResult{
		IsHit:   false,
		Offset:  offset,
		ROffset: ZeroVector(),
		Normal:  ZeroVector(),
	}
}

func NewTraceResultHit(offset Vector, roffset Vector, normal Vector) TraceResult {
	return TraceResult{
		IsHit:   true,
		Offset:  offset,
		ROffset: roffset,
		Normal:  normal.Normalize(),
	}
}

func (l *Level) GetActorBounds(except Collider) []Bounder {
	ret := []Bounder{}
	for _, c := range l.Colliders {
		if c == except {
			continue
		}

		b := c.GetColliderBounds()
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

func (l *Level) Trace(target Bounder, offset Vector, except Collider) TraceResult {
	cnt := math.Ceil(offset.Length())
	uni := offset.DivF(cnt)
	bs := l.GetActorBounds(except)

	for i := 0.0; i <= cnt; i++ {
		obj := target.Offset(uni.MulF(i))
		for _, b := range bs {
			tr := Intersect(obj, b)
			if !tr.IsZero() {
				res := uni.MulF(i - 2)
				return NewTraceResultHit(res, offset.Sub(res), tr)
			}
		}
	}

	return NewTraceResultNoHit(offset)
}
