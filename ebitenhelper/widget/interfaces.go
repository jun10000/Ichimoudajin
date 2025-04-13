package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetObjecter interface {
	Init(inherits WidgetBase)
	GetFont() *text.GoTextFace
	SetFont(font *text.GoTextFace)
	MinSize() utility.Vector
	Draw(screen *ebiten.Image, preferredArea utility.RectangleF)
}
