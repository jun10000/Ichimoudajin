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

	return ret.Add(w.WidgetContainerBase.MinSize())
}

func (w *WidgetHBox) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	if w.IsHide {
		return
	}

	s := utility.NewRectangleFFromGoRect(screen.Bounds())
	r := w.GetAlignedArea(s, &preferredArea, w.MinSize())
	w.DrawBackground(screen, *r)

	w.BackgroundToForegroundArea(r)
	for _, o := range w.Children {
		s := o.MinSize()
		r.MaxX = r.MinX + s.X
		o.Draw(screen, *r)
		r.MinX = r.MaxX
	}
}
