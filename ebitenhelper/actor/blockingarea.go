package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type BlockingArea struct {
	utility.Transform
	*component.ColliderComponent[*utility.RectangleF]

	Size utility.Vector
}

func NewBlockingArea() *BlockingArea {
	a := &BlockingArea{
		Transform: utility.DefaultTransform(),
		Size:      utility.NewVector(32, 32),
	}

	a.ColliderComponent = component.NewColliderComponent(a.getBounds)
	return a
}

func (a *BlockingArea) getBounds(output *utility.RectangleF) {
	l := a.GetLocation()

	output.MinX = l.X
	output.MinY = l.Y
	output.MaxX = l.X + a.Size.X
	output.MaxY = l.Y + a.Size.Y
}
