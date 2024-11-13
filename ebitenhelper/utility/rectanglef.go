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

func (r RectangleF) Intersect(bounds RectangleF) bool {
	return (r.MinX <= bounds.MaxX &&
		r.MaxX >= bounds.MinX &&
		r.MinY <= bounds.MaxY &&
		r.MaxY >= bounds.MinY)
}
