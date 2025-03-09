package utility

import "math"

type TraceResult struct {
	InputOffset  Vector
	InputOffsetD float64
	InputOffsetN Vector

	IsHit        bool
	HitNormal    *Vector
	HitOffset    Vector
	HitOffsetD   int
	TraceOffset  Vector
	TraceoffsetD int
}

func Trace[T ColliderComparable](colliders *Smap[T, [9]Bounder], target Bounder, offset Vector, excepts Set[T]) *TraceResult {
	offsetl, offsetn := offset.Decompose()
	oli := int(math.Trunc(offsetl)) + 1

	for i := 0; i <= oli; i++ {
		v := offsetn.MulF(float64(i))
		bo := target.Offset(v.X, v.Y, nil)
		r, n := Intersect(colliders, bo, excepts)
		if r {
			DrawDebugTraceDistance(target, i)
			if i == 0 {
				return &TraceResult{
					InputOffset:  offset,
					InputOffsetD: offsetl,
					InputOffsetN: offsetn,

					IsHit:     true,
					HitNormal: n,
				}
			} else {
				return &TraceResult{
					InputOffset:  offset,
					InputOffsetD: offsetl,
					InputOffsetN: offsetn,

					IsHit:        true,
					HitNormal:    n,
					HitOffset:    v,
					HitOffsetD:   i,
					TraceOffset:  offsetn.MulF(float64(i - 1)),
					TraceoffsetD: i - 1,
				}
			}
		}
	}

	return &TraceResult{
		InputOffset:  offset,
		InputOffsetD: offsetl,
		InputOffsetN: offsetn,

		IsHit: false,
	}
}
