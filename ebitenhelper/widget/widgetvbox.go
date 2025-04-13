package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetVBox struct {
	*WidgetContainerFields
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
}
