package utility

type Scale struct {
	value Vector
}

func NewScale(value Vector) Scale {
	scale := Scale{
		value: NewVector(1, 1),
	}
	scale.SetScale(value)

	return scale
}

func DefaultScale() Scale {
	return NewScale(NewVector(1, 1))
}

func (s *Scale) GetScale() Vector {
	return s.value
}

func (s *Scale) SetScale(value Vector) {
	if value.X <= 0 || value.Y <= 0 {
		return
	}

	s.value = value
}
