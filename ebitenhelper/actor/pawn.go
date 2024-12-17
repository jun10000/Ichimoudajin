package actor

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Pawn struct {
	*utility.Transform
	Movement  *component.MovementComponent
	Animation *component.DrawAnimationComponent
}

func NewPawn() *Pawn {
	return &Pawn{
		Transform: utility.DefaultTransform(),
		Movement:  component.NewMovementComponent(),
		Animation: component.NewDrawAnimationComponent(),
	}
}

func (p *Pawn) ReceivePressedKey(key ebiten.Key) {
}

func (p *Pawn) ReceiveReleasedKey(key ebiten.Key) {
}

func (p *Pawn) ReceivePressingKey(key ebiten.Key) {
	switch key {
	case ebiten.KeyUp:
		p.Movement.AddInput(utility.NewVector(0, -1), 1)
	case ebiten.KeyDown:
		p.Movement.AddInput(utility.NewVector(0, 1), 1)
	case ebiten.KeyLeft:
		p.Movement.AddInput(utility.NewVector(-1, 0), 1)
	case ebiten.KeyRight:
		p.Movement.AddInput(utility.NewVector(1, 0), 1)
	}
}

func (p *Pawn) ReceivePressedButton(id ebiten.GamepadID, button ebiten.StandardGamepadButton) {
}

func (p *Pawn) ReceiveReleasedButton(id ebiten.GamepadID, button ebiten.StandardGamepadButton) {
}

func (p *Pawn) ReceiveAxisValue(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis, value float64) {
	// Use DeadZone
	if -0.2 < value && value < 0.2 {
		value = 0
	}

	switch axis {
	case ebiten.StandardGamepadAxisLeftStickHorizontal:
		p.Movement.AddInput(utility.NewVector(1, 0), value)
	case ebiten.StandardGamepadAxisLeftStickVertical:
		p.Movement.AddInput(utility.NewVector(0, 1), value)
	}
}

func (p *Pawn) Tick() {
	p.Movement.Tick(p)
	p.Animation.Tick()
}

func (p *Pawn) Draw(screen *ebiten.Image) {
	p.Animation.Draw(screen, p)
}

func (p *Pawn) GetBounds() utility.Bounder {
	hs := p.Animation.FrameSize.ToVector().DivF(2)
	return utility.NewCircleF(p.GetLocation().Add(hs), math.Max(hs.X, hs.Y))
}
