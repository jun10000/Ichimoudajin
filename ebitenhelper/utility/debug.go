package utility

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	ColorGray  = color.RGBA{R: 128, G: 128, B: 128}
	ColorRed   = color.RGBA{R: 255, G: 8}
	ColorGreen = color.RGBA{G: 255}
	ColorBlue  = color.RGBA{G: 128, B: 255}
)

func DrawDebugLine(start Vector, end Vector, color color.Color) {
	GetGameInstance().AddDrawEvent(func(screen *ebiten.Image) {
		vector.StrokeLine(screen, float32(start.X), float32(start.Y), float32(end.X), float32(end.Y), 2, color, false)
	})
}
