package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type AnimatedActor struct {
	*utility.StaticTransform
	*component.DrawAnimationComponent
}

func NewAnimatedActor(sTransform *utility.StaticTransform) *AnimatedActor {
	a := &AnimatedActor{
		StaticTransform: sTransform,
	}

	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	return a
}
