package utility

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

func (r RectangleF) CenterLocation() Vector {
	return ClampLocation(NewVector((r.MinX+r.MaxX)/2, (r.MinY+r.MaxY)/2))
}

func (r RectangleF) Size() Vector {
	return NewVector(r.MaxX-r.MinX, r.MaxY-r.MinY)
}

func (r RectangleF) BoundingBox() RectangleF {
	return r
}

func (r RectangleF) Offset(value Vector) Bounder {
	return NewRectangleF(r.Location().Add(value), r.Size())
}
