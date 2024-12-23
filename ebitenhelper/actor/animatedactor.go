package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type AnimatedActor struct {
	*utility.Transform
	*component.DrawAnimationComponent
}

func NewAnimatedActor() *AnimatedActor {
	a := &AnimatedActor{
		Transform: utility.DefaultTransform(),
	}

	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	return a
}
