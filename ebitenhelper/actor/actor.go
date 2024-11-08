package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Actor struct {
	utility.Transform
	ImageCom *component.DrawImageComponent
}

func NewActor() *Actor {
	actor := &Actor{
		Transform: utility.DefaultTransform(),
		ImageCom:  component.NewDrawImageComponent(),
	}
	return actor
}

func (a *Actor) Draw(screen *ebiten.Image) {
	a.ImageCom.Draw(screen, a.Transform)
}
