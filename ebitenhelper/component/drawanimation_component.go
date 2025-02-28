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

func (c *DrawAnimationComponent) GetRectangleBounds(output *utility.RectangleF) {
	loc := c.parent.GetLocation()
	scale := c.parent.GetScale()

	output.MinX = loc.X
	output.MinY = loc.Y
	output.MaxX = loc.X + float64(c.FrameSize.X)*scale.X
	output.MaxY = loc.Y + float64(c.FrameSize.Y)*scale.Y
}

func (c *DrawAnimationComponent) GetCircleBounds(output *utility.CircleF) {
	loc := c.parent.GetLocation()
	scale := c.parent.GetScale()
	hsx := float64(c.FrameSize.X) * scale.X / 2
	hsy := float64(c.FrameSize.Y) * scale.Y / 2

	output.OrgX = loc.X + hsx
	output.OrgY = loc.Y + hsy
	output.Radius = math.Max(hsx, hsy)
}
