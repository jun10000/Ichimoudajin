package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Actor struct {
	*utility.StaticTransform
	*component.DrawImageComponent
}

func NewActor(location utility.Vector, rotation float64, scale utility.Vector) *Actor {
	a := &Actor{}
	a.StaticTransform = utility.NewStaticTransform(location, rotation, scale)
	a.DrawImageComponent = component.NewDrawImageComponent(a)
	return a
}
