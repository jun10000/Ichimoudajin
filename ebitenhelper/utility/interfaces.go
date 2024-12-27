package utility

import "github.com/hajimehoshi/ebiten/v2"

type Locator interface {
	GetLocation() Vector
	SetLocation(value Vector)
	AddLocation(value Vector)
}

type Rotator interface {
	GetRotation() float64
	SetRotation(value float64)
}

type Scaler interface {
	GetScale() Vector
	SetScale(value Vector)
}

type Transformer interface {
	Locator
	Rotator
	Scaler
}

type Collider interface {
	Transformer
	GetColliderBounds() Bounder
}

type InputReceiver interface {
	Transformer
	ReceiveKeyInput(key ebiten.Key, state PressState)
	ReceiveButtonInput(id ebiten.GamepadID, button ebiten.StandardGamepadButton, state PressState)
	ReceiveAxisInput(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis, value float64)
}

type AITicker interface {
	AITick()
}

type Ticker interface {
	Tick()
}

type Drawer interface {
	Draw(screen *ebiten.Image)
}

type Bounder interface {
	BoundingBox() RectangleF
	Offset(value Vector) Bounder
}
