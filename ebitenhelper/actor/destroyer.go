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
	DrawColor   color.Color
}

func NewDestroyer() *Destroyer {
	return &Destroyer{
		circle: utility.NewCircleF(0, 0, 0),

		GrowthValue: 1,
		MaxRadius:   120,
		DrawColor:   utility.ColorBlue,
	}
}

func (a *Destroyer) ZOrder() int {
	return utility.ZOrderEffect
}

func (a *Destroyer) Draw(screen *ebiten.Image) {
	if a.isShow {
		a.circle.Draw(screen, a.DrawColor, true)
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

	if ok, c, _ := utility.Intersect(l.MovableColliders, a.circle, excepts); ok {
		l.Remove(c)
	}
}
