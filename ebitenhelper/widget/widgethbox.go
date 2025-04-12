package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetHBox struct {
	*WidgetContainerFields
}

func (w *WidgetHBox) MinSize() utility.Vector {
	ret := utility.ZeroVector()

	for _, o := range w.Children {
		s := o.MinSize()
		ret.X += s.X
		if s.Y > ret.Y {
			ret.Y = s.Y
		}
	}

	return ret
}

func (w *WidgetHBox) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	preferredArea.MinX += w.Position.X
	preferredArea.MinY += w.Position.Y

	for _, o := range w.Children {
		s := o.MinSize()
		preferredArea.MaxX = preferredArea.MinX + s.X
		preferredArea.MaxY = preferredArea.MinY + s.Y
		o.Draw(screen, preferredArea)
		preferredArea.MinX = preferredArea.MaxX
	}
}
