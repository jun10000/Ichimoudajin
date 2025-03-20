package utility

type TraceResult[T ColliderComparable] struct {
	InputOffset  Vector
	InputOffsetD float64
	InputOffsetN Vector

	IsHit        bool
	IsFirstHit   bool
	HitCollider  T
	HitNormal    *Vector
	TraceOffset  Vector
	TraceoffsetD int
}

func NewTraceResult[T ColliderComparable](offset Vector) *TraceResult[T] {
	ol, on := offset.Decompose()
	return &TraceResult[T]{
		InputOffset:  offset,
		InputOffsetD: ol,
		InputOffsetN: on,
	}
}

func GetColliderBounds[T ColliderComparable](colliders []T, excepts Set[T]) func(yield func(T, Bounder) bool) {
	return func(yield func(T, Bounder) bool) {
		for _, c := range colliders {
			if excepts != nil && excepts.Contains(c) {
				continue
			}

			for _, b := range c.GetColliderBounds() {
				if !yield(c, b) {
					return
				}
			}
		}
	}
}

func Intersect[T ColliderComparable](colliders []T, target Bounder, excepts Set[T]) (result bool, collider T, normal *Vector) {
	for c, b := range GetColliderBounds(colliders, excepts) {
		r, n := target.IntersectTo(b)
		if r {
			return true, c, n
		}
	}

	return false, *new(T), nil
}

func IntersectAll[T ColliderComparable](colliders []T, target Bounder, excepts Set[T]) (result bool, iColliders []T, normal Vector) {
	rr := false
	rcs := make([]T, 0, len(colliders))
	rn := ZeroVector()

	for c, b := range GetColliderBounds(colliders, excepts) {
		r, n := target.IntersectTo(b)
		if r {
			rr = true
			rcs = append(rcs, c)
			rn = rn.Add(*n)
		}
	}

	return rr, rcs, rn.Normalize()
}

func Trace[T ColliderComparable](colliders []T, target Bounder, offset Vector, excepts Set[T]) *TraceResult[T] {
	ret := NewTraceResult[T](offset)

	for i := range int(ret.InputOffsetD) + 2 {
		o := ret.InputOffsetN.MulF(float64(i))
		b := target.Offset(o.X, o.Y, nil)
		ret.IsHit, ret.HitCollider, ret.HitNormal = Intersect(colliders, b, excepts)
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
