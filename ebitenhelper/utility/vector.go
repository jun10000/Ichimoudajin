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

func (v Vector) IsZero() bool {
	return (v.X == 0 && v.Y == 0)
}

func (v Vector) Add(value Vector) Vector {
	return NewVector(v.X+value.X, v.Y+value.Y)
}

func (v Vector) Sub(value Vector) Vector {
	return NewVector(v.X-value.X, v.Y-value.Y)
}

func (v Vector) Mul(value Vector) Vector {
	return NewVector(v.X*value.X, v.Y*value.Y)
}

func (v Vector) Div(value Vector) Vector {
	if value.X == 0 || value.Y == 0 {
		return ZeroVector()
	}

	return NewVector(v.X/value.X, v.Y/value.Y)
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

func (v Vector) Floor() Point {
	return NewPoint(int(math.Floor(v.X)), int(math.Floor(v.Y)))
}

func (v Vector) Ceil() Point {
	return NewPoint(int(math.Ceil(v.X)), int(math.Ceil(v.Y)))
}

func (v Vector) Negate() Vector {
	return NewVector(-v.X, -v.Y)
}

func (v Vector) ClampMin(min float64) Vector {
	if v.Length() >= min {
		return v
	}

	return v.Normalize().MulF(min)
}

func (v Vector) ClampMax(max float64) Vector {
	if v.Length() <= max {
		return v
	}

	return v.Normalize().MulF(max)
}

func (v Vector) Clamp(min float64, max float64) Vector {
	return v.ClampMin(min).ClampMax(max)
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) Normalize() Vector {
	l := v.Length()
	if l <= 0 {
		return ZeroVector()
	}

	return NewVector(v.X/l, v.Y/l)
}

func (v Vector) Distance(value Vector) float64 {
	return value.Sub(v).Length()
}

func (v Vector) Dot(value Vector) float64 {
	return v.X*value.X + v.Y*value.Y
}

func (v Vector) Cross(value Vector) Vector3 {
	return NewVector3(0, 0, v.X*value.Y-v.Y*value.X)
}

func (v Vector) CrossingAngle(value Vector) float64 {
	d1, d2 := v.Normalize(), value.Normalize()
	angle := math.Acos(d1.Dot(d2))
	if d1.Cross(d2).Dot(NewVector3(0, 0, 1)) < 0 {
		angle *= -1
	}

	return angle
}

func (v Vector) Reflect(normal Vector, factor float64) Vector {
	n := normal.Normalize()
	return n.MulF(v.Negate().Dot(n) * (1 + factor)).Add(v)
}

func (v Vector) Rotate(angle float64) Vector {
	return NewVector(
		v.X*math.Cos(angle)+v.Y*math.Sin(angle),
		v.X*math.Sin(angle)-v.Y*math.Cos(angle),
	)
}
