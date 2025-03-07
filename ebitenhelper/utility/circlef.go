package utility

import "log"

type CircleF struct {
	OrgX   float64
	OrgY   float64
	Radius float64
}

func NewCircleF(orgX, orgY, radius float64) *CircleF {
	return &CircleF{
		OrgX:   orgX,
		OrgY:   orgY,
		Radius: radius,
	}
}

func (c *CircleF) CenterLocation() Vector {
	return NewVector(c.OrgX, c.OrgY)
}

func (c *CircleF) BoundingBox() *RectangleF {
	return NewRectangleF(
		c.OrgX-c.Radius,
		c.OrgY-c.Radius,
		c.OrgX+c.Radius,
		c.OrgY+c.Radius)
}

func (c *CircleF) Offset(x, y float64, output Bounder) Bounder {
	if o, ok := output.(*CircleF); ok {
		o.OrgX = c.OrgX + x
		o.OrgY = c.OrgY + y
		o.Radius = c.Radius
		return o
	} else {
		return NewCircleF(c.OrgX+x, c.OrgY+y, c.Radius)
	}
}

/*
Intersect supports following bounder type
  - *RectangleF
  - *CircleF
*/
func (c *CircleF) Intersect(target Bounder) (result bool, normal *Vector) {
	switch t := target.(type) {
	case *RectangleF:
		return IntersectCircleToRectangle(c, t, false)
	case *CircleF:
		return IntersectCircleToCircle(c, t)
	}

	log.Println("Detected unsupported intersection type")
	return false, nil
}
