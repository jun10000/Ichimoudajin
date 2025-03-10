package component

import (
	"log"
	"reflect"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

/*
ColliderComponent gives actors Collider and Transformer role.
Available T type is pointer.
*/
type ColliderComponent[T utility.Bounder] struct {
	utility.Transform

	getBounds   func(T)
	loopOffsets [8]utility.Vector
	cache       [9]utility.Bounder
}

func NewColliderComponent[T utility.Bounder](getBounds func(T)) *ColliderComponent[T] {
	t := reflect.TypeOf(getBounds).In(0)
	if t.Kind() != reflect.Ptr {
		log.Panic("failed to create ColliderComponent: T is not pointer")
	}

	s := utility.GetScreenSize().ToVector()
	os := [8]utility.Vector{
		utility.NewVector(-s.X, -s.Y),
		utility.NewVector(0, -s.Y),
		utility.NewVector(s.X, -s.Y),
		utility.NewVector(-s.X, 0),
		utility.NewVector(s.X, 0),
		utility.NewVector(-s.X, s.Y),
		utility.NewVector(0, s.Y),
		utility.NewVector(s.X, s.Y),
	}

	c := &ColliderComponent[T]{
		Transform:   utility.DefaultTransform(),
		getBounds:   getBounds,
		loopOffsets: os,
	}

	for i := range 9 {
		c.cache[i] = reflect.New(t.Elem()).Interface().(T)
	}

	return c
}

func (c *ColliderComponent[T]) UpdateColliderBounds() {
	c.getBounds(c.cache[0].(T))
	for i, v := range c.loopOffsets {
		c.cache[0].Offset(v.X, v.Y, c.cache[i+1])
	}
}

func (c *ColliderComponent[T]) GetMainColliderBounds() utility.Bounder {
	return c.cache[0]
}

func (c *ColliderComponent[T]) GetColliderBounds() [9]utility.Bounder {
	return c.cache
}

func (c *ColliderComponent[T]) SetLocation(value utility.Vector) {
	c.Transform.SetLocation(value)
	c.UpdateColliderBounds()
}

func (c *ColliderComponent[T]) SetScale(value utility.Vector) {
	c.Transform.SetScale(value)
	c.UpdateColliderBounds()
}
