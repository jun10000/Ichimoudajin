package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetObjecter interface {
	Init(inherits WidgetBase)
	GetFontFamily() *text.GoTextFaceSource
	GetFontSize() *float64
	SetFontFamily(fontFamily *text.GoTextFaceSource)
	SetFontSize(fontSize *float64)
	MinSize() utility.Vector
	Draw(screen *ebiten.Image, preferredArea utility.RectangleF)
}
