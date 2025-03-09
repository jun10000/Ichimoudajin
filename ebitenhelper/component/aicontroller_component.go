package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type AIControllerComponent struct {
	parent utility.MovableCollider
}

func NewAIControllerComponent(parent utility.MovableCollider) *AIControllerComponent {
	return &AIControllerComponent{
		parent: parent,
	}
}

func (a *AIControllerComponent) AITick() {
	l := utility.GetLevel()
	if len(l.Players) == 0 {
		return
	}

	p := l.Players[0]
	l.AIMove(a.parent, p)
}
