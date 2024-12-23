package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type DrawImageComponent struct {
	Image  *ebiten.Image
	parent utility.Transformer
}

func NewDrawImageComponent(parent utility.Transformer) *DrawImageComponent {
	return &DrawImageComponent{
		parent: parent,
	}
}

func (c *DrawImageComponent) Draw(screen *ebiten.Image) {
	utility.DrawImage(screen, c.Image, c.parent)
}
