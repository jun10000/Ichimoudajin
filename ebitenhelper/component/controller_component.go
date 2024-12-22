package component

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type ControllerComponent struct {
	DeadZone float64
}

func NewControllerComponent() *ControllerComponent {
	return &ControllerComponent{
		DeadZone: 0.2,
	}
}

func (c *ControllerComponent) ReceiveKeyInput(movement *MovementComponent, key ebiten.Key, state utility.PressState) {
	switch key {
	case ebiten.KeyEscape:
		os.Exit(0)
	}

	if state != utility.PressState_Pressing {
		return
	}

	switch key {
	case ebiten.KeyUp:
		movement.AddInput(utility.NewVector(0, -1), 1)
	case ebiten.KeyDown:
		movement.AddInput(utility.NewVector(0, 1), 1)
	case ebiten.KeyLeft:
		movement.AddInput(utility.NewVector(-1, 0), 1)
	case ebiten.KeyRight:
		movement.AddInput(utility.NewVector(1, 0), 1)
	}
}

func (c *ControllerComponent) ReceiveButtonInput(movement *MovementComponent, id ebiten.GamepadID, button ebiten.StandardGamepadButton, state utility.PressState) {
}

func (c *ControllerComponent) ReceiveAxisInput(movement *MovementComponent, id ebiten.GamepadID, axis ebiten.StandardGamepadAxis, value float64) {
	if -c.DeadZone < value && value < c.DeadZone {
		return
	}

	switch axis {
	case ebiten.StandardGamepadAxisLeftStickHorizontal:
		movement.AddInput(utility.NewVector(1, 0), value)
	case ebiten.StandardGamepadAxisLeftStickVertical:
		movement.AddInput(utility.NewVector(0, 1), value)
	}
}
