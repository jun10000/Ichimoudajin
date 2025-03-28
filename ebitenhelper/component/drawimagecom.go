package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type DrawImageCom struct {
	*DrawCom
	parent utility.StaticTransformer

	Image *ebiten.Image
}

func NewDrawImageCom(parent utility.StaticTransformer, isVisible bool) *DrawImageCom {
	return &DrawImageCom{
		DrawCom: NewDrawCom(isVisible),
		parent:  parent,
	}
}

func (c *DrawImageCom) Draw(screen *ebiten.Image) {
	utility.DrawImage(screen, c.Image, c.parent)
}
