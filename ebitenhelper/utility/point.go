package utility

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}

func ZeroPoint() Point {
	return NewPoint(0, 0)
}

func (p Point) ToVector() Vector {
	return NewVector(float64(p.X), float64(p.Y))
}
