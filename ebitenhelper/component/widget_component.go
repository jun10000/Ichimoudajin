package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type WidgetComponent struct {
	utility.Location
	size utility.Vector
	name string
}

func NewWidgetComponent(location utility.Vector, size utility.Vector, name string) *WidgetComponent {
	return &WidgetComponent{
		Location: utility.NewLocation(location),
		size:     size,
		name:     name,
	}
}

func (c *WidgetComponent) ZOrder() int {
	return utility.ZOrderWidget
}

func (c *WidgetComponent) GetName() string {
	return c.name
}
