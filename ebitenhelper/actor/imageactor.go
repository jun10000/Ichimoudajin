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

func (g ActorGeneratorStruct) NewImageActor(options *NewActorOptions) *ImageActor {
	a := &ImageActor{}
	a.ActorCom = component.NewActorCom(options.Name)
	a.StaticTransform = utility.NewStaticTransform(options.Location, options.Rotation, options.Scale)
	a.DrawImageCom = component.NewDrawImageCom(a, options.IsVisible)
	return a
}
