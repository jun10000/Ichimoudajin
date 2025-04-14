package utility

import "image"

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

func NewRectangleFromGoRect(rect image.Rectangle) *Rectangle {
	return NewRectangle(rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y)
}

func (r *Rectangle) TopLeft() Point {
	return NewPoint(r.MinX, r.MinY)
}

func (r *Rectangle) Size() Point {
	return NewPoint(r.MaxX-r.MinX+1, r.MaxY-r.MinY+1)
}
