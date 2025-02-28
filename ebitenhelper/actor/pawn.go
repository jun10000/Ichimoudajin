package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Pawn struct {
	utility.Transform
	*component.MovementComponent
	*component.DrawAnimationComponent
	*component.ControllerComponent
	*component.CircleColliderComponent
}

func NewPawn() *Pawn {
	a := &Pawn{
		Transform: utility.DefaultTransform(),
	}

	a.MovementComponent = component.NewMovementComponent(a)
	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	a.ControllerComponent = component.NewControllerComponent(a)
	a.CircleColliderComponent = component.NewCircleColliderComponent(a.GetCircleBounds)
	return a
}

func (p *Pawn) Tick() {
	p.MovementComponent.Tick()
	p.DrawAnimationComponent.Tick()
}
