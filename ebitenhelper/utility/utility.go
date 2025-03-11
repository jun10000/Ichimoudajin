package utility

import (
	"image"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jun10000/Ichimoudajin/assets"
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

func GetImageFile(filename string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFileSystem(assets.Assets, filename)
	PanicIfError(err)
	return image
}

func GetSubImage(parentimage *ebiten.Image, location Point, size Point) *ebiten.Image {
	if parentimage == nil {
		return nil
	}

	r := image.Rect(location.X, location.Y, location.X+size.X, location.Y+size.Y)
	return parentimage.SubImage(r).(*ebiten.Image)
}

func DrawImage(dst *ebiten.Image, src *ebiten.Image, transform Transformer) {
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

func RuneToInt(r rune) int {
	return int(r - '0')
}

func StringToInt(str string, output *int) error {
	v, err := strconv.Atoi(str)
	if err != nil {
		return err
	}

	*output = v
	return nil
}

func StringToFloat(str string, output *float64) error {
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	*output = v
	return nil
}
