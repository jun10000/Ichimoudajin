package component

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type ControllerComponent struct {
	DeadZone float64
	parent   utility.Mover
}

func NewControllerComponent(parent utility.Mover) *ControllerComponent {
	return &ControllerComponent{
		DeadZone: 0.2,
		parent:   parent,
	}
}

func (c *ControllerComponent) ReceiveKeyInput(key ebiten.Key, state utility.PressState) {
	switch key {
	case ebiten.KeyEscape:
		os.Exit(0)
	}

	if state != utility.PressStatePressing {
		return
	}

	switch key {
	case ebiten.KeyUp:
		c.parent.AddInput(utility.UpVector(), 1)
	case ebiten.KeyDown:
		c.parent.AddInput(utility.DownVector(), 1)
	case ebiten.KeyLeft:
		c.parent.AddInput(utility.LeftVector(), 1)
	case ebiten.KeyRight:
		c.parent.AddInput(utility.RightVector(), 1)
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
		c.parent.AddInput(utility.RightVector(), value)
	case ebiten.StandardGamepadAxisLeftStickVertical:
		c.parent.AddInput(utility.DownVector(), value)
	}
}
