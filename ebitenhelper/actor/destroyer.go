package actor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type Destroyer struct {
	isShow bool
	circle *utility.CircleF

	GrowthValue float64
	MaxRadius   float64
	BorderWidth float32
	BorderColor color.Color
	FillColor   color.Color
}

func NewDestroyer() *Destroyer {
	return &Destroyer{
		circle: utility.NewCircleF(0, 0, 0),

		GrowthValue: 1,
		MaxRadius:   120,
		BorderWidth: 2,
		BorderColor: utility.ColorLightBlue.ToRGBA(0xff),
		FillColor:   utility.ColorLightBlue.ToRGBA(0x20),
	}
}

func (a *Destroyer) ZOrder() int {
	return utility.ZOrderEffect
}

func (a *Destroyer) Draw(screen *ebiten.Image) {
	if a.isShow {
		a.circle.Draw(screen, a.BorderWidth, a.BorderColor, a.FillColor, true)
	}
}

func (a *Destroyer) Start(location utility.Vector) {
	a.isShow = true
	a.circle.OrgX = location.X
	a.circle.OrgY = location.Y
	a.circle.Radius = a.GrowthValue
}

func (a *Destroyer) Grow() {
	a.circle.Radius += a.GrowthValue
	if a.circle.Radius > a.MaxRadius {
		a.circle.Radius = a.MaxRadius
	}
}

func (a *Destroyer) Execute() {
	a.isShow = false

	l := utility.GetLevel()
	excepts := make(utility.Set[utility.MovableCollider])
	for _, p := range l.Players {
		excepts.Add(p)
	}

	if ok, cs, _ := utility.IntersectAll(l.MovableColliders, a.circle, excepts); ok {
		for _, c := range cs {
			l.Remove(c)
		}
	}
}
