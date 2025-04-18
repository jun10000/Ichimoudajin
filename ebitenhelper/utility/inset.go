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

func NewInsetFromString(str string) Inset {
	ss := strings.Split(str, ",")
	fs, _ := StringToFloatSlice(ss)
	return NewInset(fs)
}

func (i Inset) MulF(value float64) Inset {
	return Inset{i.Top * value, i.Right * value, i.Bottom * value, i.Left * value}
}

func (i Inset) DivF(value float64) Inset {
	if value == 0 {
		return Inset{}
	}

	return Inset{i.Top / value, i.Right / value, i.Bottom / value, i.Left / value}
}
