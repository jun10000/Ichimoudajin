package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetText struct {
	*WidgetBase
	Text string
}

func (w *WidgetText) MinSize() utility.Vector {
	if w.fontFamily == nil {
		return utility.ZeroVector()
	}

	s := w.WidgetBase.MinSize()
	f := w.GetTextFace()
	x, y := text.Measure(w.Text, f, 0)
	ret := utility.NewVector(x+s.X, y+s.Y-f.Metrics().HDescent)
	return ret
}

func (w *WidgetText) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	if w.IsHide || w.fontFamily == nil {
		return
	}

	r := w.GetAlignedArea(&preferredArea, w.MinSize())
	w.DrawBackground(screen, *r)

	w.BackgroundToForegroundArea(r)
	l := r.TopLeft()
	op := &text.DrawOptions{}
	op.GeoM.Translate(l.X, l.Y)
	op.ColorScale.ScaleWithColor(w.ForegroundColor)
	text.Draw(screen, w.Text, w.GetTextFace(), op)
}
