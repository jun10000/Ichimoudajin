package utility

import "log"

type RectangleF struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

func NewRectangleF(minX, minY, maxX, maxY float64) *RectangleF {
	return &RectangleF{
		MinX: minX,
		MinY: minY,
		MaxX: maxX,
		MaxY: maxY,
	}
}

func (r *RectangleF) TopLeftLocation() Vector {
	return NewVector(r.MinX, r.MinY)
}

func (r *RectangleF) Size() Vector {
	return NewVector(r.MaxX-r.MinX, r.MaxY-r.MinY)
}

func (r *RectangleF) Type() BounderType {
	return BounderTypeRectangle
}

func (r *RectangleF) CenterLocation() Vector {
	return ClampLocation(NewVector((r.MinX+r.MaxX)/2, (r.MinY+r.MaxY)/2))
}

func (r *RectangleF) Offset(x, y float64, output Bounder) Bounder {
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
  - *RectangleF
  - *CircleF
*/
func (r *RectangleF) Intersect(target Bounder) (result bool, normal *Vector) {
	switch target.Type() {
	case BounderTypeRectangle:
		return IntersectRectangleToRectangle(r, target.(*RectangleF))
	case BounderTypeCircle:
		return IntersectCircleToRectangle(target.(*CircleF), r, true)
	default:
		log.Println("Detected unsupported intersection type")
		return false, nil
	}
}
