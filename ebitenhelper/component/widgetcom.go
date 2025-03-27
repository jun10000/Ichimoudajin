package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type WidgetCom struct {
	utility.Location
	size utility.Vector
	name string
}

func NewWidgetCom(location utility.Vector, size utility.Vector, name string) *WidgetCom {
	return &WidgetCom{
		Location: utility.NewLocation(location),
		size:     size,
		name:     name,
	}
}

func (c *WidgetCom) ZOrder() int {
	return utility.ZOrderWidget
}

func (c *WidgetCom) GetName() string {
	return c.name
}
