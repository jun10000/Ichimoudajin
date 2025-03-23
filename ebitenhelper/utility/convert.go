package utility

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	TypeBool  = reflect.TypeOf(bool(false))
	TypeInt   = reflect.TypeOf(int(0))
	TypeFloat = reflect.TypeOf(float64(0))

	TypeEbitenImagePtr = reflect.TypeOf((*ebiten.Image)(nil))
)

func RuneToInt(r rune) int {
	return int(r - '0')
}

func ConvertFromString(str string, typeTo reflect.Type) (any, error) {
	switch typeTo {
	case TypeBool:
		v, err := strconv.ParseBool(str)
		if err != nil {
			return nil, err
		}
		return v, nil
	case TypeInt:
		v, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		return v, nil
	case TypeFloat:
		v, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		return v, nil
	case TypeEbitenImagePtr:
		v, err := GetImageFromFile(str)
		if err != nil {
			return nil, err
		}
		return v, nil
	default:
		return nil, fmt.Errorf("found unsupported convertion type: %s", typeTo)
	}
}
