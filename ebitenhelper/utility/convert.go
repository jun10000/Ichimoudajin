package utility

import (
	"fmt"
	"image/color"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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

func HexStringToColor(hex string, defColor color.Color) color.Color {
	switch utf8.RuneCountInString(hex) {
	case 7:
		r, err := strconv.ParseUint(hex[1:3], 16, 8)
		if err != nil {
			return defColor
		}

		g, err := strconv.ParseUint(hex[3:5], 16, 8)
		if err != nil {
			return defColor
		}

		b, err := strconv.ParseUint(hex[5:7], 16, 8)
		if err != nil {
			return defColor
		}

		return RGB{uint8(r), uint8(g), uint8(b)}
	case 9:
		a, err := strconv.ParseUint(hex[1:3], 16, 8)
		if err != nil {
			return defColor
		}

		r, err := strconv.ParseUint(hex[3:5], 16, 8)
		if err != nil {
			return defColor
		}

		g, err := strconv.ParseUint(hex[5:7], 16, 8)
		if err != nil {
			return defColor
		}

		b, err := strconv.ParseUint(hex[7:9], 16, 8)
		if err != nil {
			return defColor
		}

		return color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	default:
		return defColor
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
		return HexStringToColor(str, ColorTransparent), nil
	default:
		return nil, fmt.Errorf("found unsupported convertion type: %s", typeTo)
	}
}

func StringToFloatSlice(strings []string) ([]float64, error) {
	ret := make([]float64, 0, len(strings))
	for _, s := range strings {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}

		ret = append(ret, f)
	}

	return ret, nil
}

func AppendFontFamiliesFromFilePathsString(in []*text.GoTextFaceSource, pathsString string) []*text.GoTextFaceSource {
	for _, s := range strings.Split(pathsString, ",") {
		f, err := GetFontFromFile(s)
		if err == nil {
			in = append(in, f)
		}
	}

	return in
}
