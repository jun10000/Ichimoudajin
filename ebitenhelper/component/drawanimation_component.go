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

	currentTickIndex int
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

	c.currentTickIndex++
}

func (c *DrawAnimationComponent) Draw(screen *ebiten.Image, transformer utility.Transformer) {
	if c.Source == nil {
		return
	}

	excount := 2*c.FrameCount - 2
	index := (c.currentTickIndex * c.FPS / utility.TickCount) % excount
	if index >= c.FrameCount {
		index = excount - index
	}

	direction := c.FrameDirectionMap[3]
	switch r := transformer.GetRotation(); {
	case r < math.Pi*-3/4:
		direction = c.FrameDirectionMap[3]
	case r < math.Pi*-1/4:
		direction = c.FrameDirectionMap[2]
	case r < math.Pi*1/4:
		direction = c.FrameDirectionMap[0]
	case r <= math.Pi*3/4:
		direction = c.FrameDirectionMap[1]
	}

	location := transformer.GetLocation()
	scale := transformer.GetScale()
	sublocation := utility.NewPoint(
		index*c.FrameSize.X,
		direction*c.FrameSize.Y,
	)

	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(scale.X, scale.Y)
	o.GeoM.Translate(location.X, location.Y)
	screen.DrawImage(utility.GetSubImage(c.Source, sublocation, c.FrameSize), o)
}
