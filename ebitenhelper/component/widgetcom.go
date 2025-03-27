package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type WidgetCom struct {
	utility.Location
	size utility.Vector
}

func NewWidgetCom(location utility.Vector, size utility.Vector) *WidgetCom {
	return &WidgetCom{
		Location: utility.NewLocation(location),
		size:     size,
	}
}

func (c *WidgetCom) ZOrder() int {
	return utility.ZOrderWidget
}
