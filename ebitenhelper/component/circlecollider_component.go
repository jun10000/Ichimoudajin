package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type CircleColliderComponent struct {
	getBounds   func(*utility.CircleF)
	boundsCache *utility.CircleF
}

func NewCircleColliderComponent(getBounds func(*utility.CircleF)) *CircleColliderComponent {
	return &CircleColliderComponent{
		getBounds:   getBounds,
		boundsCache: &utility.CircleF{},
	}
}

func (c *CircleColliderComponent) GetColliderBounds() utility.Bounder {
	c.getBounds(c.boundsCache)
	return c.boundsCache
}
