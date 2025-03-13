package utility

type StaticTransform struct {
	StaticLocation
	StaticRotation
	StaticScale
}

func NewStaticTransform(location Vector, rotation float64, scale Vector) *StaticTransform {
	return &StaticTransform{
		StaticLocation: NewStaticLocation(location),
		StaticRotation: NewStaticRotation(rotation),
		StaticScale:    NewStaticScale(scale),
	}
}

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
