package utility

import "github.com/hajimehoshi/ebiten/v2"

type StaticLocator interface {
	GetLocation() Vector
}

type Locator interface {
	StaticLocator
	SetLocation(value Vector)
}

type StaticRotator interface {
	GetRotation() float64
}

type Rotator interface {
	StaticRotator
	SetRotation(value float64)
}

type StaticScaler interface {
	GetScale() Vector
}

type Scaler interface {
	StaticScaler
	SetScale(value Vector)
}

type StaticTransformer interface {
	StaticLocator
	StaticRotator
	StaticScaler
}

type Transformer interface {
	Locator
	Rotator
	Scaler
}

type ColliderBase interface {
	GetMainColliderBounds() Bounder
	GetColliderBounds() [9]Bounder
}

type Collider interface {
	ColliderBase
	StaticTransformer
}

type MovableCollider interface {
	ColliderBase
	Transformer
	AddInput(normal Vector, scale float64)
	AddLocation(offset Vector) *TraceResult
}

type ColliderComparable interface {
	Collider
	comparable
}

type InputReceiver interface {
	ReceiveKeyInput(key ebiten.Key, state PressState)
	ReceiveMouseButtonInput(button ebiten.MouseButton, state PressState, pos Point)
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
