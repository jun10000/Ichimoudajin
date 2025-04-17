package widget

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type WidgetContainerBase struct {
	*WidgetBase
	Children []WidgetObjecter
}

func (w *WidgetContainerBase) Init(inherits WidgetBase) {
	if len(w.fontFamilies) == 0 {
		w.fontFamilies = inherits.fontFamilies
	} else {
		inherits.fontFamilies = w.fontFamilies
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

func (w *WidgetContainerBase) SetFontFamilies(fontFamilies []*text.GoTextFaceSource) {
	oldFontFamilies := w.fontFamilies
	w.fontFamilies = fontFamilies
	for _, o := range w.Children {
		if slices.Equal(o.GetFontFamilies(), oldFontFamilies) {
			o.SetFontFamilies(fontFamilies)
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

func (w *WidgetContainerBase) GetWidgetObject(name string) WidgetObjecter {
	for _, o := range w.Children {
		if o.GetName() == name {
			return o
		}

		gc := o.GetWidgetObject(name)
		if gc != nil {
			return gc
		}
	}

	return nil
}
