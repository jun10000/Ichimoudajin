package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Pawn struct {
	*component.MovementComponent
	*component.DrawAnimationComponent
	*component.ControllerComponent
	*component.ColliderComponent[*utility.CircleF]
}

func NewPawn() *Pawn {
	a := &Pawn{}

	a.MovementComponent = component.NewMovementComponent(a)
	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	a.ControllerComponent = component.NewControllerComponent(a)
	a.ColliderComponent = component.NewColliderComponent(a.GetCircleBounds)
	return a
}

func (a *Pawn) Tick() {
	a.MovementComponent.Tick()
	a.DrawAnimationComponent.Tick()
}
