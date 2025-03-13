package utility

type StaticScale struct {
	value Vector
}

func NewStaticScale(value Vector) StaticScale {
	s := StaticScale{}
	s.value = value.Abs()
	return s
}

func (s *StaticScale) GetScale() Vector {
	return s.value
}

type Scale struct {
	StaticScale
}

func NewScale(value Vector) Scale {
	s := Scale{}
	s.value = value.Abs()
	return s
}

func (s *Scale) SetScale(value Vector) {
	s.value = value.Abs()
}
