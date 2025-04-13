package widget

import "github.com/hajimehoshi/ebiten/v2/text/v2"

type WidgetContainerBase struct {
	*WidgetBase
	Children []WidgetObjecter
}

func (f *WidgetContainerBase) Init(inherits WidgetBase) {
	if f.font == nil {
		f.font = inherits.font
	} else {
		inherits.font = f.font
	}

	for _, o := range f.Children {
		o.Init(inherits)
	}
}

func (f *WidgetContainerBase) SetFont(font *text.GoTextFace) {
	oldFont := f.font
	f.font = font
	for _, o := range f.Children {
		if o.GetFont() == oldFont {
			o.SetFont(font)
		}
	}
}
