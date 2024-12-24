package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type AIPawn struct {
	*utility.Transform
	*component.MovementComponent
	*component.DrawAnimationComponent
}

func NewAIPawn() *AIPawn {
	a := &AIPawn{
		Transform: utility.DefaultTransform(),
	}

	a.MovementComponent = component.NewMovementComponent(a)
	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	return a
}

func (a *AIPawn) Tick() {
	a.MovementComponent.Tick()
	a.DrawAnimationComponent.Tick()
}

func (a *AIPawn) GetColliderBounds() utility.Bounder {
	return a.GetCircleBounds()
}
