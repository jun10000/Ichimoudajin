package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type AIControllerComponent struct {
	AIGridSize utility.Vector

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
	sl := src.GetColliderBounds().BoundingBox().CenterLocation()
	sp := sl.Div(a.AIGridSize).Floor()

	dl := dst.GetColliderBounds().BoundingBox().CenterLocation()
	dp := dl.Div(a.AIGridSize).Floor()

	a.currentExcepts = []utility.Collider{src, dst}

	pr := a.pathfinding.Run(sp, dp, a.IsPointLocationValid)
	switch c := len(pr); {
	case c >= 2:
		dl = pr[1].ToVector().Mul(a.AIGridSize).Add(a.AIGridSize.DivF(2))
		a.target.AddInput(dl.Sub(sl), 1)
	case c == 1:
		a.target.AddInput(dl.Sub(sl), 1)
	default:
	}

	// for debug
	for _, p := range pr {
		utility.DrawDebugRectangle(p.ToVector().Mul(a.AIGridSize), a.AIGridSize, utility.ColorGreen)
	}
}

func (a *AIControllerComponent) IsPointLocationValid(location utility.Point) bool {
	ss := utility.GetGameInstance().ScreenSize
	tl := location.ToVector().Mul(a.AIGridSize)
	if tl.X < 0 || tl.Y < 0 || tl.X > float64(ss.X) || tl.Y > float64(ss.Y) {
		return false
	}

	rc := utility.NewRectangleF(tl, a.AIGridSize)
	tr := utility.GetLevel().Trace(rc, utility.ZeroVector(), a.currentExcepts)
	return !tr.IsHit
}
