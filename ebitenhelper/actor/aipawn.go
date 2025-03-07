package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type AIPawn struct {
	utility.Transform
	*component.MovementComponent
	*component.DrawAnimationComponent
	*component.AIControllerComponent
	*component.ColliderComponent[*utility.CircleF]
}

func NewAIPawn() *AIPawn {
	a := &AIPawn{
		Transform: utility.DefaultTransform(),
	}

	a.MovementComponent = component.NewMovementComponent(a)
	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	a.AIControllerComponent = component.NewAIControllerComponent(a)
	a.ColliderComponent = component.NewColliderComponent(a.GetCircleBounds)
	return a
}

func (a *AIPawn) SetLocation(value utility.Vector) {
	a.Transform.SetLocation(value)
	a.UpdateColliderBounds()
}

func (a *AIPawn) SetScale(value utility.Vector) {
	a.Transform.SetScale(value)
	a.UpdateColliderBounds()
}

func (a *AIPawn) Tick() {
	a.MovementComponent.Tick()
	a.DrawAnimationComponent.Tick()
}
