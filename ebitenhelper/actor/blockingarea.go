package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type BlockingArea struct {
	*component.ActorCom
	*component.StaticColliderCom[*utility.RectangleF]
	utility.StaticSize
}

func (g ActorGeneratorStruct) NewBlockingArea(options *NewActorOptions) *BlockingArea {
	t := utility.NewStaticTransform(options.Location, options.Rotation, options.Scale)

	a := &BlockingArea{}
	a.ActorCom = component.NewActorCom(options.Name)
	a.StaticColliderCom = component.NewStaticColliderCom(t, a.GetRectangleBounds)
	a.StaticSize = utility.NewStaticSize(options.Size)

	a.UpdateBounds()
	return a
}

func (a *BlockingArea) GetRectangleBounds(output *utility.RectangleF) {
	l := a.GetLocation()
	sz := a.GetSize()
	output.MinX = l.X
	output.MinY = l.Y
	output.MaxX = l.X + sz.X
	output.MaxY = l.Y + sz.Y
}
