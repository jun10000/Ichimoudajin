package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type RectangleColliderComponent struct {
	getBounds   func(*utility.RectangleF)
	boundsCache []utility.Bounder
}

func NewRectangleColliderComponent(getBounds func(*utility.RectangleF)) *RectangleColliderComponent {
	c := &RectangleColliderComponent{
		getBounds: getBounds,
		boundsCache: []utility.Bounder{
			&utility.RectangleF{},
			&utility.RectangleF{},
			&utility.RectangleF{},
			&utility.RectangleF{},
			&utility.RectangleF{},
			&utility.RectangleF{},
			&utility.RectangleF{},
			&utility.RectangleF{},
			&utility.RectangleF{},
		},
	}

	return c
}

func (c *RectangleColliderComponent) GetMainColliderBounds() utility.Bounder {
	b := c.boundsCache[0].(*utility.RectangleF)
	c.getBounds(b)
	return b
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

		s := utility.GetGameInstance().ScreenSize.ToVector()
		b.Offset(-s.X, -s.Y, c.boundsCache[1])
		if !yield(c.boundsCache[1]) {
			return
		}
		b.Offset(0, -s.Y, c.boundsCache[2])
		if !yield(c.boundsCache[2]) {
			return
		}
		b.Offset(s.X, -s.Y, c.boundsCache[3])
		if !yield(c.boundsCache[3]) {
			return
		}
		b.Offset(-s.X, 0, c.boundsCache[4])
		if !yield(c.boundsCache[4]) {
			return
		}
		b.Offset(s.X, 0, c.boundsCache[5])
		if !yield(c.boundsCache[5]) {
			return
		}
		b.Offset(-s.X, s.Y, c.boundsCache[6])
		if !yield(c.boundsCache[6]) {
			return
		}
		b.Offset(0, s.Y, c.boundsCache[7])
		if !yield(c.boundsCache[7]) {
			return
		}
		b.Offset(s.X, s.Y, c.boundsCache[8])
		if !yield(c.boundsCache[8]) {
			return
		}
	}
}
