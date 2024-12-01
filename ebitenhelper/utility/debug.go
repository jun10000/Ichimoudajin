package utility

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	ColorRed   = color.RGBA{R: 224, G: 32}
	ColorGreen = color.RGBA{G: 255}
	ColorBlue  = color.RGBA{G: 128, B: 255}
)

func DrawDebugLine(image *ebiten.Image, start Vector, end Vector, color color.Color) {
	vector.StrokeLine(image, float32(start.X), float32(start.Y), float32(end.X), float32(end.Y), 2, color, false)
}
