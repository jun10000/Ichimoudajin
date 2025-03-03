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
	for r := range l.InputReceivers {
		if p, ok := r.(utility.Collider); ok {
			l.AIMove(a.parent, p)
			return
		}
	}
}
