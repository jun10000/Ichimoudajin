package utility

import (
	"image"
	"math"
)

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

func NewRectangleFFromGoRect(rect image.Rectangle) *RectangleF {
	return NewRectangleF(
		float64(rect.Min.X),
		float64(rect.Min.Y),
		float64(rect.Max.X),
		float64(rect.Max.Y))
}

func (r *RectangleF) TopLeft() Vector {
	return NewVector(r.MinX, r.MinY)
}

func (r *RectangleF) Size() Vector {
	return NewVector(r.MaxX-r.MinX, r.MaxY-r.MinY)
}

func (r *RectangleF) ToCircle() *CircleF {
	org := r.CenterLocation()
	rad := r.TopLeft().Sub(org).Length()
	return NewCircleF(org.X, org.Y, rad)
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

func (r *RectangleF) IntersectTo(target Bounder) (result bool, normal *Vector) {
	return target.IntersectFromRectangle(r)
}

func (r *RectangleF) IntersectFromRectangle(src *RectangleF) (result bool, normal *Vector) {
	xleft := src.MaxX - r.MinX
	if xleft < 0 {
		return false, nil
	}

	xright := r.MaxX - src.MinX
	if xright < 0 {
		return false, nil
	}

	ytop := src.MaxY - r.MinY
	if ytop < 0 {
		return false, nil
	}

	ybottom := r.MaxY - src.MinY
	if ybottom < 0 {
		return false, nil
	}

	if math.Min(xleft, xright) > math.Min(ytop, ybottom) {
		if ytop > ybottom {
			return true, DownVectorPtr()
		} else {
			return true, UpVectorPtr()
		}
	} else {
		if xleft > xright {
			return true, RightVectorPtr()
		} else {
			return true, LeftVectorPtr()
		}
	}
}

func (r *RectangleF) IntersectFromCircle(src *CircleF) (result bool, normal *Vector) {
	p := NewVector(
		ClampFloat(src.OrgX, r.MinX, r.MaxX),
		ClampFloat(src.OrgY, r.MinY, r.MaxY))
	o := NewVector(src.OrgX-p.X, src.OrgY-p.Y) // CircleF to RectangleF version
	rll := o.Length2()

	if rll > (src.Radius * src.Radius) {
		return false, nil
	}

	n := o.DivF(math.Sqrt(rll))
	return true, &n
}
