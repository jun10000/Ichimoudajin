package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetButton struct {
	*WidgetCommonFields
	Text string
}

func (w *WidgetButton) MinSize() utility.Vector {
	return utility.ZeroVector()
}

func (w *WidgetButton) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {

}
