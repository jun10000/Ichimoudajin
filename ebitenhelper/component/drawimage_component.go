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

func (c *DrawImageComponent) Draw(screen *ebiten.Image, transform utility.Transform) {
	if c.Source == nil {
		return
	}

	location := transform.GetLocation()
	scale := transform.GetScale()

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(scale.X, scale.Y)
	o.GeoM.Rotate(transform.GetRotation())
	o.GeoM.Translate(location.X, location.Y)
	screen.DrawImage(c.Source, o)
}
