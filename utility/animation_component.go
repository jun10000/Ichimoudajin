package utility

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationComponent struct {
	Parent            *Pawn
	FrameCount        int
	FrameSize         Point
	FPS               int
	FrameDirectionMap []int // Front, Left, Right, Back

	currentTickIndex int
}

func NewAnimationComponent(pawn *Pawn) *AnimationComponent {
	return &AnimationComponent{
		Parent:            pawn,
		FrameCount:        3,
		FrameSize:         NewPoint(32, 32),
		FPS:               4,
		FrameDirectionMap: []int{0, 1, 2, 3},
	}
}

func (c *AnimationComponent) Tick() {
	c.currentTickIndex++
}

func (c *AnimationComponent) Draw(screen *ebiten.Image) {
	excount := 2*c.FrameCount - 2
	index := (c.currentTickIndex * c.FPS / TickCount) % excount
	if index >= c.FrameCount {
		index = excount - index
	}

	direction := c.FrameDirectionMap[3]
	switch r := c.Parent.Rotation.Get(); {
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
	o.GeoM.Scale(c.Parent.Scale.X, c.Parent.Scale.Y)
	o.GeoM.Translate(c.Parent.Location.X, c.Parent.Location.Y)
	screen.DrawImage(GetSubImage(c.Parent.Image, location, c.FrameSize), o)
}
