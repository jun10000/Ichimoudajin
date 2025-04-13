package utility

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"reflect"
	"runtime"
	"slices"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type PressState int

const (
	PressStatePressed PressState = iota
	PressStateReleased
	PressStatePressing
)

type Empty struct{}

func RemoveAllStrings(src string, targets ...string) string {
	ret := src
	for _, s := range targets {
		ret = strings.ReplaceAll(ret, s, "")
	}
	return ret
}

func CallMethodByName(parent any, name string, args ...any) ([]any, error) {
	if parent == nil {
		return nil, fmt.Errorf("argument 'parent' in CallMethodByName is nil")
	}

	m := reflect.ValueOf(parent).MethodByName(name)
	if !m.IsValid() {
		return nil, fmt.Errorf("method '%s' is not found", name)
	}

	argc := len(args)
	rargc := m.Type().NumIn()
	if rargc > argc {
		return nil, fmt.Errorf("method '%s' has invalid argument counts: %d", name, rargc)
	}

	rargs := make([]reflect.Value, 0, argc)
	for _, arg := range args {
		rargs = append(rargs, reflect.ValueOf(arg))
	}

	rrets := m.Call(rargs[:rargc])
	rets := make([]any, 0, len(rrets))
	for _, ret := range rrets {
		rets = append(rets, ret.Interface())
	}

	return rets, nil
}

func Exit(code int) {
	if runtime.GOOS != "js" {
		os.Exit(code)
	}
}

func GetSubImage(parentimage *ebiten.Image, location Point, size Point) *ebiten.Image {
	if parentimage == nil {
		return nil
	}

	r := image.Rect(location.X, location.Y, location.X+size.X, location.Y+size.Y)
	return parentimage.SubImage(r).(*ebiten.Image)
}

func DrawLine(screen *ebiten.Image, start Vector, end Vector, width float32, color color.Color, antialias bool) {
	if screen == nil || width <= 0 || color == nil {
		return
	}

	vector.StrokeLine(screen, float32(start.X), float32(start.Y), float32(end.X), float32(end.Y), width, color, antialias)
}

func DrawRectangle(screen *ebiten.Image, topLeft Vector, size Vector, borderWidth float32, borderColor color.Color, fillColor color.Color, antialias bool) {
	if screen == nil || size.X <= 0 || size.Y <= 0 {
		return
	}

	if fillColor != nil {
		if _, _, _, a := fillColor.RGBA(); a > 0 {
			vector.DrawFilledRect(screen, float32(topLeft.X), float32(topLeft.Y), float32(size.X), float32(size.Y), fillColor, antialias)
		}
	}

	if borderColor != nil && borderWidth > 0 {
		vector.StrokeRect(screen, float32(topLeft.X), float32(topLeft.Y), float32(size.X), float32(size.Y), borderWidth, borderColor, antialias)
	}
}

func DrawCircle(screen *ebiten.Image, center Vector, radius float32, borderWidth float32, borderColor color.Color, fillColor color.Color, antialias bool) {
	if screen == nil || radius <= 0 {
		return
	}

	cx := float32(center.X)
	cy := float32(center.Y)
	cr := float32(radius)

	if fillColor != nil {
		if _, _, _, a := fillColor.RGBA(); a > 0 {
			vector.DrawFilledCircle(screen, cx, cy, cr, fillColor, antialias)
		}
	}

	if borderColor != nil && borderWidth > 0 {
		vector.StrokeCircle(screen, cx, cy, cr, borderWidth, borderColor, antialias)
	}
}

func DrawImage(dst *ebiten.Image, src *ebiten.Image, transform StaticTransformer) {
	if dst == nil || src == nil {
		return
	}

	tl := transform.GetLocation()
	tr := transform.GetRotation()
	ts := transform.GetScale()

	var ls []Vector
	if !GetLevel().IsLooping {
		ls = []Vector{tl}
	} else {
		ss := GetScreenSize().ToVector()
		ls = []Vector{
			tl,
			tl.AddXY(-ss.X, -ss.Y),
			tl.AddXY(0, -ss.Y),
			tl.AddXY(ss.X, -ss.Y),
			tl.AddXY(-ss.X, 0),
			tl.AddXY(ss.X, 0),
			tl.AddXY(-ss.X, ss.Y),
			tl.AddXY(0, ss.Y),
			tl.AddXY(ss.X, ss.Y),
		}
	}

	for _, l := range ls {
		o := &ebiten.DrawImageOptions{}
		o.GeoM.Scale(ts.X, ts.Y)
		o.GeoM.Rotate(tr)
		o.GeoM.Translate(l.X, l.Y)

		dst.DrawImage(src, o)
	}
}

func GetCursorPosition() Point {
	x, y := ebiten.CursorPosition()
	return NewPoint(x, y)
}

func PanicIfError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ClampFloat(value float64, min float64, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}

	return value
}

func ClampLocation(location Vector) Vector {
	r := location
	lv := GetLevel()

	if lv == nil || !lv.IsLooping {
		return r
	}

	ss := GetScreenSize().ToVector()
	r = r.Mod(ss)
	if r.X < 0 {
		r.X += ss.X
	}
	if r.Y < 0 {
		r.Y += ss.Y
	}

	return r
}

func ClampRotation(rotation float64) float64 {
	r := math.Mod(rotation, 2*math.Pi)
	if r >= math.Pi {
		r -= 2 * math.Pi
	} else if r <= math.Pi*-1 {
		r += 2 * math.Pi
	}

	return r
}

func RemoveSliceItem[T comparable](slice []T, item T) []T {
	i := slices.Index(slice, item)
	if i == -1 {
		return slice
	}

	return slices.Delete(slice, i, i+1)
}

func DegreeToRadian(degree float64) float64 {
	return degree / 180.0 * math.Pi
}

func RadianToDegree(radian float64) float64 {
	return radian / math.Pi * 180.0
}
