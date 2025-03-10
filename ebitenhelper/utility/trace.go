package utility

type TraceResult struct {
	InputOffset  Vector
	InputOffsetD float64
	InputOffsetN Vector

	IsHit        bool
	IsFirstHit   bool
	HitNormal    *Vector
	TraceOffset  Vector
	TraceoffsetD int
}

func NewTraceResult(offset Vector) *TraceResult {
	ol, on := offset.Decompose()
	return &TraceResult{
		InputOffset:  offset,
		InputOffsetD: ol,
		InputOffsetN: on,
	}
}

func GetColliderBounds[T ColliderComparable](colliders *Smap[T, [9]Bounder], excepts Set[T]) func(yield func(Bounder) bool) {
	return func(yield func(Bounder) bool) {
		loop := GetLevel().IsLooping
		for c, bs := range colliders.Range() {
			if excepts != nil && excepts.Contains(c) {
				continue
			}

			if loop {
				for i := range 9 {
					if !yield(bs[i]) {
						return
					}
				}
			} else {
				if !yield(bs[0]) {
					return
				}
			}
		}
	}
}

func Intersect[T ColliderComparable](colliders *Smap[T, [9]Bounder], target Bounder, excepts Set[T]) (result bool, normal *Vector) {
	for b := range GetColliderBounds(colliders, excepts) {
		r, n := target.IntersectTo(b)
		if r {
			return true, n
		}
	}

	return false, nil
}

func Trace[T ColliderComparable](colliders *Smap[T, [9]Bounder], target Bounder, offset Vector, excepts Set[T]) *TraceResult {
	ret := NewTraceResult(offset)

	for i := range int(ret.InputOffsetD) + 2 {
		o := ret.InputOffsetN.MulF(float64(i))
		b := target.Offset(o.X, o.Y, nil)
		ret.IsHit, ret.HitNormal = Intersect(colliders, b, excepts)
		if ret.IsHit {
			DrawDebugTraceDistance(target, i)
			if i == 0 {
				ret.IsFirstHit = true
				return ret
			} else {
				ret.TraceOffset = o.Sub(ret.InputOffsetN)
				ret.TraceoffsetD = i - 1
				DrawDebugTraceResult(ret, target)
				return ret
			}
		}
	}

	return ret
}
