package component

import (
	"os"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type ControllerCom struct {
	parent utility.MovableCollider
}

func NewControllerCom(parent utility.MovableCollider) *ControllerCom {
	return &ControllerCom{
		parent: parent,
	}
}

func (c *ControllerCom) ReceiveKeyInput(key ebiten.Key, state utility.PressState) {
	switch key {
	case ebiten.KeyEscape:
		if runtime.GOOS != "js" {
			os.Exit(0)
		}
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

func (c *ControllerCom) ReceiveMouseButtonInput(button ebiten.MouseButton, state utility.PressState, pos utility.Point) {
}

func (c *ControllerCom) ReceiveGamepadButtonInput(id ebiten.GamepadID, button ebiten.StandardGamepadButton, state utility.PressState) {
}

func (c *ControllerCom) ReceiveGamepadAxisInput(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis, value float64) {
	if -utility.GamepadDeadZone < value && value < utility.GamepadDeadZone {
		return
	}

	switch axis {
	case ebiten.StandardGamepadAxisLeftStickHorizontal:
		c.parent.AddInput(utility.RightVector(), value)
	case ebiten.StandardGamepadAxisLeftStickVertical:
		c.parent.AddInput(utility.DownVector(), value)
	}
}
