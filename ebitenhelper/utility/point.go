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

func (p Point) AddXY(x int, y int) Point {
	return NewPoint(p.X+x, p.Y+y)
}

func (p Point) SubXY(x int, y int) Point {
	return NewPoint(p.X-x, p.Y-y)
}

func (p Point) MulXY(x int, y int) Point {
	return NewPoint(p.X*x, p.Y*y)
}

func (p Point) DivXY(x int, y int) Point {
	if x == 0 || y == 0 {
		return ZeroPoint()
	}

	return NewPoint(p.X/x, p.Y/y)
}

func (p Point) Add(value Point) Point {
	return NewPoint(p.X+value.X, p.Y+value.Y)
}

func (p Point) Sub(value Point) Point {
	return NewPoint(p.X-value.X, p.Y-value.Y)
}

func (p Point) Mul(value Point) Point {
	return NewPoint(p.X*value.X, p.Y*value.Y)
}

func (p Point) Div(value Point) Point {
	if value.X == 0 || value.Y == 0 {
		return ZeroPoint()
	}

	return NewPoint(p.X/value.X, p.Y/value.Y)
}

func (p Point) Length2() int {
	return p.X*p.X + p.Y*p.Y
}

func (p Point) Length() float64 {
	return math.Sqrt(float64(p.Length2()))
}

func (p Point) Distance2(value Point) int {
	return value.Sub(p).Length2()
}

func (p Point) Distance(value Point) float64 {
	return value.Sub(p).Length()
}
