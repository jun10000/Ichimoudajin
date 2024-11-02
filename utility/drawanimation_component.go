package utility

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawAnimationComponent struct {
	parent            *Pawn
	Image             *ebiten.Image
	FrameCount        int
	FrameSize         Point
	FPS               int
	FrameDirectionMap []int // Front, Left, Right, Back

	currentTickIndex int
}

func NewDrawAnimationComponent(parentActor *Pawn) *DrawAnimationComponent {
	return &DrawAnimationComponent{
		parent:            parentActor,
		FrameCount:        3,
		FrameSize:         NewPoint(32, 32),
		FPS:               4,
		FrameDirectionMap: []int{0, 1, 2, 3},
	}
}

func (c *DrawAnimationComponent) Tick() {
	if c.Image == nil {
		return
	}

	c.currentTickIndex++
}

func (c *DrawAnimationComponent) Draw(screen *ebiten.Image) {
	if c.Image == nil {
		return
	}

	excount := 2*c.FrameCount - 2
	index := (c.currentTickIndex * c.FPS / TickCount) % excount
	if index >= c.FrameCount {
		index = excount - index
	}

	direction := c.FrameDirectionMap[3]
	switch r := c.parent.Rotation.Get(); {
	case r < math.Pi*-3/4:
		direction = c.FrameDirectionMap[3]
	case r < math.Pi*-1/4:
		direction = c.FrameDirectionMap[2]
	case r < math.Pi*1/4:
		direction = c.FrameDirectionMap[0]
	case r <= math.Pi*3/4:
		direction = c.FrameDirectionMap[1]
	}

	location := NewPoint(
		index*c.FrameSize.X,
		direction*c.FrameSize.Y,
	)

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(c.parent.Scale.X, c.parent.Scale.Y)
	o.GeoM.Translate(c.parent.Location.X, c.parent.Location.Y)
	screen.DrawImage(GetSubImage(c.Image, location, c.FrameSize), o)
}
