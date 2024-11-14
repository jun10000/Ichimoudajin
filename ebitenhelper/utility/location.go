package utility

type Location struct {
	value Vector
}

func NewLocation(value Vector) Location {
	location := Location{}
	location.SetLocation(value)
	return location
}

func ZeroLocation() Location {
	return NewLocation(ZeroVector())
}

func (l *Location) GetLocation() Vector {
	return l.value
}

func (l *Location) SetLocation(value Vector) {
	l.value = value
}

func (l *Location) AddLocation(value Vector) {
	l.value = l.value.Add(value)
}
