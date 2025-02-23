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
	l := a.GetLocation()
	s := a.Size
	return utility.NewRectangleF(l.X, l.Y, l.X+s.X, l.Y+s.Y)
}
