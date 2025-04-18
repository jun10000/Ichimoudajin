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

func (w *WidgetText) MinSize(screenSize *utility.Vector) utility.Vector {
	if len(w.fontFamilies) == 0 {
		return utility.ZeroVector()
	}

	s := w.WidgetBase.MinSize(screenSize)
	x, y := text.Measure(w.Text, w.GetTextFace(), 0)
	ret := utility.NewVector(x+s.X, y+s.Y)
	return ret
}

func (w *WidgetText) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	if w.IsHide || len(w.fontFamilies) == 0 {
		return
	}

	s := utility.NewRectangleFFromGoRect(screen.Bounds())
	ssz := s.Size()
	r := w.GetAlignedArea(s, &preferredArea, w.MinSize(&ssz))
	w.DrawBackground(screen, *r)

	w.BackgroundToForegroundArea(&ssz, r)
	l := r.TopLeft()
	op := &text.DrawOptions{}
	op.GeoM.Translate(l.X, l.Y)
	op.ColorScale.ScaleWithColor(w.ForegroundColor)
	text.Draw(screen, w.Text, w.GetTextFace(), op)
}
