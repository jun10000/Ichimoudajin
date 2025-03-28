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

func (g ActorGeneratorStruct) NewBlockingArea(options *NewActorOptions) *BlockingArea {
	t := utility.NewStaticTransform(options.Location, options.Rotation, options.Scale)

	a := &BlockingArea{}
	a.ActorCom = component.NewActorCom(options.Name)
	a.StaticColliderCom = component.NewStaticColliderCom(t, a.GetRectangleBounds)
	a.size = options.Size

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
