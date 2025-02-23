package utility

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

func (c CircleF) BoundingBox() RectangleF {
	return NewRectangleF(
		c.Origin.X-c.Radius,
		c.Origin.Y-c.Radius,
		c.Origin.X+c.Radius,
		c.Origin.Y+c.Radius)
}

func (c CircleF) Offset(value Vector) Bounder {
	return NewCircleF(c.Origin.Add(value), c.Radius)
}
