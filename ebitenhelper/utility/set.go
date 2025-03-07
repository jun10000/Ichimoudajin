package utility

type Set[T comparable] map[T]Empty

func (s Set[T]) Add(value T) {
	s[value] = Empty{}
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) UnionRange(s2 Set[T]) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for v := range s {
			if !yield(v) {
				return
			}
		}
		for v := range s2 {
			if _, ok := s[v]; !ok {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func (s Set[T]) IntersectRange(s2 Set[T]) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for v := range s {
			if _, ok := s2[v]; ok {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func (s Set[T]) SubRange(s2 Set[T]) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for v := range s {
			if _, ok := s2[v]; !ok {
				if !yield(v) {
					return
				}
			}
		}
	}
}
