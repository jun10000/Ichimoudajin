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

func (a *AIPawn) Tick() {
	a.MovementComponent.Tick()
	a.DrawAnimationComponent.Tick()
}
