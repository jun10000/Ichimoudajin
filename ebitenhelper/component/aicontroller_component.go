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
	t := utility.GetLevel().InputReceivers[0]
	a.AIMoveTo(t)
}

func (a *AIControllerComponent) AIMoveTo(dest any) {
	sl := a.parent.GetLocation()
	gt, ok := dest.(utility.Transformer)
	if !ok {
		return
	}
	gl := gt.GetLocation()

	sp := sl.Div(a.AIGridSize).Floor()
	gp := gl.Div(a.AIGridSize).Floor()
	a.currentExcepts = []utility.Collider{a.parent, dest.(utility.Collider)}

	pr := a.pathfinding.Run(sp, gp, a.IsPointLocationValid)
	switch c := len(pr); {
	case c >= 2:
		// to do
	case c == 1:
		a.target.AddInput(gl.Sub(sl), 1)
	default:
	}

	// for debug
	for _, p := range pr {
		utility.DrawDebugRectangle(p.ToVector().Mul(a.AIGridSize), a.AIGridSize, utility.ColorGreen)
	}
}

func (a *AIControllerComponent) IsPointLocationValid(location utility.Point) bool {
	tl := location.ToVector().Mul(a.AIGridSize)
	rc := utility.NewRectangleF(tl, a.AIGridSize)
	tr := utility.GetLevel().Trace(rc, utility.ZeroVector(), a.currentExcepts)
	return !tr.IsHit
}
