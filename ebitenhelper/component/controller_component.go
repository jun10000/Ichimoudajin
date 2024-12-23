package component

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type ControllerComponent struct {
	DeadZone float64
	target   *MovementComponent
}

func NewControllerComponent(target *MovementComponent) *ControllerComponent {
	return &ControllerComponent{
		DeadZone: 0.2,
		target:   target,
	}
}

func (c *ControllerComponent) ReceiveKeyInput(key ebiten.Key, state utility.PressState) {
	switch key {
	case ebiten.KeyEscape:
		os.Exit(0)
	}

	if state != utility.PressState_Pressing {
		return
	}

	switch key {
	case ebiten.KeyUp:
		c.target.AddInput(utility.NewVector(0, -1), 1)
	case ebiten.KeyDown:
		c.target.AddInput(utility.NewVector(0, 1), 1)
	case ebiten.KeyLeft:
		c.target.AddInput(utility.NewVector(-1, 0), 1)
	case ebiten.KeyRight:
		c.target.AddInput(utility.NewVector(1, 0), 1)
	}
}

func (c *ControllerComponent) ReceiveButtonInput(id ebiten.GamepadID, button ebiten.StandardGamepadButton, state utility.PressState) {
}

func (c *ControllerComponent) ReceiveAxisInput(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis, value float64) {
	if -c.DeadZone < value && value < c.DeadZone {
		return
	}

	switch axis {
	case ebiten.StandardGamepadAxisLeftStickHorizontal:
		c.target.AddInput(utility.NewVector(1, 0), value)
	case ebiten.StandardGamepadAxisLeftStickVertical:
		c.target.AddInput(utility.NewVector(0, 1), value)
	}
}
