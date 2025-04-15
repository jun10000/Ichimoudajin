package utility

import (
	"errors"
	"fmt"
	"image/color"
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

func HexStringToColor(hex string, defColor color.Color) (color.Color, error) {
	switch utf8.RuneCountInString(hex) {
	case 7:
		r, err := strconv.ParseUint(hex[1:3], 16, 8)
		if err != nil {
			return defColor, err
		}

		g, err := strconv.ParseUint(hex[3:5], 16, 8)
		if err != nil {
			return defColor, err
		}

		b, err := strconv.ParseUint(hex[5:7], 16, 8)
		if err != nil {
			return defColor, err
		}

		return RGB{uint8(r), uint8(g), uint8(b)}, nil
	case 9:
		a, err := strconv.ParseUint(hex[1:3], 16, 8)
		if err != nil {
			return defColor, err
		}

		r, err := strconv.ParseUint(hex[3:5], 16, 8)
		if err != nil {
			return defColor, err
		}

		g, err := strconv.ParseUint(hex[5:7], 16, 8)
		if err != nil {
			return defColor, err
		}

		b, err := strconv.ParseUint(hex[7:9], 16, 8)
		if err != nil {
			return defColor, err
		}

		return color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}, nil
	default:
		return defColor, errors.New("failed to parse HexColor")
	}
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
		v, err := HexStringToColor(str, ColorTransparent)
		if err != nil {
			return nil, err
		}
		return v, nil
	default:
		return nil, fmt.Errorf("found unsupported convertion type: %s", typeTo)
	}
}
