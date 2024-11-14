package utility

import "math"

type Rotation struct {
	value float64
}

func NewRotation(value float64) Rotation {
	rot := Rotation{}
	rot.SetRotation(value)
	return rot
}

func ZeroRotation() Rotation {
	return NewRotation(0)
}

func (r *Rotation) GetRotation() float64 {
	return r.value
}

func (r *Rotation) SetRotation(value float64) {
	v := math.Mod(value, 2*math.Pi)
	if v >= math.Pi {
		v -= 2 * math.Pi
	} else if v <= math.Pi*-1 {
		v += 2 * math.Pi
	}
	r.value = v
}
