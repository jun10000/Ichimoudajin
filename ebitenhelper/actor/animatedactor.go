package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type AnimatedActor struct {
	*component.ActorCom
	*utility.StaticTransform
	*component.DrawAnimationCom
}

func (g ActorGeneratorStruct) NewAnimatedActor(options *NewActorOptions) *AnimatedActor {
	a := &AnimatedActor{}
	a.ActorCom = component.NewActorCom(options.Name)
	a.StaticTransform = utility.NewStaticTransform(options.Location, options.Rotation, options.Scale)
	a.DrawAnimationCom = component.NewDrawAnimationCom(a, options.IsVisible)
	return a
}
