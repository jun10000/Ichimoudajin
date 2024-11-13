package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type BlockingArea struct {
	Location utility.Vector
	Size     utility.Vector
}

func NewBlockingArea() *BlockingArea {
	return &BlockingArea{
		Location: utility.ZeroVector(),
		Size:     utility.NewVector(32, 32),
	}
}

func (a *BlockingArea) GetBounds() utility.RectangleF {
	return utility.NewRectangleF(a.Location, a.Size)
}
