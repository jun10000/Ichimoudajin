package component

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type DrawAnimationCom struct {
	*DrawCom
	parent utility.StaticTransformer

	Image             *ebiten.Image
	FrameCount        int
	FrameSize         utility.Point
	FPS               int
	FrameDirectionMap []int // Front, Left, Right, Back
}

func NewDrawAnimationCom(parent utility.StaticTransformer, isVisible bool) *DrawAnimationCom {
	return &DrawAnimationCom{
		DrawCom: NewDrawCom(isVisible),
		parent:  parent,

		FrameCount:        3,
		FrameSize:         utility.NewPoint(32, 32),
		FPS:               4,
		FrameDirectionMap: []int{0, 1, 2, 3},
	}
}

func (c *DrawAnimationCom) Draw(screen *ebiten.Image) {
	// Determine sub image index X (Pose)
	idxe := 2*c.FrameCount - 2
	idx := (utility.GetTickIndex() * c.FPS / utility.TickCount) % idxe
	if idx >= c.FrameCount {
		idx = idxe - idx
	}

	// Determine sub image index Y (Direction)
	idy := c.FrameDirectionMap[3]
	switch r := c.parent.GetRotation(); {
	case r < math.Pi*-3/4:
		idy = c.FrameDirectionMap[3]
	case r <= math.Pi*-1/4:
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

func (c *DrawAnimationCom) GetRectangleBounds(output *utility.RectangleF) {
	l := c.parent.GetLocation()
	s := c.parent.GetScale()

	output.MinX = l.X
	output.MinY = l.Y
	output.MaxX = l.X + float64(c.FrameSize.X)*s.X
	output.MaxY = l.Y + float64(c.FrameSize.Y)*s.Y
}

func (c *DrawAnimationCom) GetCircleBounds(output *utility.CircleF) {
	l := c.parent.GetLocation()
	s := c.parent.GetScale()
	sz := c.FrameSize.ToVector().Mul(s)

	output.OrgX = l.X + sz.X/2
	output.OrgY = l.Y + sz.Y/2
	output.Radius = math.Max(sz.X, sz.Y) / 2
}
