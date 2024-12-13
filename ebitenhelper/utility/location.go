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

	lv, gi := GetLevel(), GetGameInstance()
	if lv == nil || !lv.IsLooping || gi == nil {
		return
	}

	ss := gi.ScreenSize.ToVector()
	l.value = l.value.Mod(ss)
	if l.value.X < 0 {
		l.value.X += ss.X
	}
	if l.value.Y < 0 {
		l.value.Y += ss.Y
	}
}

func (l *Location) AddLocation(value Vector) {
	l.SetLocation(l.value.Add(value))
}
