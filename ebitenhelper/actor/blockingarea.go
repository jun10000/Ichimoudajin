package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type BlockingArea struct {
	*component.StaticColliderComponent[*utility.RectangleF]
	size utility.Vector
}

func NewBlockingArea(sTransform *utility.StaticTransform, size utility.Vector) *BlockingArea {
	a := &BlockingArea{
		size: size,
	}

	a.StaticColliderComponent = component.NewStaticColliderComponent(sTransform, a.GetBounds)
	a.UpdateColliderBounds()
	return a
}

func (a *BlockingArea) GetBounds(output *utility.RectangleF) {
	l := a.GetLocation()
	output.MinX = l.X
	output.MinY = l.Y
	output.MaxX = l.X + a.size.X
	output.MaxY = l.Y + a.size.Y
}

func (a *BlockingArea) GetSize() utility.Vector {
	return a.size
}
