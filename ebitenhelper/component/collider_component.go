package component

import (
	"log"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

/*
ColliderComponent gives actors collider.

Available T types:
  - *utility.RectangleF
  - *utility.CircleF
*/
type ColliderComponent[T utility.Bounder] struct {
	getBounds   func(T)
	loopOffsets []utility.Vector
	mainCache   T
	offsetCache T
}

func NewColliderComponent[T utility.Bounder](getBounds func(T)) *ColliderComponent[T] {
	s := utility.ScreenSize.ToVector()
	c := &ColliderComponent[T]{
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
	}

	c.ClearCache()
	return c
}

func (c *ColliderComponent[T]) ClearCache() {
	switch any(c.mainCache).(type) {
	case *utility.RectangleF:
		c.mainCache = any(new(utility.RectangleF)).(T)
		c.offsetCache = any(new(utility.RectangleF)).(T)
	case *utility.CircleF:
		c.mainCache = any(new(utility.CircleF)).(T)
		c.offsetCache = any(new(utility.CircleF)).(T)
	default:
		log.Panicln("Failed to clear ColliderComponent cache.")
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
