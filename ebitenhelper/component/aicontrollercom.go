package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type AIControllerCom struct {
	parent utility.MovableCollider
}

func NewAIControllerCom(parent utility.MovableCollider) *AIControllerCom {
	return &AIControllerCom{
		parent: parent,
	}
}

func (a *AIControllerCom) AITick() {
	l := utility.GetLevel()
	if len(l.Players) == 0 {
		return
	}

	p := l.Players[0]
	l.AIMove(a.parent, p)
}
