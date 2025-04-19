package actor

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/widget"
)

type DestroyerStatus int

const (
	DestroyerStatusDisable DestroyerStatus = iota
	DestroyerStatusGrowing
	DestroyerStatusShrinking
)

type Destroyer struct {
	*component.ActorCom
	*component.DrawCom
	status  DestroyerStatus
	circle  *utility.CircleF
	targets []utility.MovableCollider
	points  int
	widget  *widget.WidgetText

	GrowSpeed   float64
	ShrinkSpeed float64
	MaxRadius   float64
	BorderWidth float32
	BorderColor color.Color
	FillColor   color.Color
}

func (g ActorGeneratorStruct) NewDestroyer(options *NewActorOptions) *Destroyer {
	return &Destroyer{
		ActorCom: component.NewActorCom(options.Name),
		DrawCom:  component.NewDrawCom(options.IsVisible),
		status:   DestroyerStatusDisable,
		circle:   utility.NewCircleF(0, 0, 0),

		GrowSpeed:   1,
		ShrinkSpeed: 2,
		MaxRadius:   120,
		BorderWidth: 2,
		BorderColor: utility.ColorOrange,
		FillColor:   utility.ColorOrange.ToNRGBA(0x40),
	}
}

func (a *Destroyer) ZOrder() int {
	return utility.ZOrderEffect
}

func (a *Destroyer) BeginPlay() {
	a.widget = widget.GetWidgetObjectByNameP[*widget.WidgetText]("mainwidget", "PointText")
}

func (a *Destroyer) Tick() {
	switch a.status {
	case DestroyerStatusGrowing:
		a.circle.Radius += a.GrowSpeed
		if a.circle.Radius > a.MaxRadius {
			a.circle.Radius = a.MaxRadius
		}
	case DestroyerStatusShrinking:
		a.circle.Radius -= a.ShrinkSpeed
		if a.circle.Radius <= 0 {
			a.circle.Radius = 0
		}

		tTrashes := make([]utility.MovableCollider, 0)
		cLocation := a.circle.CenterLocation()
		cRadius := a.circle.Radius
		for _, t := range a.targets {
			tCircle := t.GetRealFirstBounds().ToCircle()
			tLocation := tCircle.CenterLocation()
			tRadius := tCircle.Radius
			if cRadius <= tRadius {
				tTrashes = append(tTrashes, t)
				continue
			}

			dtLength, dtNormal := cLocation.Sub(tLocation).Decompose()
			dtLength -= cRadius - tRadius
			if dtLength > 0 {
				t.AddLocation(dtNormal.MulF(dtLength))
			}
		}

		lv := utility.GetLevel()
		for _, t := range tTrashes {
			a.targets = utility.RemoveSliceItem(a.targets, t)
			lv.Remove(t)
			a.points += 1
		}

		if cRadius == 0 {
			a.status = DestroyerStatusDisable
		}
	}

	a.ApplyPointsToWidget()
}

func (a *Destroyer) ApplyPointsToWidget() {
	a.widget.Text = fmt.Sprintf("%d", a.points)
}

func (a *Destroyer) Draw(screen *ebiten.Image) {
	if a.status != DestroyerStatusDisable {
		a.circle.Draw(screen, a.BorderWidth, a.BorderColor, a.FillColor, true)
	}
}

func (a *Destroyer) Start(location utility.Vector) {
	if a.status != DestroyerStatusDisable {
		return
	}

	a.circle.OrgX = location.X
	a.circle.OrgY = location.Y
	a.circle.Radius = a.GrowSpeed

	a.status = DestroyerStatusGrowing
}

func (a *Destroyer) Finish() {
	if a.status != DestroyerStatusGrowing {
		return
	}

	l := utility.GetLevel()
	excepts := make(utility.Set[utility.MovableCollider])
	for _, p := range l.Players {
		excepts.Add(p)
	}
	_, a.targets, _ = utility.IntersectAll(l.MovableColliders, a.circle, excepts)
	for _, t := range a.targets {
		t.DisableBounds()
	}

	a.status = DestroyerStatusShrinking
}
