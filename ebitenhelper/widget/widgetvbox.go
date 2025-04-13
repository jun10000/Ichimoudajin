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
	r := w.GetAlignedArea(&preferredArea, w.MinSize())
	for _, o := range w.Children {
		s := o.MinSize()
		r.MaxY = r.MinY + s.Y
		o.Draw(screen, *r)
		r.MinY = r.MaxY
	}
}
