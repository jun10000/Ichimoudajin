package utility

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TickCount    int     = 60
	TickDuration float64 = 1.0 / float64(TickCount)
)

var (
	ColorGray  = color.RGBA{R: 128, G: 128, B: 128}
	ColorRed   = color.RGBA{R: 255, G: 8}
	ColorGreen = color.RGBA{G: 255}
	ColorBlue  = color.RGBA{G: 128, B: 255}
)

type PressState int

const (
	PressState_Pressed PressState = iota
	PressState_Released
	PressState_Pressing
)

func NewRectangle(location Point, size Point) image.Rectangle {
	return image.Rect(location.X, location.Y, location.X+size.X, location.Y+size.Y)
}

func GetSubImage(parentimage *ebiten.Image, location Point, size Point) *ebiten.Image {
	if parentimage == nil {
		return nil
	}

	return parentimage.SubImage(NewRectangle(location, size)).(*ebiten.Image)
}

func DrawImage(dst *ebiten.Image, src *ebiten.Image, transform Transformer) {
	if dst == nil || src == nil {
		return
	}

	tl := transform.GetLocation()
	tr := transform.GetRotation()
	ts := transform.GetScale()

	ls := []Vector{tl}
	if GetLevel().IsLooping {
		ss := GetGameInstance().ScreenSize.ToVector()
		ls = append(ls,
			tl.Add(ss.Mul(NewVector(-1, -1))),
			tl.Add(ss.Mul(NewVector(0, -1))),
			tl.Add(ss.Mul(NewVector(1, -1))),
			tl.Add(ss.Mul(NewVector(-1, 0))),
			tl.Add(ss.Mul(NewVector(1, 0))),
			tl.Add(ss.Mul(NewVector(-1, 1))),
			tl.Add(ss.Mul(NewVector(0, 1))),
			tl.Add(ss.Mul(NewVector(1, 1))),
		)
	}

	for _, l := range ls {
		o := &ebiten.DrawImageOptions{}
		o.GeoM.Scale(ts.X, ts.Y)
		o.GeoM.Rotate(tr)
		o.GeoM.Translate(l.X, l.Y)

		dst.DrawImage(src, o)
	}
}

func ExitIfError(err error) {
	if err != nil {
		log.Fatal(err)
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

func RemoveSliceItem[T comparable](slice []T, item T) []T {
	r := make([]T, len(slice))
	for _, v := range slice {
		if v == item {
			continue
		}

		r = append(r, v)
	}

	return r
}
