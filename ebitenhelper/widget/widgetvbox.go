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

	return ret.AddF(w.BorderWidth*2 + w.Margin*2 + w.Padding*2)
}

func (w *WidgetVBox) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	if w.IsHide {
		return
	}

	r := w.GetAlignedArea(&preferredArea, w.MinSize())
	r.MinX += w.Margin
	r.MinY += w.Margin
	r.MaxX -= w.Margin
	r.MaxY -= w.Margin
	utility.DrawRectangle(screen, r.TopLeft(), r.Size(), float32(w.BorderWidth), w.BorderColor, w.BackgroundColor, true)

	r.MinX += w.Padding
	r.MinY += w.Padding
	r.MaxX -= w.Padding
	for _, o := range w.Children {
		s := o.MinSize()
		r.MaxY = r.MinY + s.Y
		o.Draw(screen, *r)
		r.MinY = r.MaxY
	}
}
