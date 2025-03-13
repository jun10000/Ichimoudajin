package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Actor struct {
	*utility.StaticTransform
	*component.DrawImageComponent
}

func NewActor(sTransform *utility.StaticTransform) *Actor {
	a := &Actor{
		StaticTransform: sTransform,
	}

	a.DrawImageComponent = component.NewDrawImageComponent(a)
	return a
}
