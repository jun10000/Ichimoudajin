package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type BlockingArea struct {
	*component.StaticColliderComponent[*utility.RectangleF]
	size utility.Vector
}

func NewBlockingArea(location utility.Vector, size utility.Vector) *BlockingArea {
	t := utility.NewStaticTransform(location, 0, utility.DefaultScale())

	a := &BlockingArea{}
	a.StaticColliderComponent = component.NewStaticColliderComponent(t, a.GetRectangleBounds)
	a.size = size

	a.UpdateBounds()
	return a
}

func (a *BlockingArea) GetRectangleBounds(output *utility.RectangleF) {
	l := a.GetLocation()
	output.MinX = l.X
	output.MinY = l.Y
	output.MaxX = l.X + a.size.X
	output.MaxY = l.Y + a.size.Y
}

func (a *BlockingArea) GetSize() utility.Vector {
	return a.size
}
