package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetObject interface {
}

type WidgetCommonFields struct {
	Name     string
	Origin   utility.Vector
	Position utility.Vector
	IsHide   bool
}

type WidgetHBox struct {
	*WidgetCommonFields
	Children []WidgetObject
}

type WidgetVBox struct {
	*WidgetCommonFields
	Children []WidgetObject
}

type WidgetText struct {
	*WidgetCommonFields
	Text string `xml:",chardata"`
}

type WidgetButton struct {
	*WidgetCommonFields
	Text string `xml:",chardata"`
}

type Widget struct {
	*component.ActorCom
	*component.DrawCom
	Children []WidgetObject
}

func (a *Widget) Draw(screen *ebiten.Image) {

}

func (a *Widget) ZOrder() int {
	return utility.ZOrderWidget
}
