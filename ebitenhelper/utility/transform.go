package utility

type Transform struct {
	Location
	Rotation
	Scale
}

func NewTransform(location Location, rotation Rotation, scale Scale) Transform {
	return Transform{
		Location: location,
		Rotation: rotation,
		Scale:    scale,
	}
}

func DefaultTransform() Transform {
	return NewTransform(ZeroLocation(), ZeroRotation(), DefaultScale())
}
