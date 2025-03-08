package utility

import "sync"

type Smap[K, V any] struct {
	m sync.Map
}

func NewSmap[K, V any]() *Smap[K, V] {
	return &Smap[K, V]{
		m: sync.Map{},
	}
}

func (s *Smap[K, V]) Load(key K) (value V, ok bool) {
	if v, o := s.m.Load(key); o {
		return v.(V), true
	} else {
		return *new(V), false
	}
}

func (s *Smap[K, V]) Store(key K, value V) {
	s.m.Store(key, value)
}

func (s *Smap[K, V]) Clear() {
	s.m.Clear()
}

func (s *Smap[K, V]) Delete(key K) {
	s.m.Delete(key)
}

func (s *Smap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	if p, l := s.m.Swap(key, value); l {
		return p.(V), true
	} else {
		return *new(V), false
	}
}

func (s *Smap[K, V]) Range() func(yield func(K, V) bool) {
	return func(yield func(K, V) bool) {
		s.m.Range(func(key, value any) bool {
			return yield(key.(K), value.(V))
		})
	}
}
