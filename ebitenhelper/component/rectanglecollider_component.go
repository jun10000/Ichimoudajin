package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type RectangleColliderComponent struct {
	getBounds   func(*utility.RectangleF)
	boundsCache *utility.RectangleF
}

func NewRectangleColliderComponent(getBounds func(*utility.RectangleF)) *RectangleColliderComponent {
	return &RectangleColliderComponent{
		getBounds:   getBounds,
		boundsCache: &utility.RectangleF{},
	}
}

func (c *RectangleColliderComponent) GetColliderBounds() utility.Bounder {
	c.getBounds(c.boundsCache)
	return c.boundsCache
}
