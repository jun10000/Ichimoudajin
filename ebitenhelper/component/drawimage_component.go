package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type DrawImageComponent struct {
	Source *ebiten.Image
}

func NewDrawImageComponent() *DrawImageComponent {
	return &DrawImageComponent{}
}

func (c *DrawImageComponent) Draw(screen *ebiten.Image, transformer utility.Transformer) {
	if c.Source == nil {
		return
	}

	location := transformer.GetLocation()
	scale := transformer.GetScale()

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(scale.X, scale.Y)
	o.GeoM.Rotate(transformer.GetRotation())
	o.GeoM.Translate(location.X, location.Y)
	screen.DrawImage(c.Source, o)
}
