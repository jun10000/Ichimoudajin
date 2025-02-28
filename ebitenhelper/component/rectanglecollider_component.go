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

func (c *RectangleColliderComponent) GetColliderBounds() []utility.Bounder {
	b := c.boundsCache[0].(*utility.RectangleF)
	c.getBounds(b)
	if !utility.GetLevel().IsLooping {
		return c.boundsCache[:1]
	}

	s := utility.GetGameInstance().ScreenSize.ToVector()
	b.Offset(-s.X, -s.Y, c.boundsCache[1])
	b.Offset(0, -s.Y, c.boundsCache[2])
	b.Offset(s.X, -s.Y, c.boundsCache[3])
	b.Offset(-s.X, 0, c.boundsCache[4])
	b.Offset(s.X, 0, c.boundsCache[5])
	b.Offset(-s.X, s.Y, c.boundsCache[6])
	b.Offset(0, s.Y, c.boundsCache[7])
	b.Offset(s.X, s.Y, c.boundsCache[8])
	return c.boundsCache
}
