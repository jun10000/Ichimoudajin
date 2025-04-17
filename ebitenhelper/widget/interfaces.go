package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetObjecter interface {
	Init(inherits WidgetBase)
	GetFontFamilies() []*text.GoTextFaceSource
	GetFontSize() *float64
	SetFontFamilies(fontFamilies []*text.GoTextFaceSource)
	SetFontSize(fontSize *float64)
	GetName() string
	GetWidgetObject(name string) WidgetObjecter
	MinSize() utility.Vector
	Draw(screen *ebiten.Image, preferredArea utility.RectangleF)
}
