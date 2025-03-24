package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type ImageActor struct {
	*utility.StaticTransform
	*component.DrawImageComponent
}

func (g ActorGeneratorStruct) NewImageActor(location utility.Vector, rotation float64, scale utility.Vector) *ImageActor {
	a := &ImageActor{}
	a.StaticTransform = utility.NewStaticTransform(location, rotation, scale)
	a.DrawImageComponent = component.NewDrawImageComponent(a)
	return a
}
