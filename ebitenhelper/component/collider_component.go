package component

import (
	"log"
	"reflect"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

/*
colliderComponentBase gives actors Collider role.
Available T type is pointer.
*/
type colliderComponentBase[T utility.Bounder] struct {
	getBounds   func(T)
	isEnable    bool
	loopOffsets [8]utility.Vector
	cache       [9]utility.Bounder
}

func newColliderComponentBase[T utility.Bounder](getBounds func(T)) *colliderComponentBase[T] {
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

	c := &colliderComponentBase[T]{
		isEnable:    true,
		getBounds:   getBounds,
		loopOffsets: os,
	}

	for i := range 9 {
		c.cache[i] = reflect.New(t.Elem()).Interface().(T)
	}

	return c
}

func (c *colliderComponentBase[T]) UpdateBounds() {
	c.getBounds(c.cache[0].(T))
	for i, v := range c.loopOffsets {
		c.cache[0].Offset(v.X, v.Y, c.cache[i+1])
	}
}

func (c *colliderComponentBase[T]) EnableBounds() {
	c.isEnable = true
}

func (c *colliderComponentBase[T]) DisableBounds() {
	c.isEnable = false
}

func (c *colliderComponentBase[T]) GetRealFirstBounds() utility.Bounder {
	return c.cache[0]
}

func (c *colliderComponentBase[T]) GetRealBounds() []utility.Bounder {
	if utility.GetLevel().IsLooping {
		return c.cache[:]
	} else {
		return c.cache[:1]
	}
}

func (c *colliderComponentBase[T]) GetFirstBounds() utility.Bounder {
	if c.isEnable {
		return c.GetRealFirstBounds()
	} else {
		return nil
	}
}

func (c *colliderComponentBase[T]) GetBounds() []utility.Bounder {
	if c.isEnable {
		return c.GetRealBounds()
	} else {
		return c.cache[:0]
	}
}

/*
StaticColliderComponent gives actors Collider and StaticTransformer role.
Available T type is pointer.
*/
type StaticColliderComponent[T utility.Bounder] struct {
	*colliderComponentBase[T]
	*utility.StaticTransform
}

func NewStaticColliderComponent[T utility.Bounder](sTransform *utility.StaticTransform, getBounds func(T)) *StaticColliderComponent[T] {
	return &StaticColliderComponent[T]{
		colliderComponentBase: newColliderComponentBase(getBounds),

		StaticTransform: sTransform,
	}
}

/*
ColliderComponent gives actors Collider and Transformer role.
Available T type is pointer.
*/
type ColliderComponent[T utility.Bounder] struct {
	*colliderComponentBase[T]
	*utility.Transform
}

func NewColliderComponent[T utility.Bounder](transform *utility.Transform, getBounds func(T)) *ColliderComponent[T] {
	return &ColliderComponent[T]{
		colliderComponentBase: newColliderComponentBase(getBounds),

		Transform: transform,
	}
}

func (c *ColliderComponent[T]) SetLocation(value utility.Vector) {
	c.Transform.SetLocation(value)
	c.UpdateBounds()
}

func (c *ColliderComponent[T]) SetScale(value utility.Vector) {
	c.Transform.SetScale(value)
	c.UpdateBounds()
}
