package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type CircleColliderComponent struct {
	getBounds   func(*utility.CircleF)
	loopOffsets []utility.Vector
	mainCache   *utility.CircleF
	offsetCache *utility.CircleF
}

func NewCircleColliderComponent(getBounds func(*utility.CircleF)) *CircleColliderComponent {
	s := utility.ScreenSize.ToVector()
	return &CircleColliderComponent{
		getBounds: getBounds,
		loopOffsets: []utility.Vector{
			utility.NewVector(-s.X, -s.Y),
			utility.NewVector(0, -s.Y),
			utility.NewVector(s.X, -s.Y),
			utility.NewVector(-s.X, 0),
			utility.NewVector(s.X, 0),
			utility.NewVector(-s.X, s.Y),
			utility.NewVector(0, s.Y),
			utility.NewVector(s.X, s.Y),
		},
		mainCache:   &utility.CircleF{},
		offsetCache: &utility.CircleF{},
	}
}

func (c *CircleColliderComponent) GetMainColliderBounds() utility.Bounder {
	c.getBounds(c.mainCache)
	return c.mainCache
}

func (c *CircleColliderComponent) GetColliderBounds() func(yield func(utility.Bounder) bool) {
	return func(yield func(utility.Bounder) bool) {
		b := c.GetMainColliderBounds()
		if !yield(b) {
			return
		}
		if !utility.GetLevel().IsLooping {
			return
		}

		for _, v := range c.loopOffsets {
			b.Offset(v.X, v.Y, c.offsetCache)
			if !yield(*c.offsetCache) {
				return
			}
		}
	}
}
