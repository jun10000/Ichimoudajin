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
	Movement *MovementComponent
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
		Movement: NewMovementComponent(),
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
		a.Movement.AddInput(NewVector(0, -1), 1)
	case ebiten.KeyDown:
		a.Movement.AddInput(NewVector(0, 1), 1)
	case ebiten.KeyLeft:
		a.Movement.AddInput(NewVector(-1, 0), 1)
	case ebiten.KeyRight:
		a.Movement.AddInput(NewVector(1, 0), 1)
	}
}

func (a *Actor) Event_Tick() {
	a.Movement.Event_Tick(a)
}
