package utility

type Inset struct {
	Top, Right, Bottom, Left float64
}

func NewInset(values ...float64) Inset {
	ret := Inset{}
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
