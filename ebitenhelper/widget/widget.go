package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Widget struct {
	*component.ActorCom
	*component.DrawCom
	*WidgetContainerBase
}

func (a *Widget) Draw(screen *ebiten.Image) {
	r := utility.NewRectangleFFromGoRect(screen.Bounds())
	for _, o := range a.Children {
		o.Draw(screen, *r)
	}
}

func (a *Widget) ZOrder() int {
	return utility.ZOrderWidget
}
