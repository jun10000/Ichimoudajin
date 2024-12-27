package utility

type Transform struct {
	Location
	Rotation
	Scale
}

func NewTransform(location Vector, rotation float64, scale Vector) *Transform {
	return &Transform{
		Location: NewLocation(location),
		Rotation: NewRotation(rotation),
		Scale:    NewScale(scale),
	}
}

func DefaultTransform() Transform {
	return Transform{
		Location: ZeroLocation(),
		Rotation: ZeroRotation(),
		Scale:    DefaultScale(),
	}
}
