package actor

import (
	"math"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Pawn struct {
	*utility.Transform
	*component.MovementComponent
	*component.DrawAnimationComponent
	*component.ControllerComponent
}

func NewPawn() *Pawn {
	a := &Pawn{
		Transform: utility.DefaultTransform(),
	}

	a.MovementComponent = component.NewMovementComponent(a)
	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	a.ControllerComponent = component.NewControllerComponent(a.MovementComponent)
	return a
}

func (p *Pawn) Tick() {
	p.MovementComponent.Tick()
	p.DrawAnimationComponent.Tick()
}

func (p *Pawn) GetBounds() utility.Bounder {
	hs := p.FrameSize.ToVector().DivF(2).Mul(p.GetScale())
	return utility.NewCircleF(p.GetLocation().Add(hs), math.Max(hs.X, hs.Y))
}
