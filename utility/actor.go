package utility

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Actor struct {
	Image    *ebiten.Image
	Location Vector
	Rotation float64
	Scale    Vector
	MaxSpeed float64
}

func NewActor(imagefile string) *Actor {
	image, _, err := ebitenutil.NewImageFromFile(imagefile)
	if err != nil {
		log.Fatal(err)
	}

	return &Actor{
		Image:    image,
		Location: NewVector(0, 0),
		Rotation: 0,
		Scale:    NewVector(1, 1),
		MaxSpeed: 5,
	}
}

func (a *Actor) Draw(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(a.Scale.X, a.Scale.Y)
	o.GeoM.Rotate(a.Rotation)
	o.GeoM.Translate(a.Location.X, a.Location.Y)
	screen.DrawImage(a.Image, o)
}

func (a *Actor) Event_KeyPressed(key ebiten.Key) {
}

func (a *Actor) Event_KeyReleased(key ebiten.Key) {
}

func (a *Actor) Event_KeyPressing(key ebiten.Key) {
	switch key {
	case ebiten.KeyUp:
		a.Location.Y -= a.MaxSpeed
	case ebiten.KeyDown:
		a.Location.Y += a.MaxSpeed
	case ebiten.KeyLeft:
		a.Location.X -= a.MaxSpeed
	case ebiten.KeyRight:
		a.Location.X += a.MaxSpeed
	}
}
