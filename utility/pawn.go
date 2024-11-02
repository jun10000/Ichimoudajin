package utility

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Pawn struct {
	*Actor
	Movement  *MovementComponent
	Animation *AnimationComponent
}

func NewPawn() *Pawn {
	actor := NewActor()
	actor.Scale = NewVector(2, 2)
	pawn := &Pawn{
		Actor: actor,
	}
	pawn.Movement = NewMovementComponent(pawn)
	pawn.Animation = NewAnimationComponent(pawn)

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
	p.Movement.Tick()
	p.Animation.Tick()
}

func (p *Pawn) Draw(screen *ebiten.Image) {
	p.Animation.Draw(screen)
}
