package utility

import "math"

type RectangleF struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

func NewRectangleF(location Vector, size Vector) RectangleF {
	return RectangleF{
		MinX: location.X,
		MinY: location.Y,
		MaxX: location.X + size.X,
		MaxY: location.Y + size.Y,
	}
}

func (r RectangleF) Location() Vector {
	return NewVector(r.MinX, r.MinY)
}

func (r RectangleF) Size() Vector {
	return NewVector(r.MaxX-r.MinX, r.MaxY-r.MinY)
}

func (rect1 RectangleF) Intersect(rect2 RectangleF) (normal Vector, result bool) {
	xleft := rect1.MaxX - rect2.MinX
	xright := rect2.MaxX - rect1.MinX
	ytop := rect1.MaxY - rect2.MinY
	ybottom := rect2.MaxY - rect1.MinY

	if xleft < 0 || xright < 0 || ytop < 0 || ybottom < 0 {
		return ZeroVector(), false
	}

	isright := xleft > xright
	isbottom := ytop > ybottom
	isy := math.Min(xleft, xright) > math.Min(ytop, ybottom)

	if isy {
		if isbottom {
			return NewVector(0, 1), true
		} else {
			return NewVector(0, -1), true
		}
	} else {
		if isright {
			return NewVector(1, 0), true
		} else {
			return NewVector(-1, 0), true
		}
	}
}
