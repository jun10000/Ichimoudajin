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
	return NewRectangleF(c.Origin.SubF(c.Radius), NewVector(c.Radius*2, c.Radius*2))
}

func (c CircleF) Offset(value Vector) Bounder {
	return NewCircleF(c.Origin.Add(value), c.Radius)
}
