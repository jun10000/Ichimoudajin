package widget

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetBase struct {
	font *text.GoTextFace

	Name     string
	Origin   utility.Vector
	Position utility.Vector
	IsHide   bool
}

func (f *WidgetBase) Init(inherits WidgetBase) {
	if f.font == nil {
		f.font = inherits.font
	}
}

func (f *WidgetBase) GetFont() *text.GoTextFace {
	return f.font
}

func (f *WidgetBase) SetFont(font *text.GoTextFace) {
	f.font = font
}

func (f *WidgetBase) GetAlignedArea(outerArea *utility.RectangleF, innerSize utility.Vector) *utility.RectangleF {
	outerPos := outerArea.TopLeft()
	outerSize := outerArea.Size()
	retPos := f.Origin.Mul(outerSize).Sub(f.Origin.Mul(innerSize)).Add(f.Position.Mul(outerSize)).Add(outerPos)
	return utility.NewRectangleF(retPos.X, retPos.Y, retPos.X+innerSize.X, retPos.Y+innerSize.Y)
}
