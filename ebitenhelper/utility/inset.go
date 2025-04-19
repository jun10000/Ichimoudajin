package utility

import "strings"

type Inset struct {
	Top, Right, Bottom, Left float64
}

func NewInset(values []float64) Inset {
	ret := Inset{}
	if values == nil {
		return ret
	}

	switch len(values) {
	case 1:
		ret.Top = values[0]
		ret.Right = values[0]
		ret.Bottom = values[0]
		ret.Left = values[0]
	case 2:
		ret.Top = values[0]
		ret.Right = values[1]
		ret.Bottom = values[0]
		ret.Left = values[1]
	case 4:
		ret.Top = values[0]
		ret.Right = values[1]
		ret.Bottom = values[2]
		ret.Left = values[3]
	}

	return ret
}

func NewInsetFromString(str string, unit float64) Inset {
	if unit == 0 {
		return Inset{}
	}

	ss := strings.Split(str, ",")
	fs, _ := StringToFloatSlice(ss)
	for i := range fs {
		fs[i] /= unit
	}

	return NewInset(fs)
}
