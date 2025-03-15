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

func NewPawn(location utility.Vector, rotation float64, scale utility.Vector) *Pawn {
	t := utility.NewTransform(location, rotation, scale)

	a := &Pawn{}
	a.MovementComponent = component.NewMovementComponent(a)
	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	a.ControllerComponent = component.NewControllerComponent(a)
	a.ColliderComponent = component.NewColliderComponent(t, a.GetCircleBounds)

	a.destroyer = NewDestroyer()

	a.UpdateColliderBounds()
	return a
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
		a.destroyer.IsShow = true
		a.destroyer.Circle.OrgX = float64(pos.X)
		a.destroyer.Circle.OrgY = float64(pos.Y)
		a.destroyer.Circle.Radius = 0
	case utility.PressStatePressing:
		a.destroyer.Circle.Radius += 5
		if a.destroyer.Circle.Radius > 100 {
			a.destroyer.Circle.Radius = 100
		}
	case utility.PressStateReleased:
		a.destroyer.IsShow = false
	}
}

func (a *Pawn) Tick() {
	a.MovementComponent.Tick()
	a.DrawAnimationComponent.Tick()
}
