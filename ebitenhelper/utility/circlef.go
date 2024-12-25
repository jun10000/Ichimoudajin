package utility

import (
	"log"
)

type CircleF struct {
	Origin Vector
	Radius float64
}

func NewCircleF(origin Vector, radius float64) CircleF {
	return CircleF{
		Origin: origin,
		Radius: radius,
	}
}

func (c CircleF) intersectRectangleF(rect RectangleF) (normal Vector) {
	p := NewVector(
		ClampFloat(c.Origin.X, rect.MinX, rect.MaxX),
		ClampFloat(c.Origin.Y, rect.MinY, rect.MaxY))
	r := c.Origin.Sub(p)

	if r.Length() > c.Radius {
		return ZeroVector()
	}

	return r.Normalize()
}

func (c CircleF) intersectCircleF(circle CircleF) (normal Vector) {
	d := c.Origin.Sub(circle.Origin)
	if d.Length() > c.Radius+circle.Radius {
		return ZeroVector()
	}

	return d.Normalize()
}

func (c CircleF) Intersect(target Bounder) (normal Vector) {
	switch v := target.(type) {
	case RectangleF:
		return c.intersectRectangleF(v)
	case CircleF:
		return c.intersectCircleF(v)
	default:
		log.Println("Detected not supported intersect target type")
		return ZeroVector()
	}
}

func (c CircleF) BoundingBox() RectangleF {
	return NewRectangleF(c.Origin.SubF(c.Radius), NewVector(c.Radius*2, c.Radius*2))
}

func (c CircleF) Offset(value Vector) Bounder {
	return NewCircleF(c.Origin.Add(value), c.Radius)
}
