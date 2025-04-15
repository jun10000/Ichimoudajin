package widget

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetBase struct {
	fontFamily *text.GoTextFaceSource
	fontSize   *float64

	Name            string
	Origin          utility.Vector
	Position        utility.Vector
	Padding         float64
	IsHide          bool
	BorderWidth     float64
	BorderColor     color.Color
	BackgroundColor color.Color
	ForegroundColor color.Color
}

func (w *WidgetBase) Init(inherits WidgetBase) {
	if w.fontFamily == nil {
		w.fontFamily = inherits.fontFamily
	}
	if w.fontSize == nil {
		w.fontSize = inherits.fontSize
	}
}

func (w *WidgetBase) GetFontFamily() *text.GoTextFaceSource {
	return w.fontFamily
}

func (w *WidgetBase) GetFontSize() *float64 {
	return w.fontSize
}

func (w *WidgetBase) SetFontFamily(fontFamily *text.GoTextFaceSource) {
	w.fontFamily = fontFamily
}

func (w *WidgetBase) SetFontSize(fontSize *float64) {
	w.fontSize = fontSize
}

func (w *WidgetBase) GetTextFace() text.Face {
	var s float64
	if w.fontSize != nil {
		s = *w.fontSize
	}

	return &text.GoTextFace{
		Source: w.fontFamily,
		Size:   s,
	}
}

func (w *WidgetBase) GetAlignedArea(outerArea *utility.RectangleF, innerSize utility.Vector) *utility.RectangleF {
	outerPos := outerArea.TopLeft()
	outerSize := outerArea.Size()
	retPos := w.Origin.Mul(outerSize).Sub(w.Origin.Mul(innerSize)).Add(w.Position.Mul(outerSize)).Add(outerPos)
	return utility.NewRectangleF(retPos.X, retPos.Y, retPos.X+innerSize.X, retPos.Y+innerSize.Y)
}
