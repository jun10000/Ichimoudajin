package component

import (
	"math/rand"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type AIControllerComponent struct {
	AIGridSize        utility.Vector
	LocationDeviation float64

	parent         utility.Collider
	target         *MovementComponent
	pathfinding    *utility.AStar
	currentExcepts []utility.Collider
}

func NewAIControllerComponent(parent utility.Collider, target *MovementComponent) *AIControllerComponent {
	return &AIControllerComponent{
		AIGridSize:        utility.NewVector(32, 32),
		LocationDeviation: 0.5,
		parent:            parent,
		target:            target,
		pathfinding:       utility.NewAStar(),
	}
}

func (a *AIControllerComponent) AITick() {
	c, ok := utility.GetLevel().InputReceivers[0].(utility.Collider)
	if ok {
		a.AIMoveToActor(c)
	}
}

func (a *AIControllerComponent) AIMoveToActor(dst utility.Collider) {
	src := a.parent
	a.currentExcepts = []utility.Collider{dst}
	for _, t := range utility.GetLevel().AITickers {
		if c, ok := t.(utility.Collider); ok {
			a.currentExcepts = append(a.currentExcepts, c)
		}
	}
	sl := src.GetColliderBounds().BoundingBox().CenterLocation()
	dl := dst.GetColliderBounds().BoundingBox().CenterLocation()

	pr := a.pathfinding.Run(a.RealToPFLocation(sl), a.RealToPFLocation(dl), a.IsPointLocationValid)
	switch c := len(pr); {
	case c > 2:
		dl1 := a.PFToRealLocation(pr[1], true, a.LocationDeviation)
		dl2 := a.PFToRealLocation(pr[2], true, a.LocationDeviation)
		dl = dl1.Add(dl2.Sub(dl1).DivF(2))
		a.target.AddInput(dl.Sub(sl), 1)
	case c == 2:
		dl = a.PFToRealLocation(pr[1], true, a.LocationDeviation)
		a.target.AddInput(dl.Sub(sl), 1)
	case c == 1:
		a.target.AddInput(dl.Sub(sl), 1)
	}

	if utility.IsShowDebugAIPath {
		for _, p := range pr {
			utility.DrawDebugRectangle(a.PFToRealLocation(p, false, 0), a.AIGridSize, utility.ColorGreen)
		}
	}
}

func (a *AIControllerComponent) IsPointLocationValid(location utility.Point) bool {
	s := utility.GetGameInstance().ScreenSize
	l := a.PFToRealLocation(location, false, 0)
	if l.X < 0 || l.Y < 0 || l.X >= float64(s.X) || l.Y >= float64(s.Y) {
		return false
	}

	b := utility.NewRectangleF(l, a.AIGridSize.SubF(1))
	r, _ := utility.GetLevel().Intersect(b, a.currentExcepts)
	return !r
}

func (a *AIControllerComponent) RealToPFLocation(realLocation utility.Vector) utility.Point {
	return realLocation.Div(a.AIGridSize).Floor()
}

func (a *AIControllerComponent) PFToRealLocation(pfLocation utility.Point, isCenter bool, deviation float64) utility.Vector {
	rr := a.AIGridSize.MulF((rand.Float64() - 0.5) * deviation)
	r := pfLocation.ToVector().Mul(a.AIGridSize).Add(rr)
	if isCenter {
		r = r.Add(a.AIGridSize.DivF(2))
	}
	return r
}
