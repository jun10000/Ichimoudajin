package utility

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Actor struct {
	Transform
	Image *DrawImageComponent
}

func NewActor() *Actor {
	actor := &Actor{
		Transform: DefaultTransform(),
	}
	actor.Image = NewDrawImageComponent(actor)
	return actor
}

func (a *Actor) Draw(screen *ebiten.Image) {
	a.Image.Draw(screen)
}
