package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type BlockingArea struct {
	utility.Transform
	Size utility.Vector
}

func NewBlockingArea() *BlockingArea {
	return &BlockingArea{
		Transform: utility.DefaultTransform(),
		Size:      utility.NewVector(32, 32),
	}
}

func (a *BlockingArea) GetColliderBounds() utility.Bounder {
	return utility.NewRectangleF(a.GetLocation(), a.Size)
}
