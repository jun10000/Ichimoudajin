package utility

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Landscape struct {
	Image     *ebiten.Image
	TileScale Vector
}

func NewLandscape() *Landscape {
	return &Landscape{
		TileScale: NewVector(1, 1),
	}
}

func (l *Landscape) Draw(screen *ebiten.Image) {
	if l.Image == nil {
		return
	}

	screensize := NewVector(
		float64(screen.Bounds().Dx()),
		float64(screen.Bounds().Dy()),
	)
	tilesize := NewVector(
		float64(l.Image.Bounds().Dx()),
		float64(l.Image.Bounds().Dy()),
	).Mul(l.TileScale)
	tilecount := screensize.Div(tilesize).Ceil()

	for y := 0; y < tilecount.Y; y++ {
		for x := 0; x < tilecount.X; x++ {
			o := &ebiten.DrawImageOptions{}
			o.GeoM.Scale(l.TileScale.X, l.TileScale.Y)
			o.GeoM.Translate(tilesize.X*float64(x), tilesize.Y*float64(y))
			screen.DrawImage(l.Image, o)
		}
	}
}
