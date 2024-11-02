package ebitenhelper

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
