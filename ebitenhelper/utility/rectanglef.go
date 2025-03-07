package utility

import "log"

type RectangleF struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

func NewRectangleF(minX, minY, maxX, maxY float64) RectangleF {
	return RectangleF{
		MinX: minX,
		MinY: minY,
		MaxX: maxX,
		MaxY: maxY,
	}
}

func (r RectangleF) Location() Vector {
	return NewVector(r.MinX, r.MinY)
}

func (r RectangleF) CenterLocation() Vector {
	return ClampLocation(NewVector((r.MinX+r.MaxX)/2, (r.MinY+r.MaxY)/2))
}

func (r RectangleF) Size() Vector {
	return NewVector(r.MaxX-r.MinX, r.MaxY-r.MinY)
}

func (r RectangleF) BoundingBox() RectangleF {
	return r
}

func (r RectangleF) Offset(x, y float64, output Bounder) Bounder {
	if o, ok := output.(*RectangleF); ok {
		o.MinX = r.MinX + x
		o.MinY = r.MinY + y
		o.MaxX = r.MaxX + x
		o.MaxY = r.MaxY + y
		return o
	} else {
		return NewRectangleF(r.MinX+x, r.MinY+y, r.MaxX+x, r.MaxY+y)
	}
}

/*
Intersect supports following bounder type
  - RectangleF
  - CircleF
*/
func (r RectangleF) Intersect(target Bounder) (result bool, normal *Vector) {
	switch t := target.(type) {
	case RectangleF:
		return IntersectRectangleToRectangle(r, t)
	case *RectangleF:
		return IntersectRectangleToRectangle(r, *t)
	case CircleF:
		return IntersectCircleToRectangle(t, r, true)
	case *CircleF:
		return IntersectCircleToRectangle(*t, r, true)
	}

	log.Println("Detected unsupported intersection type")
	return false, nil
}
