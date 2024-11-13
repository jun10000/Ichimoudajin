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
