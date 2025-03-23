package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Pawn struct {
	*component.MovementComponent
	*component.DrawAnimationComponent
	*component.ControllerComponent
	*component.ColliderComponent[*utility.CircleF]

	destroyer *Destroyer
}

func (a *Pawn) Children() []any {
	return []any{
		a.destroyer,
	}
}

func (a *Pawn) ReceiveMouseButtonInput(button ebiten.MouseButton, state utility.PressState, pos utility.Point) {
	a.ControllerComponent.ReceiveMouseButtonInput(button, state, pos)
	if button != ebiten.MouseButtonLeft {
		return
	}

	switch state {
	case utility.PressStatePressed:
		a.destroyer.Start(pos.ToVector())
	case utility.PressStateReleased:
		a.destroyer.Finish()
	}
}

func (a *Pawn) Tick() {
	a.MovementComponent.Tick()
	a.DrawAnimationComponent.Tick()
}
