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

func (c CircleF) Intersect(target any) (normal Vector) {
	switch v := target.(type) {
	case RectangleF:
		return c.intersectRectangleF(v)
	default:
		log.Println("Detected not supported intersect target type")
		return ZeroVector()
	}
}
