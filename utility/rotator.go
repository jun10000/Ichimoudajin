package utility

import "math"

type Rotator struct {
	value float64
}

func NewRotator(value float64) Rotator {
	rot := Rotator{}
	rot.Set(value)
	return rot
}

func (r *Rotator) Get() float64 {
	return r.value
}

func (r *Rotator) Set(value float64) {
	v := math.Mod(value, 2*math.Pi)
	if v >= math.Pi {
		v -= 2 * math.Pi
	} else if v <= math.Pi*-1 {
		v += 2 * math.Pi
	}
	r.value = v
}
