package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Destroyer struct {
	IsShow bool
	Circle *utility.CircleF
}

func NewDestroyer() *Destroyer {
	return &Destroyer{
		Circle: utility.NewCircleF(0, 0, 0),
	}
}

func (a *Destroyer) Draw(screen *ebiten.Image) {
	if a.IsShow {
		a.Circle.Draw(screen, utility.DebugColorBlue, true)
	}
}
