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

func (c *CircleF) Type() BounderType {
	return BounderTypeCircle
}

func (c *CircleF) CenterLocation() Vector {
	return ClampLocation(NewVector(c.OrgX, c.OrgY))
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
	switch target.Type() {
	case BounderTypeRectangle:
		return IntersectCircleToRectangle(c, target.(*RectangleF), false)
	case BounderTypeCircle:
		return IntersectCircleToCircle(c, target.(*CircleF))
	default:
		log.Println("Detected unsupported intersection type")
		return false, nil
	}
}
