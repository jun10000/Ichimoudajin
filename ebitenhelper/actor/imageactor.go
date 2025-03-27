package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type ImageActor struct {
	*component.ActorCom
	*utility.StaticTransform
	*component.DrawImageCom
}

func (g ActorGeneratorStruct) NewImageActor(name string, location utility.Vector, rotation float64, scale utility.Vector) *ImageActor {
	a := &ImageActor{}
	a.ActorCom = component.NewActorCom(name)
	a.StaticTransform = utility.NewStaticTransform(location, rotation, scale)
	a.DrawImageCom = component.NewDrawImageCom(a)
	return a
}
