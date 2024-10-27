package utility

import "math"

type Vector struct {
	X float64
	Y float64
}

func NewVector(x float64, y float64) Vector {
	return Vector{X: x, Y: y}
}

func (v Vector) Add(vector Vector) Vector {
	return NewVector(v.X+vector.X, v.Y+vector.Y)
}

func (v Vector) MultiplyFloat(f float64) Vector {
	return NewVector(v.X*f, v.Y*f)
}

func (v Vector) DivideFloat(f float64) Vector {
	if f == 0 {
		return NewVector(0, 0)
	}

	return NewVector(v.X/f, v.Y/f)
}

func (v Vector) GetLength() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) Normalize() Vector {
	l := v.GetLength()
	if l <= 0 {
		return NewVector(0, 0)
	}

	return NewVector(v.X/l, v.Y/l)
}

func (v Vector) Clamp(min float64, max float64) Vector {
	l := v.GetLength()
	if l < min {
		l = min
	} else if l > max {
		l = max
	}

	n := v.Normalize()
	return n.MultiplyFloat(l)
}
