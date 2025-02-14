package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type AIControllerComponent struct {
	parent utility.Mover
}

func NewAIControllerComponent(parent utility.Mover) *AIControllerComponent {
	return &AIControllerComponent{
		parent: parent,
	}
}

func (a *AIControllerComponent) AITick() {
	l := utility.GetLevel()
	p := l.InputReceivers[0].(utility.Collider)
	l.AIMove(a.parent, p)
}
