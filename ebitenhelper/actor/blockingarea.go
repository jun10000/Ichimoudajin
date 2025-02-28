package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type BlockingArea struct {
	utility.Transform
	*component.RectangleColliderComponent

	Size utility.Vector
}

func NewBlockingArea() *BlockingArea {
	a := &BlockingArea{
		Transform: utility.DefaultTransform(),
		Size:      utility.NewVector(32, 32),
	}

	a.RectangleColliderComponent = component.NewRectangleColliderComponent(a.getBounds)
	return a
}

func (a *BlockingArea) getBounds(input *utility.RectangleF) {
	l := a.GetLocation()

	input.MinX = l.X
	input.MinY = l.Y
	input.MaxX = l.X + a.Size.X
	input.MaxY = l.Y + a.Size.Y
}
