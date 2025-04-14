package widget

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetText struct {
	*WidgetBase
	Text        string
	BorderWidth float64
	BorderColor color.Color
	FillColor   color.Color
}

func (w *WidgetText) MinSize() utility.Vector {
	if w.fontFamily == nil {
		return utility.ZeroVector()
	}

	x, y := text.Measure(w.Text, w.GetTextFace(), 0)
	x += w.BorderWidth * 2
	y += w.BorderWidth * 2
	return utility.NewVector(x, y)
}

func (w *WidgetText) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	if w.IsHide || w.fontFamily == nil {
		return
	}

	utility.DrawRectangle(screen, preferredArea.TopLeft(), preferredArea.Size(), float32(w.BorderWidth), w.BorderColor, w.FillColor, true)

	r := w.GetAlignedArea(&preferredArea, w.MinSize())
	l := r.TopLeft()
	op := &text.DrawOptions{}
	op.GeoM.Translate(l.X, l.Y)
	text.Draw(screen, w.Text, w.GetTextFace(), op)
}
