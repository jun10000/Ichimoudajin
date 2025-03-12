package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type BlockingArea struct {
	*component.ColliderComponent[*utility.RectangleF]

	size utility.Vector
}

func NewBlockingArea() *BlockingArea {
	a := &BlockingArea{
		size: utility.NewVector(32, 32),
	}

	a.ColliderComponent = component.NewColliderComponent(a.getBounds)
	return a
}

func (a *BlockingArea) GetSize() utility.Vector {
	return a.size
}

func (a *BlockingArea) SetSize(size utility.Vector) {
	a.size = size
	a.UpdateColliderBounds()
}

func (a *BlockingArea) getBounds(output *utility.RectangleF) {
	l := a.GetLocation()
	output.MinX = l.X
	output.MinY = l.Y
	output.MaxX = l.X + a.size.X
	output.MaxY = l.Y + a.size.Y
}
