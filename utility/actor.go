package utility

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Actor struct {
	Location Vector
	Rotation Rotator
	Scale    Vector
	Image    *ebiten.Image
}

func NewActor() *Actor {
	return &Actor{
		Location: NewVector(0, 0),
		Rotation: NewRotator(0),
		Scale:    NewVector(1, 1),
	}
}

func (a *Actor) Draw(screen *ebiten.Image) {
	if a.Image == nil {
		return
	}

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(a.Scale.X, a.Scale.Y)
	o.GeoM.Rotate(a.Rotation.Get())
	o.GeoM.Translate(a.Location.X, a.Location.Y)
	screen.DrawImage(a.Image, o)
}
