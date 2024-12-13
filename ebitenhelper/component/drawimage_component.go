package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type DrawImageComponent struct {
	Source *ebiten.Image
}

func NewDrawImageComponent() *DrawImageComponent {
	return &DrawImageComponent{}
}

func (c *DrawImageComponent) Draw(screen *ebiten.Image, transform utility.Transformer) {
	utility.DrawImage(screen, c.Source, transform)
}
