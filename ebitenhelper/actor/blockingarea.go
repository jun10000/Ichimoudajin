package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type BlockingArea struct {
	*component.ActorCom
	*component.StaticColliderCom[*utility.RectangleF]
	size utility.Vector
}

func (g ActorGeneratorStruct) NewBlockingArea(name string, location utility.Vector, rotation float64, scale utility.Vector, size utility.Vector) *BlockingArea {
	t := utility.NewStaticTransform(location, rotation, scale)

	a := &BlockingArea{}
	a.ActorCom = component.NewActorCom(name)
	a.StaticColliderCom = component.NewStaticColliderCom(t, a.GetRectangleBounds)
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
