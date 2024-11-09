package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Actor struct {
	utility.Transform
	Image *component.DrawImageComponent
}

func NewActor() *Actor {
	return &Actor{
		Transform: utility.DefaultTransform(),
		Image:     component.NewDrawImageComponent(),
	}
}

func (a *Actor) Draw(screen *ebiten.Image) {
	a.Image.Draw(screen, a.Transform)
}
