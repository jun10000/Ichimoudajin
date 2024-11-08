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

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(transform.Scale.X, transform.Scale.Y)
	o.GeoM.Rotate(transform.Rotation.Get())
	o.GeoM.Translate(transform.Location.X, transform.Location.Y)
	screen.DrawImage(c.Source, o)
}
