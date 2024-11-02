package ebitenhelper

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Landscape struct {
	Image *DrawFullScreenComponent
}

func NewLandscape() *Landscape {
	return &Landscape{
		Image: NewDrawFullScreenComponent(),
	}
}

func (l *Landscape) Draw(screen *ebiten.Image) {
	l.Image.Draw(screen)
}
