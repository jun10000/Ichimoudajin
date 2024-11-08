package utility

type Transform struct {
	Location Vector
	Rotation Rotator
	Scale    Vector
}

func NewTransform(location Vector, rotation Rotator, scale Vector) Transform {
	return Transform{
		Location: location,
		Rotation: rotation,
		Scale:    scale,
	}
}

func DefaultTransform() Transform {
	return NewTransform(ZeroVector(), ZeroRotator(), NewVector(1, 1))
}
