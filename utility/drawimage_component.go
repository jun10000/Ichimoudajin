package utility

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawImageComponent struct {
	parent *Actor
	Image  *ebiten.Image
}

func NewDrawImageComponent(parentActor *Actor) *DrawImageComponent {
	if parentActor == nil {
		log.Fatal("Failed to create DrawImageComponent")
	}

	return &DrawImageComponent{
		parent: parentActor,
	}
}

func (c *DrawImageComponent) Draw(screen *ebiten.Image) {
	if c.Image == nil {
		return
	}

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(c.parent.Scale.X, c.parent.Scale.Y)
	o.GeoM.Rotate(c.parent.Rotation.Get())
	o.GeoM.Translate(c.parent.Location.X, c.parent.Location.Y)
	screen.DrawImage(c.Image, o)
}
