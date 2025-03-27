package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type AnimatedActor struct {
	*utility.StaticTransform
	*component.DrawAnimationCom
}

func (g ActorGeneratorStruct) NewAnimatedActor(location utility.Vector, rotation float64, scale utility.Vector) *AnimatedActor {
	a := &AnimatedActor{}
	a.StaticTransform = utility.NewStaticTransform(location, rotation, scale)
	a.DrawAnimationCom = component.NewDrawAnimationCom(a)
	return a
}
