package utility

import "math"

type TraceResult struct {
	IsHit       bool
	HitDistance int
	TraceOffset Vector
	HitNormal   *Vector
}

func Trace[T ColliderComparable](colliders *Smap[T, [9]Bounder], target Bounder, offset Vector, excepts Set[T]) *TraceResult {
	ol, on := offset.Decompose()
	oli := int(math.Trunc(ol)) + 1

	for i := 0; i <= oli; i++ {
		v := on.MulF(float64(i))
		bo := target.Offset(v.X, v.Y, nil)
		r, n := Intersect(colliders, bo, excepts)
		if r {
			DrawDebugTraceDistance(target, i)
			if i == 0 {
				return &TraceResult{
					HitDistance: 0,
					TraceOffset: ZeroVector(),
					HitNormal:   n,
					IsHit:       true,
				}
			} else {
				o := on.MulF(float64(i - 1))
				return &TraceResult{
					HitDistance: i,
					TraceOffset: o,
					HitNormal:   n,
					IsHit:       true,
				}
			}
		}
	}

	return &TraceResult{
		HitDistance: oli + 1,
		TraceOffset: offset,
		HitNormal:   nil,
		IsHit:       false,
	}
}
