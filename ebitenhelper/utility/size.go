package utility

type StaticSize struct {
	value Vector
}

func NewStaticSize(value Vector) StaticSize {
	s := StaticSize{}
	s.value = value.Abs()
	return s
}

func (s *StaticSize) GetSize() Vector {
	return s.value
}

type Size struct {
	StaticSize
}

func NewSize(value Vector) Size {
	s := Size{}
	s.value = value.Abs()
	return s
}

func (s *Size) SetSize(value Vector) {
	s.value = value.Abs()
}
