package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Pawn struct {
	*utility.Transform
	Movement   *component.MovementComponent
	Controller *component.ControllerComponent
	Animation  *component.DrawAnimationComponent
}

func NewPawn() *Pawn {
	return &Pawn{
		Transform:  utility.DefaultTransform(),
		Movement:   component.NewMovementComponent(),
		Controller: component.NewControllerComponent(),
		Animation:  component.NewDrawAnimationComponent(),
	}
}

func (p *Pawn) ReceiveKeyInput(key ebiten.Key, state utility.PressState) {
	p.Controller.ReceiveKeyInput(p.Movement, key, state)
}

func (p *Pawn) ReceiveButtonInput(id ebiten.GamepadID, button ebiten.StandardGamepadButton, state utility.PressState) {
	p.Controller.ReceiveButtonInput(p.Movement, id, button, state)
}

func (p *Pawn) ReceiveAxisInput(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis, value float64) {
	p.Controller.ReceiveAxisInput(p.Movement, id, axis, value)
}

func (p *Pawn) Tick() {
	p.Movement.Tick(p)
	p.Animation.Tick()
}

func (p *Pawn) Draw(screen *ebiten.Image) {
	p.Animation.Draw(screen, p)
}

func (p *Pawn) GetBounds() utility.Bounder {
	return p.Animation.GetCircleBounds(p)
}
