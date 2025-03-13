package utility

type StaticRotation struct {
	value float64
}

func NewStaticRotation(value float64) StaticRotation {
	r := StaticRotation{}
	r.value = ClampRotation(value)
	return r
}

func (r *StaticRotation) GetRotation() float64 {
	return r.value
}

type Rotation struct {
	StaticRotation
}

func NewRotation(value float64) Rotation {
	r := Rotation{}
	r.value = ClampRotation(value)
	return r
}

func (r *Rotation) SetRotation(value float64) {
	r.value = ClampRotation(value)
}
