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
