package widget

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetText struct {
	*WidgetCommonFields
	Text        string
	BorderWidth float64
	BorderColor color.Color
	FillColor   color.Color
}

func (w *WidgetText) MinSize() utility.Vector {
	if w.font == nil {
		return utility.ZeroVector()
	}

	x, y := text.Measure(w.Text, w.font, 0)
	x += w.BorderWidth * 2
	y += w.BorderWidth * 2
	return utility.NewVector(x, y)
}

func (w *WidgetText) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	if w.IsHide || w.font == nil {
		return
	}

	innerSize := w.MinSize()
	outerSize := preferredArea.Size()
	origin := w.Origin.DivF(100)
	offset := origin.Mul(outerSize).Sub(origin.Mul(innerSize)).Add(w.Position)

	op := &text.DrawOptions{}
	op.GeoM.Translate(preferredArea.MinX+offset.X, preferredArea.MinY+offset.Y)

	utility.DrawRectangle(screen, preferredArea.TopLeft(), preferredArea.Size(), float32(w.BorderWidth), w.BorderColor, w.FillColor, true)
	text.Draw(screen, w.Text, w.font, op)
}
