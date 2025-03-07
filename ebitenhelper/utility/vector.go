package utility

import (
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func NewVector(x float64, y float64) Vector {
	return Vector{X: x, Y: y}
}

func NewVectorPtr(x float64, y float64) *Vector {
	return &Vector{X: x, Y: y}
}

func ZeroVector() Vector {
	return NewVector(0, 0)
}

func ZeroVectorPtr() *Vector {
	return NewVectorPtr(0, 0)
}

func UpVector() Vector {
	return NewVector(0, -1)
}

func UpVectorPtr() *Vector {
	return NewVectorPtr(0, -1)
}

func DownVector() Vector {
	return NewVector(0, 1)
}

func DownVectorPtr() *Vector {
	return NewVectorPtr(0, 1)
}

func LeftVector() Vector {
	return NewVector(-1, 0)
}

func LeftVectorPtr() *Vector {
	return NewVectorPtr(-1, 0)
}

func RightVector() Vector {
	return NewVector(1, 0)
}

func RightVectorPtr() *Vector {
	return NewVectorPtr(1, 0)
}

func (v Vector) IsZero() bool {
	return (v.X == 0 && v.Y == 0)
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

func (v Vector) ModF(value float64) Vector {
	if value == 0 {
		return ZeroVector()
	}

	return NewVector(math.Mod(v.X, value), math.Mod(v.Y, value))
}

func (v Vector) AddXY(x, y float64) Vector {
	return NewVector(v.X+x, v.Y+y)
}

func (v Vector) SubXY(x, y float64) Vector {
	return NewVector(v.X-x, v.Y-y)
}

func (v Vector) MulXY(x, y float64) Vector {
	return NewVector(v.X*x, v.Y*y)
}

func (v Vector) DivXY(x, y float64) Vector {
	if x == 0 || y == 0 {
		return ZeroVector()
	}

	return NewVector(v.X/x, v.Y/y)
}

func (v Vector) ModXY(x, y float64) Vector {
	if x == 0 || y == 0 {
		return ZeroVector()
	}

	return NewVector(math.Mod(v.X, x), math.Mod(v.Y, y))
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

func (v Vector) Mod(value Vector) Vector {
	if value.X == 0 || value.Y == 0 {
		return ZeroVector()
	}

	return NewVector(math.Mod(v.X, value.X), math.Mod(v.Y, value.Y))
}

func (v Vector) Trunc() Point {
	return NewPoint(int(math.Trunc(v.X)), int(math.Trunc(v.Y)))
}

func (v Vector) Negate() Vector {
	return NewVector(-v.X, -v.Y)
}

func (v Vector) Length2() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.Length2())
}

func (v Vector) Normalize() Vector {
	ll := v.Length2()
	if ll == 0 {
		return ZeroVector()
	}

	l := math.Sqrt(ll)
	return NewVector(v.X/l, v.Y/l)
}

func (v Vector) Decompose() (length float64, normal Vector) {
	l := v.Length()
	return l, v.DivF(l)
}

func (v Vector) ClampMin(min float64) Vector {
	vll := v.Length2()
	if vll >= (min * min) {
		return v
	}

	return v.DivF(math.Sqrt(vll)).MulF(min)
}

func (v Vector) ClampMax(max float64) Vector {
	vll := v.Length2()
	if vll <= (max * max) {
		return v
	}

	return v.DivF(math.Sqrt(vll)).MulF(max)
}

func (v Vector) Clamp(min float64, max float64) Vector {
	return v.ClampMin(min).ClampMax(max)
}

func (v Vector) Dot(value Vector) float64 {
	return v.X*value.X + v.Y*value.Y
}

func (v Vector) CrossZ(value Vector) float64 {
	return v.X*value.Y - v.Y*value.X
}

func (v Vector) CrossingAngle(value Vector) float64 {
	d1, d2 := v.Normalize(), value.Normalize()
	angle := math.Acos(d1.Dot(d2))
	if d1.CrossZ(d2) < 0 {
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
		-v.X*math.Sin(angle)+v.Y*math.Cos(angle),
	)
}

func (v Vector) String() string {
	return fmt.Sprintf("(%v, %v)", v.X, v.Y)
}
