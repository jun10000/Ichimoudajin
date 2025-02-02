package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type AIControllerComponent struct {
	AIGridSize      utility.Vector
	IsDrawDebugPath bool

	parent         utility.Collider
	target         *MovementComponent
	pathfinding    *utility.AStar
	currentExcepts []utility.Collider
}

func NewAIControllerComponent(parent utility.Collider, target *MovementComponent) *AIControllerComponent {
	return &AIControllerComponent{
		AIGridSize:  utility.NewVector(32, 32),
		parent:      parent,
		target:      target,
		pathfinding: utility.NewAStar(),
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
	case c >= 2:
		dl = a.PFToRealLocation(pr[1], true)
		a.target.AddInput(dl.Sub(sl), 1)
	case c == 1:
		a.target.AddInput(dl.Sub(sl), 1)
	}

	if a.IsDrawDebugPath {
		for _, p := range pr {
			utility.DrawDebugRectangle(a.PFToRealLocation(p, false), a.AIGridSize, utility.ColorGreen)
		}
	}
}

func (a *AIControllerComponent) IsPointLocationValid(location utility.Point) bool {
	s := utility.GetGameInstance().ScreenSize
	l := a.PFToRealLocation(location, false)
	if l.X < 0 || l.Y < 0 || l.X >= float64(s.X) || l.Y >= float64(s.Y) {
		return false
	}

	b := utility.NewRectangleF(l, a.AIGridSize.SubF(1))
	r, _ := utility.GetLevel().Intersect(b, a.currentExcepts)
	return !r
}

func (a *AIControllerComponent) RealToPFLocation(real utility.Vector) utility.Point {
	return real.Div(a.AIGridSize).Floor()
}

func (a *AIControllerComponent) PFToRealLocation(pf utility.Point, isCenter bool) utility.Vector {
	r := pf.ToVector().Mul(a.AIGridSize)
	if isCenter {
		r = r.Add(a.AIGridSize.DivF(2))
	}
	return r
}
