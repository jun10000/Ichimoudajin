package ebitenhelper

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type DrawImageComponent struct {
	parent *Actor
	Source *ebiten.Image
}

func NewDrawImageComponent(parentActor *Actor) *DrawImageComponent {
	return &DrawImageComponent{
		parent: parentActor,
	}
}

func (c *DrawImageComponent) Draw(screen *ebiten.Image) {
	if c.Source == nil {
		return
	}

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(c.parent.Scale.X, c.parent.Scale.Y)
	o.GeoM.Rotate(c.parent.Rotation.Get())
	o.GeoM.Translate(c.parent.Location.X, c.parent.Location.Y)
	screen.DrawImage(c.Source, o)
}
