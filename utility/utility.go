package utility

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TickCount    int     = 60
	TickDuration float64 = 1.0 / float64(TickCount)
)

func NewRectangle(location Point, size Point) image.Rectangle {
	return image.Rect(location.X, location.Y, location.X+size.X, location.Y+size.Y)
}

func GetSubImage(parentimage *ebiten.Image, location Point, size Point) *ebiten.Image {
	return parentimage.SubImage(NewRectangle(location, size)).(*ebiten.Image)
}
