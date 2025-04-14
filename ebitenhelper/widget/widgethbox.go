package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetHBox struct {
	*WidgetContainerBase
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

	return ret.AddF(w.BorderWidth * 2)
}

func (w *WidgetHBox) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	if w.IsHide {
		return
	}

	r := w.GetAlignedArea(&preferredArea, w.MinSize())
	utility.DrawRectangle(screen, r.TopLeft(), r.Size(), float32(w.BorderWidth), w.BorderColor, w.BackgroundColor, true)

	for _, o := range w.Children {
		s := o.MinSize()
		r.MaxX = r.MinX + s.X
		o.Draw(screen, *r)
		r.MinX = r.MaxX
	}
}
