package utility

import "math"

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

func (p Point) Add(value Point) Point {
	return NewPoint(p.X+value.X, p.Y+value.Y)
}

func (p Point) Sub(value Point) Point {
	return NewPoint(p.X-value.X, p.Y-value.Y)
}

func (p Point) AddXY(x int, y int) Point {
	return NewPoint(p.X+x, p.Y+y)
}

func (p Point) SubXY(x int, y int) Point {
	return NewPoint(p.X-x, p.Y-y)
}

func (p Point) Length() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func (p Point) Distance(value Point) float64 {
	return value.Sub(p).Length()
}
