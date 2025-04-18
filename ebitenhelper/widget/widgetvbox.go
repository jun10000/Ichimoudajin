package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetVBox struct {
	*WidgetContainerBase
}

func (w *WidgetVBox) MinSize(screenSize *utility.Vector) utility.Vector {
	ret := utility.ZeroVector()
	for _, o := range w.Children {
		s := o.MinSize(screenSize)
		ret.Y += s.Y
		if s.X > ret.X {
			ret.X = s.X
		}
	}

	return ret.Add(w.WidgetContainerBase.MinSize(screenSize))
}

func (w *WidgetVBox) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	if w.IsHide {
		return
	}

	s := utility.NewRectangleFFromGoRect(screen.Bounds())
	ssz := s.Size()
	r := w.GetAlignedArea(s, &preferredArea, w.MinSize(&ssz))
	w.DrawBackground(screen, *r)

	w.BackgroundToForegroundArea(&ssz, r)
	for _, o := range w.Children {
		s := o.MinSize(&ssz)
		r.MaxY = r.MinY + s.Y
		o.Draw(screen, *r)
		r.MinY = r.MaxY
	}
}
