package utility

import "github.com/hajimehoshi/ebiten/v2"

type Locator interface {
	GetLocation() Vector
	SetLocation(value Vector)
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
	GetMainColliderBounds() Bounder
	GetColliderBounds() [9]Bounder
}

type ColliderComparable interface {
	Collider
	comparable
}

type MovableCollider interface {
	Collider
	AddInput(normal Vector, scale float64)
	AddLocation(offset Vector) *TraceResult
}

type InputReceiver interface {
	ReceiveKeyInput(key ebiten.Key, state PressState)
	ReceiveButtonInput(id ebiten.GamepadID, button ebiten.StandardGamepadButton, state PressState)
	ReceiveAxisInput(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis, value float64)
}

type Player interface {
	InputReceiver
	MovableCollider
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
	CenterLocation() Vector
	Offset(x, y float64, output Bounder) Bounder
	IntersectTo(target Bounder) (result bool, normal *Vector)
	IntersectFromRectangle(src *RectangleF) (result bool, normal *Vector)
	IntersectFromCircle(src *CircleF) (result bool, normal *Vector)
}
