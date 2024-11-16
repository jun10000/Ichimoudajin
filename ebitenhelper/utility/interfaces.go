package utility

import "github.com/hajimehoshi/ebiten/v2"

type Drawer interface {
	Draw(screen *ebiten.Image)
}

type Ticker interface {
	Tick()
}

type KeyReceiver interface {
	ReceivePressedKey(key ebiten.Key)
	ReceiveReleasedKey(key ebiten.Key)
	ReceivePressingKey(key ebiten.Key)
}

type Collider interface {
	GetBounds() RectangleF
}

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

type Mover interface {
	Transformer
	Collider
}
