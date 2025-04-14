package widget

import "github.com/hajimehoshi/ebiten/v2/text/v2"

type WidgetContainerBase struct {
	*WidgetBase
	Children []WidgetObjecter
}

func (w *WidgetContainerBase) Init(inherits WidgetBase) {
	if w.fontFamily == nil {
		w.fontFamily = inherits.fontFamily
	} else {
		inherits.fontFamily = w.fontFamily
	}

	if w.fontSize == nil {
		w.fontSize = inherits.fontSize
	} else {
		inherits.fontSize = w.fontSize
	}

	for _, o := range w.Children {
		o.Init(inherits)
	}
}

func (w *WidgetContainerBase) SetFontFamily(fontFamily *text.GoTextFaceSource) {
	oldFontFamily := w.fontFamily
	w.fontFamily = fontFamily
	for _, o := range w.Children {
		if o.GetFontFamily() == oldFontFamily {
			o.SetFontFamily(fontFamily)
		}
	}
}

func (w *WidgetContainerBase) SetFontSize(fontSize *float64) {
	oldFontSize := w.fontSize
	w.fontSize = fontSize
	for _, o := range w.Children {
		if o.GetFontSize() == oldFontSize {
			o.SetFontSize(fontSize)
		}
	}
}
