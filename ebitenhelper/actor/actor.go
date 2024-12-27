package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Actor struct {
	utility.Transform
	*component.DrawImageComponent
}

func NewActor() *Actor {
	a := &Actor{
		Transform: utility.DefaultTransform(),
	}

	a.DrawImageComponent = component.NewDrawImageComponent(a)
	return a
}
