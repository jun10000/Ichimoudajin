package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type RectangleColliderComponent struct {
	getBounds   func(*utility.RectangleF)
	loopOffsets []utility.Vector
	mainCache   *utility.RectangleF
	offsetCache *utility.RectangleF
}

func NewRectangleColliderComponent(getBounds func(*utility.RectangleF)) *RectangleColliderComponent {
	s := utility.ScreenSize.ToVector()
	return &RectangleColliderComponent{
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
		mainCache:   &utility.RectangleF{},
		offsetCache: &utility.RectangleF{},
	}
}

func (c *RectangleColliderComponent) GetMainColliderBounds() utility.Bounder {
	c.getBounds(c.mainCache)
	return c.mainCache
}

func (c *RectangleColliderComponent) GetColliderBounds() func(yield func(utility.Bounder) bool) {
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
