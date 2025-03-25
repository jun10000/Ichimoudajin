package utility

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unicode/utf8"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	TypeBool  = reflect.TypeOf(bool(false))
	TypeInt   = reflect.TypeOf(int(0))
	TypeFloat = reflect.TypeOf(float64(0))

	TypeEbitenImagePtr = reflect.TypeOf((*ebiten.Image)(nil))
	TypeRGB            = reflect.TypeOf(RGB{})
)

func RuneToInt(r rune) int {
	return int(r - '0')
}

func HexColorToRGB(hex string) (RGB, error) {
	if utf8.RuneCountInString(hex) != 7 {
		return RGB{}, errors.New("failed to parse HexColor")
	}

	r, err := strconv.ParseUint(hex[1:3], 16, 8)
	if err != nil {
		return RGB{}, err
	}

	g, err := strconv.ParseUint(hex[3:5], 16, 8)
	if err != nil {
		return RGB{}, err
	}

	b, err := strconv.ParseUint(hex[5:7], 16, 8)
	if err != nil {
		return RGB{}, err
	}

	return RGB{uint8(r), uint8(g), uint8(b)}, nil
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
	case TypeRGB:
		v, err := HexColorToRGB(str)
		if err != nil {
			return nil, err
		}
		return v, nil
	default:
		return nil, fmt.Errorf("found unsupported convertion type: %s", typeTo)
	}
}
