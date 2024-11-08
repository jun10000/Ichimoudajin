package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
)

type Landscape struct {
	Image *component.DrawFullScreenComponent
}

func NewLandscape() *Landscape {
	return &Landscape{
		Image: component.NewDrawFullScreenComponent(),
	}
}

func (l *Landscape) Draw(screen *ebiten.Image) {
	l.Image.Draw(screen)
}
