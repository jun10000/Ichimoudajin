package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type AIPawn struct {
	*component.MovementComponent
	*component.DrawAnimationComponent
	*component.AIControllerComponent
	*component.ColliderComponent[*utility.CircleF]
}

func NewAIPawn(location utility.Vector, rotation float64, scale utility.Vector) *AIPawn {
	t := utility.NewTransform(location, rotation, scale)

	a := &AIPawn{}
	a.MovementComponent = component.NewMovementComponent(a)
	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	a.AIControllerComponent = component.NewAIControllerComponent(a)
	a.ColliderComponent = component.NewColliderComponent(t, a.GetCircleBounds)
	a.UpdateColliderBounds()
	return a
}

func (a *AIPawn) Tick() {
	a.MovementComponent.Tick()
	a.DrawAnimationComponent.Tick()
}
