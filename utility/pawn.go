package utility

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Pawn struct {
	*Actor
	Movement *MovementComponent

	currentTickIndex int
}

func NewPawn() *Pawn {
	actor := NewActor()
	actor.Scale = NewVector(2, 2)
	pawn := &Pawn{
		Actor: actor,
	}
	pawn.Movement = NewMovementComponent(pawn)

	return pawn
}

func (p *Pawn) ReceivePressedKey(key ebiten.Key) {
}

func (p *Pawn) ReceiveReleasedKey(key ebiten.Key) {
}

func (p *Pawn) ReceivePressingKey(key ebiten.Key) {
	switch key {
	case ebiten.KeyUp:
		p.Movement.AddInput(NewVector(0, -1), 1)
	case ebiten.KeyDown:
		p.Movement.AddInput(NewVector(0, 1), 1)
	case ebiten.KeyLeft:
		p.Movement.AddInput(NewVector(-1, 0), 1)
	case ebiten.KeyRight:
		p.Movement.AddInput(NewVector(1, 0), 1)
	}
}

func (p *Pawn) Tick() {
	p.currentTickIndex++
	p.Movement.Tick()
}

func (p *Pawn) Draw(screen *ebiten.Image) {
	tipframecount := 4
	tipspeed := 4
	tipsize := NewPoint(32, 32)

	tipindex := (p.currentTickIndex * tipspeed / TickCount) % tipframecount
	if tipindex == 3 {
		tipindex = 1
	}

	tipdirection := 3
	switch {
	case p.Rotation.Get() < math.Pi*-3/4:
		tipdirection = 3
	case p.Rotation.Get() < math.Pi*-1/4:
		tipdirection = 2
	case p.Rotation.Get() < math.Pi*1/4:
		tipdirection = 0
	case p.Rotation.Get() < math.Pi*3/4:
		tipdirection = 1
	}

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(p.Scale.X, p.Scale.Y)
	o.GeoM.Translate(p.Location.X, p.Location.Y)
	screen.DrawImage(p.Image.SubImage(image.Rect(tipindex*tipsize.X, tipdirection*tipsize.Y, tipindex*tipsize.X+tipsize.X, tipdirection*tipsize.Y+tipsize.Y)).(*ebiten.Image), o)
}
