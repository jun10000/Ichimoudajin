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

func (g ActorGeneratorStruct) NewPawn(location utility.Vector, rotation float64, scale utility.Vector) *Pawn {
	t := utility.NewTransform(location, rotation, scale)

	a := &Pawn{}
	a.MovementComponent = component.NewMovementComponent(a)
	a.DrawAnimationComponent = component.NewDrawAnimationComponent(a)
	a.ControllerComponent = component.NewControllerComponent(a)
	a.ColliderComponent = component.NewColliderComponent(t, a.GetCircleBounds)

	a.destroyer = g.NewDestroyer()

	a.UpdateBounds()
	return a
}

// NewPawn1 creates another version of Pawn
func (g ActorGeneratorStruct) NewPawn1(location utility.Vector, rotation float64, scale utility.Vector) *Pawn {
	a := g.NewPawn(location, rotation, scale)
	a.Image = utility.GetImageFromFileP("images/ぴぽやキャラチップ32出力素材/現代系/女_スーツ1.png")
	a.MaxSpeed = 200
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
		a.destroyer.Start(pos.ToVector())
	case utility.PressStateReleased:
		a.destroyer.Finish()
	}
}

func (a *Pawn) Tick() {
	a.MovementComponent.Tick()
	a.DrawAnimationComponent.Tick()
}
