package utility

type StaticLocation struct {
	value Vector
}

func NewStaticLocation(value Vector) StaticLocation {
	l := StaticLocation{}
	l.value = ClampLocation(value)
	return l
}

func (l *StaticLocation) GetLocation() Vector {
	return l.value
}

type Location struct {
	StaticLocation
}

func NewLocation(value Vector) Location {
	l := Location{}
	l.value = ClampLocation(value)
	return l
}

func (l *Location) SetLocation(value Vector) {
	l.value = ClampLocation(value)
}
