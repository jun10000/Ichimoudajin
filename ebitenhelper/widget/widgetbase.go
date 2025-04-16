package widget

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetBase struct {
	fontFamilies []*text.GoTextFaceSource
	fontSize     *float64

	Name            string
	Origin          utility.Vector
	Offset          utility.Vector
	Margin          utility.Inset
	Padding         utility.Inset
	IsHide          bool
	BorderWidth     float64
	BorderColor     color.Color
	BackgroundColor color.Color
	ForegroundColor color.Color
}

func (w *WidgetBase) Init(inherits WidgetBase) {
	if len(w.fontFamilies) == 0 {
		w.fontFamilies = inherits.fontFamilies
	}
	if w.fontSize == nil {
		w.fontSize = inherits.fontSize
	}
}

func (w *WidgetBase) GetFontFamilies() []*text.GoTextFaceSource {
	return w.fontFamilies
}

func (w *WidgetBase) GetFontSize() *float64 {
	return w.fontSize
}

func (w *WidgetBase) SetFontFamilies(fontFamilies []*text.GoTextFaceSource) {
	w.fontFamilies = fontFamilies
}

func (w *WidgetBase) SetFontSize(fontSize *float64) {
	w.fontSize = fontSize
}

func (w *WidgetBase) MinSize() utility.Vector {
	x := w.Padding.Left + w.Padding.Right + w.BorderWidth*2 + w.Margin.Left + w.Margin.Right
	y := w.Padding.Top + w.Padding.Bottom + w.BorderWidth*2 + w.Margin.Top + w.Margin.Bottom
	return utility.NewVector(x, y)
}

func (w *WidgetBase) GetTextFace() text.Face {
	var s float64
	if w.fontSize != nil {
		s = *w.fontSize
	}

	fs := make([]text.Face, 0, utility.InitialWidgetFontCap)
	for _, f := range w.fontFamilies {
		fs = append(fs, &text.GoTextFace{
			Source: f,
			Size:   s,
		})
	}

	mf, err := text.NewMultiFace(fs...)
	utility.PanicIfError(err)
	return mf
}

func (w *WidgetBase) GetAlignedArea(outerArea *utility.RectangleF, innerSize utility.Vector) *utility.RectangleF {
	outerPos := outerArea.TopLeft()
	outerSize := outerArea.Size()
	retPos := w.Origin.Mul(outerSize).Sub(w.Origin.Mul(innerSize)).Add(w.Offset.Mul(outerSize)).Add(outerPos)
	return utility.NewRectangleF(retPos.X, retPos.Y, retPos.X+innerSize.X, retPos.Y+innerSize.Y)
}

func (w *WidgetBase) DrawBackground(screen *ebiten.Image, preferredArea utility.RectangleF) {
	preferredArea.MinX += w.Margin.Left + w.BorderWidth/2
	preferredArea.MinY += w.Margin.Top + w.BorderWidth/2
	preferredArea.MaxX -= w.Margin.Right + w.BorderWidth/2
	preferredArea.MaxY -= w.Margin.Bottom + w.BorderWidth/2
	utility.DrawRectangle(screen, preferredArea.TopLeft(), preferredArea.Size(), float32(w.BorderWidth), w.BorderColor, w.BackgroundColor, true)
}

func (w *WidgetBase) BackgroundToForegroundArea(out *utility.RectangleF) {
	out.MinX += w.Margin.Left + w.BorderWidth + w.Padding.Left
	out.MinY += w.Margin.Top + w.BorderWidth + w.Padding.Top
	out.MaxX -= w.Margin.Right + w.BorderWidth + w.Padding.Right
	out.MaxY -= w.Margin.Bottom + w.BorderWidth + w.Padding.Bottom
}
