package component

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type DrawAnimationComponent struct {
	Image             *ebiten.Image
	FrameCount        int
	FrameSize         utility.Point
	FPS               int
	FrameDirectionMap []int // Front, Left, Right, Back

	parent    utility.Transformer
	tickIndex int
}

func NewDrawAnimationComponent(parent utility.Transformer) *DrawAnimationComponent {
	return &DrawAnimationComponent{
		FrameCount:        3,
		FrameSize:         utility.NewPoint(32, 32),
		FPS:               4,
		FrameDirectionMap: []int{0, 1, 2, 3},
		parent:            parent,
	}
}

func (c *DrawAnimationComponent) Tick() {
	if c.Image == nil {
		return
	}

	c.tickIndex++
}

func (c *DrawAnimationComponent) Draw(screen *ebiten.Image) {
	// Determine sub image index X (Pose)
	idxe := 2*c.FrameCount - 2
	idx := (c.tickIndex * c.FPS / utility.TickCount) % idxe
	if idx >= c.FrameCount {
		idx = idxe - idx
	}

	// Determine sub image index Y (Direction)
	idy := c.FrameDirectionMap[3]
	switch r := c.parent.GetRotation(); {
	case r < math.Pi*-3/4:
		idy = c.FrameDirectionMap[3]
	case r < math.Pi*-1/4:
		idy = c.FrameDirectionMap[2]
	case r < math.Pi*1/4:
		idy = c.FrameDirectionMap[0]
	case r <= math.Pi*3/4:
		idy = c.FrameDirectionMap[1]
	}

	// Draw images
	il := utility.NewPoint(idx*c.FrameSize.X, idy*c.FrameSize.Y)
	img := utility.GetSubImage(c.Image, il, c.FrameSize)
	utility.DrawImage(screen, img, utility.NewTransform(
		c.parent.GetLocation(),
		0,
		c.parent.GetScale(),
	))
}

func (c *DrawAnimationComponent) GetRectangleBounds() utility.RectangleF {
	l := c.parent.GetLocation()
	s := c.FrameSize.ToVector().Mul(c.parent.GetScale())
	return utility.NewRectangleF(l.X, l.Y, l.X+s.X, l.Y+s.Y)
}

func (c *DrawAnimationComponent) GetCircleBounds() utility.CircleF {
	hs := c.FrameSize.ToVector().Mul(c.parent.GetScale()).DivF(2)
	cl := c.parent.GetLocation().Add(hs)
	return utility.NewCircleF(cl.X, cl.Y, math.Max(hs.X, hs.Y))
}
