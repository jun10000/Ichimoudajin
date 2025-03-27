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

type Bounder interface {
	ToCircle() *CircleF
	CenterLocation() Vector
	Offset(x, y float64, output Bounder) Bounder
	IntersectTo(target Bounder) (result bool, normal *Vector)
	IntersectFromRectangle(src *RectangleF) (result bool, normal *Vector)
	IntersectFromCircle(src *CircleF) (result bool, normal *Vector)
}

type Actor interface {
	GetName() string
}

type ColliderBase interface {
	Actor
	UpdateBounds()
	EnableBounds()
	DisableBounds()
	GetRealFirstBounds() Bounder
	GetRealBounds() []Bounder
	GetFirstBounds() Bounder
	GetBounds() []Bounder
	ReceiveHit(result *TraceResult[Collider])
}

type Collider interface {
	ColliderBase
	StaticTransformer
}

type MovableCollider interface {
	ColliderBase
	Transformer
	AddInput(normal Vector, scale float64)
	AddLocation(offset Vector) *TraceResult[Collider]
}

type ColliderComparable interface {
	Collider
	comparable
}

type InputReceiver interface {
	Actor
	ReceiveKeyInput(key ebiten.Key, state PressState)
	ReceiveMouseButtonInput(button ebiten.MouseButton, state PressState, pos Point)
	ReceiveGamepadButtonInput(id ebiten.GamepadID, button ebiten.StandardGamepadButton, state PressState)
	ReceiveGamepadAxisInput(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis, value float64)
}

type Player interface {
	InputReceiver
	MovableCollider
}

type BeginPlayer interface {
	Actor
	BeginPlay()
}

type EndPlayer interface {
	Actor
	EndPlay()
}

type AITicker interface {
	Actor
	AITick()
}

type Ticker interface {
	Actor
	Tick()
}

type Drawer interface {
	Actor
	GetVisibility() bool
	SetVisibility(isVisible bool)
	Draw(screen *ebiten.Image)
}

type ZSpecifiedDrawer interface {
	Drawer
	ZOrder() int
}
