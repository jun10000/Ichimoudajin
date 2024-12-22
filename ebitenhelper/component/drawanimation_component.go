package component

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type DrawAnimationComponent struct {
	Source            *ebiten.Image
	FrameCount        int
	FrameSize         utility.Point
	FPS               int
	FrameDirectionMap []int // Front, Left, Right, Back

	tickIndex int
}

func NewDrawAnimationComponent() *DrawAnimationComponent {
	return &DrawAnimationComponent{
		FrameCount:        3,
		FrameSize:         utility.NewPoint(32, 32),
		FPS:               4,
		FrameDirectionMap: []int{0, 1, 2, 3},
	}
}

func (c *DrawAnimationComponent) Tick() {
	if c.Source == nil {
		return
	}

	c.tickIndex++
}

func (c *DrawAnimationComponent) Draw(screen *ebiten.Image, transform utility.Transformer) {
	// Determine sub image index X (Pose)
	idxe := 2*c.FrameCount - 2
	idx := (c.tickIndex * c.FPS / utility.TickCount) % idxe
	if idx >= c.FrameCount {
		idx = idxe - idx
	}

	// Determine sub image index Y (Direction)
	idy := c.FrameDirectionMap[3]
	switch r := transform.GetRotation(); {
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
	img := utility.GetSubImage(c.Source, il, c.FrameSize)
	utility.DrawImage(screen, img, utility.NewTransform(
		transform.GetLocation(),
		0,
		transform.GetScale(),
	))
}

func (c *DrawAnimationComponent) GetCircleBounds(transform utility.Transformer) utility.CircleF {
	hs := c.FrameSize.ToVector().DivF(2).Mul(transform.GetScale())
	return utility.NewCircleF(transform.GetLocation().Add(hs), math.Max(hs.X, hs.Y))
}
