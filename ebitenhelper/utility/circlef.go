package utility

import "math"

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

func (c *CircleF) IntersectTo(target Bounder) (result bool, normal *Vector) {
	return target.IntersectFromCircle(c)
}

func (c *CircleF) IntersectFromRectangle(src *RectangleF) (result bool, normal *Vector) {
	p := NewVector(
		ClampFloat(c.OrgX, src.MinX, src.MaxX),
		ClampFloat(c.OrgY, src.MinY, src.MaxY))
	o := NewVector(p.X-c.OrgX, p.Y-c.OrgY) // RectangleF to CircleF version
	rll := o.Length2()

	if rll > (c.Radius * c.Radius) {
		return false, nil
	}

	n := o.DivF(math.Sqrt(rll))
	return true, &n
}

func (c *CircleF) IntersectFromCircle(src *CircleF) (result bool, normal *Vector) {
	d := NewVector(src.OrgX-c.OrgX, src.OrgY-c.OrgY)
	dll := d.Length2()
	if dll > ((src.Radius + c.Radius) * (src.Radius + c.Radius)) {
		return false, nil
	}

	n := d.DivF(math.Sqrt(dll))
	return true, &n
}
