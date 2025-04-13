package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetVBox struct {
	*WidgetContainerBase
}

func (w *WidgetVBox) MinSize() utility.Vector {
	ret := utility.ZeroVector()

	for _, o := range w.Children {
		s := o.MinSize()
		ret.Y += s.Y
		if s.X > ret.X {
			ret.X = s.X
		}
	}

	return ret
}

func (w *WidgetVBox) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	parentSize := preferredArea.Size()
	preferredArea.MinX += parentSize.X * w.Position.X
	preferredArea.MinY += parentSize.Y * w.Position.Y

	for _, o := range w.Children {
		s := o.MinSize()
		preferredArea.MaxX = preferredArea.MinX + s.X
		preferredArea.MaxY = preferredArea.MinY + s.Y
		o.Draw(screen, preferredArea)
		preferredArea.MinY = preferredArea.MaxY
	}
}
