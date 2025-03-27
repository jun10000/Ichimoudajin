package component

import (
	"log"
	"reflect"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

/*
colliderComBase gives actors Collider role.
Available T type is pointer.
*/
type colliderComBase[T utility.Bounder] struct {
	getBounds   func(T)
	isEnable    bool
	loopOffsets [8]utility.Vector
	cache       [9]utility.Bounder
}

func newColliderComBase[T utility.Bounder](getBounds func(T)) *colliderComBase[T] {
	t := reflect.TypeOf(getBounds).In(0)
	if t.Kind() != reflect.Ptr {
		log.Panic("failed to create colliderComBase: T is not pointer")
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

	c := &colliderComBase[T]{
		isEnable:    true,
		getBounds:   getBounds,
		loopOffsets: os,
	}

	for i := range 9 {
		c.cache[i] = reflect.New(t.Elem()).Interface().(T)
	}

	return c
}

func (c *colliderComBase[T]) UpdateBounds() {
	c.getBounds(c.cache[0].(T))
	for i, v := range c.loopOffsets {
		c.cache[0].Offset(v.X, v.Y, c.cache[i+1])
	}
}

func (c *colliderComBase[T]) EnableBounds() {
	c.isEnable = true
}

func (c *colliderComBase[T]) DisableBounds() {
	c.isEnable = false
}

func (c *colliderComBase[T]) GetRealFirstBounds() utility.Bounder {
	return c.cache[0]
}

func (c *colliderComBase[T]) GetRealBounds() []utility.Bounder {
	if utility.GetLevel().IsLooping {
		return c.cache[:]
	} else {
		return c.cache[:1]
	}
}

func (c *colliderComBase[T]) GetFirstBounds() utility.Bounder {
	if c.isEnable {
		return c.GetRealFirstBounds()
	} else {
		return nil
	}
}

func (c *colliderComBase[T]) GetBounds() []utility.Bounder {
	if c.isEnable {
		return c.GetRealBounds()
	} else {
		return c.cache[:0]
	}
}

func (c *colliderComBase[T]) ReceiveHit(result *utility.TraceResult[utility.Collider]) {
}

/*
StaticColliderCom gives actors Collider and StaticTransformer role.
Available T type is pointer.
*/
type StaticColliderCom[T utility.Bounder] struct {
	*colliderComBase[T]
	*utility.StaticTransform
}

func NewStaticColliderCom[T utility.Bounder](sTransform *utility.StaticTransform, getBounds func(T)) *StaticColliderCom[T] {
	return &StaticColliderCom[T]{
		colliderComBase: newColliderComBase(getBounds),

		StaticTransform: sTransform,
	}
}

/*
ColliderCom gives actors Collider and Transformer role.
Available T type is pointer.
*/
type ColliderCom[T utility.Bounder] struct {
	*colliderComBase[T]
	*utility.Transform
}

func NewColliderCom[T utility.Bounder](transform *utility.Transform, getBounds func(T)) *ColliderCom[T] {
	return &ColliderCom[T]{
		colliderComBase: newColliderComBase(getBounds),

		Transform: transform,
	}
}

func (c *ColliderCom[T]) SetLocation(value utility.Vector) {
	c.Transform.SetLocation(value)
	c.UpdateBounds()
}

func (c *ColliderCom[T]) SetScale(value utility.Vector) {
	c.Transform.SetScale(value)
	c.UpdateBounds()
}
