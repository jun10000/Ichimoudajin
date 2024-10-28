package utility

import "math"

type Vector struct {
	X float64
	Y float64
}

func NewVector(x float64, y float64) Vector {
	return Vector{X: x, Y: y}
}

func ZeroVector() Vector {
	return NewVector(0, 0)
}

func (v Vector) Add(value Vector) Vector {
	return NewVector(v.X+value.X, v.Y+value.Y)
}

func (v Vector) Sub(value Vector) Vector {
	return NewVector(v.X-value.X, v.Y-value.Y)
}

func (v Vector) AddF(value float64) Vector {
	return NewVector(v.X+value, v.Y+value)
}

func (v Vector) SubF(value float64) Vector {
	return NewVector(v.X-value, v.Y-value)
}

func (v Vector) MulF(value float64) Vector {
	return NewVector(v.X*value, v.Y*value)
}

func (v Vector) DivF(value float64) Vector {
	if value == 0 {
		return ZeroVector()
	}

	return NewVector(v.X/value, v.Y/value)
}

func (v Vector) GetLength() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) Normalize() Vector {
	l := v.GetLength()
	if l <= 0 {
		return ZeroVector()
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
	return n.MulF(l)
}
