package utility

type Rectangle struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
}

func NewRectangle(minX, minY, maxX, maxY int) *Rectangle {
	return &Rectangle{
		MinX: minX,
		MinY: minY,
		MaxX: maxX,
		MaxY: maxY,
	}
}

func (r *Rectangle) TopLeft() Point {
	return NewPoint(r.MinX, r.MinY)
}

func (r *Rectangle) Size() Point {
	return NewPoint(r.MaxX-r.MinX+1, r.MaxY-r.MinY+1)
}
