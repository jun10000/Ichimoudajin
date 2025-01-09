package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type AIControllerComponent struct {
	parent utility.Transformer
	target *MovementComponent
}

func NewAIControllerComponent(parent utility.Transformer, target *MovementComponent) *AIControllerComponent {
	return &AIControllerComponent{
		parent: parent,
		target: target,
	}
}

func (a *AIControllerComponent) AITick() {
	pl := utility.GetLevel().InputReceivers[0].GetLocation()
	el := a.parent.GetLocation()
	a.target.AddInput(pl.Sub(el), 1)
}
