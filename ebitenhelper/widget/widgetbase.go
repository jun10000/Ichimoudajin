package widget

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetCommonFields struct {
	font *text.GoTextFace

	Name     string
	Origin   utility.Vector
	Position utility.Vector
	IsHide   bool
}

func (f *WidgetCommonFields) Init(inherits WidgetCommonFields) {
	if f.font == nil {
		f.font = inherits.font
	}
}

func (f *WidgetCommonFields) GetFont() *text.GoTextFace {
	return f.font
}

func (f *WidgetCommonFields) SetFont(font *text.GoTextFace) {
	f.font = font
}

type WidgetContainerFields struct {
	*WidgetCommonFields
	Children []WidgetObjecter
}

func (f *WidgetContainerFields) Init(inherits WidgetCommonFields) {
	if f.font == nil {
		f.font = inherits.font
	} else {
		inherits.font = f.font
	}

	for _, o := range f.Children {
		o.Init(inherits)
	}
}

func (f *WidgetContainerFields) SetFont(font *text.GoTextFace) {
	oldFont := f.font
	f.font = font
	for _, o := range f.Children {
		if o.GetFont() == oldFont {
			o.SetFont(font)
		}
	}
}
