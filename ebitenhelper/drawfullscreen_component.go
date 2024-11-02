package ebitenhelper

import "github.com/hajimehoshi/ebiten/v2"

type DrawFullScreenComponent struct {
	Source    *ebiten.Image
	TileScale Vector
}

func NewDrawFullScreenComponent() *DrawFullScreenComponent {
	return &DrawFullScreenComponent{
		TileScale: NewVector(1, 1),
	}
}

func (c *DrawFullScreenComponent) Draw(screen *ebiten.Image) {
	if c.Source == nil {
		return
	}

	screensize := NewVector(
		float64(screen.Bounds().Dx()),
		float64(screen.Bounds().Dy()),
	)
	tilesize := NewVector(
		float64(c.Source.Bounds().Dx()),
		float64(c.Source.Bounds().Dy()),
	).Mul(c.TileScale)
	tilecount := screensize.Div(tilesize).Ceil()

	for y := 0; y < tilecount.Y; y++ {
		for x := 0; x < tilecount.X; x++ {
			o := &ebiten.DrawImageOptions{}
			o.GeoM.Scale(c.TileScale.X, c.TileScale.Y)
			o.GeoM.Translate(tilesize.X*float64(x), tilesize.Y*float64(y))
			screen.DrawImage(c.Source, o)
		}
	}
}
