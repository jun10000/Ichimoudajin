package component

import (
	"log"
	"reflect"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

/*
ColliderComponent gives actors collider.
Available T type is pointer.
*/
type ColliderComponent[T utility.Bounder] struct {
	getBounds   func(T)
	loopOffsets []utility.Vector
	mainCache   T
	offsetCache T
}

func NewColliderComponent[T utility.Bounder](getBounds func(T)) *ColliderComponent[T] {
	t := reflect.TypeOf(getBounds).In(0)
	if t.Kind() != reflect.Ptr {
		log.Panic("failed to create ColliderComponent: T is not pointer")
	}

	s := utility.ScreenSize.ToVector()

	return &ColliderComponent[T]{
		getBounds: getBounds,
		loopOffsets: []utility.Vector{
			utility.NewVector(-s.X, -s.Y),
			utility.NewVector(0, -s.Y),
			utility.NewVector(s.X, -s.Y),
			utility.NewVector(-s.X, 0),
			utility.NewVector(s.X, 0),
			utility.NewVector(-s.X, s.Y),
			utility.NewVector(0, s.Y),
			utility.NewVector(s.X, s.Y),
		},
		mainCache:   reflect.New(t.Elem()).Interface().(T),
		offsetCache: reflect.New(t.Elem()).Interface().(T),
	}
}

func (c *ColliderComponent[T]) GetMainColliderBounds() utility.Bounder {
	c.getBounds(c.mainCache)
	return c.mainCache
}

func (c *ColliderComponent[T]) GetColliderBounds() func(yield func(utility.Bounder) bool) {
	return func(yield func(utility.Bounder) bool) {
		b := c.GetMainColliderBounds()
		if !yield(b) {
			return
		}
		if !utility.GetLevel().IsLooping {
			return
		}

		for _, v := range c.loopOffsets {
			b.Offset(v.X, v.Y, c.offsetCache)
			if !yield(c.offsetCache) {
				return
			}
		}
	}
}
