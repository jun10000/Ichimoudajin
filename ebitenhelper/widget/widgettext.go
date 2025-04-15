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

	f := w.GetTextFace()
	m := f.Metrics()
	x, y := text.Measure(w.Text, f, 0)
	x += w.BorderWidth*2 + w.Margin*2 + w.Padding*2
	y += w.BorderWidth*2 + w.Margin*2 + w.Padding*2 - m.HDescent
	return utility.NewVector(x, y)
}

func (w *WidgetText) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	if w.IsHide || w.fontFamily == nil {
		return
	}

	r := w.GetAlignedArea(&preferredArea, w.MinSize())
	r.MinX += w.Margin + w.BorderWidth/2
	r.MinY += w.Margin + w.BorderWidth/2
	r.MaxX -= w.Margin + w.BorderWidth/2
	r.MaxY -= w.Margin + w.BorderWidth/2
	utility.DrawRectangle(screen, r.TopLeft(), r.Size(), float32(w.BorderWidth), w.BorderColor, w.BackgroundColor, true)

	l := r.TopLeft().AddF(w.BorderWidth/2 + w.Padding)
	op := &text.DrawOptions{}
	op.GeoM.Translate(l.X, l.Y)
	op.ColorScale.ScaleWithColor(w.ForegroundColor)
	text.Draw(screen, w.Text, w.GetTextFace(), op)
}
