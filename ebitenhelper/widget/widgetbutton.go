package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetButton struct {
	*WidgetText
	Border Border
}

func (w *WidgetButton) MinSize() utility.Vector {
	return w.WidgetText.MinSize().AddXY(w.Border.Left+w.Border.Right, w.Border.Top+w.Border.Bottom)
}

func (w *WidgetButton) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {

}
