package utility

type Transform struct {
	location Vector
	rotation Rotation
	scale    Vector
}

func NewTransform(location Vector, rotation Rotation, scale Vector) Transform {
	return Transform{
		location: location,
		rotation: rotation,
		scale:    scale,
	}
}

func DefaultTransform() Transform {
	return NewTransform(ZeroVector(), ZeroRotation(), NewVector(1, 1))
}

func (t *Transform) GetLocation() Vector {
	return t.location
}

func (t *Transform) SetLocation(location Vector) {
	t.location = location
}

func (t *Transform) AddLocation(location Vector) {
	t.location = t.location.Add(location)
}

func (t *Transform) GetRotation() float64 {
	return t.rotation.Get()
}

func (t *Transform) SetRotation(value float64) {
	t.rotation.Set(value)
}

func (t *Transform) GetScale() Vector {
	return t.scale
}

func (t *Transform) SetScale(scale Vector) {
	t.scale = scale
}
