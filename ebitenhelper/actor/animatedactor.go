package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type AnimatedActor struct {
	utility.Transform
	Animation *component.DrawAnimationComponent
}

func NewAnimatedActor() *AnimatedActor {
	return &AnimatedActor{
		Transform: utility.DefaultTransform(),
		Animation: component.NewDrawAnimationComponent(),
	}
}

func (a *AnimatedActor) Tick() {
	a.Animation.Tick()
}

func (a *AnimatedActor) Draw(screen *ebiten.Image) {
	a.Animation.Draw(screen, a.Transform)
}
