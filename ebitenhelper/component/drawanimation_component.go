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

func (c *DrawAnimationComponent) Draw(screen *ebiten.Image, transformer utility.Transformer) {
	if c.Source == nil {
		return
	}

	// Determine sub image index X (Pose)
	idxe := 2*c.FrameCount - 2
	idx := (c.tickIndex * c.FPS / utility.TickCount) % idxe
	if idx >= c.FrameCount {
		idx = idxe - idx
	}

	// Determine sub image index Y (Direction)
	idy := c.FrameDirectionMap[3]
	switch r := transformer.GetRotation(); {
	case r < math.Pi*-3/4:
		idy = c.FrameDirectionMap[3]
	case r < math.Pi*-1/4:
		idy = c.FrameDirectionMap[2]
	case r < math.Pi*1/4:
		idy = c.FrameDirectionMap[0]
	case r <= math.Pi*3/4:
		idy = c.FrameDirectionMap[1]
	}

	il := utility.NewPoint(idx*c.FrameSize.X, idy*c.FrameSize.Y)
	al := transformer.GetLocation()
	as := transformer.GetScale()
	als := []utility.Vector{al}
	if utility.GetLevel().IsLooping {
		ss := utility.GetGameInstance().ScreenSize.ToVector()
		als = append(als,
			al.Add(ss.Mul(utility.NewVector(-1, -1))),
			al.Add(ss.Mul(utility.NewVector(0, -1))),
			al.Add(ss.Mul(utility.NewVector(1, -1))),
			al.Add(ss.Mul(utility.NewVector(-1, 0))),
			al.Add(ss.Mul(utility.NewVector(1, 0))),
			al.Add(ss.Mul(utility.NewVector(-1, 1))),
			al.Add(ss.Mul(utility.NewVector(0, 1))),
			al.Add(ss.Mul(utility.NewVector(1, 1))),
		)
	}

	// Draw images
	for _, l := range als {
		o := &ebiten.DrawImageOptions{}
		o.GeoM.Scale(as.X, as.Y)
		o.GeoM.Translate(l.X, l.Y)
		screen.DrawImage(utility.GetSubImage(c.Source, il, c.FrameSize), o)
	}
}
