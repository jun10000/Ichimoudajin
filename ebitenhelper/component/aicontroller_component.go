package component

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type AIControllerComponent struct {
	AIGridSize utility.Vector

	parent      utility.Transformer
	target      *MovementComponent
	pathfinding *utility.AStar
}

func NewAIControllerComponent(parent utility.Transformer, target *MovementComponent) *AIControllerComponent {
	return &AIControllerComponent{
		AIGridSize:  utility.NewVector(64, 64),
		parent:      parent,
		target:      target,
		pathfinding: utility.NewAStar(),
	}
}

func (a *AIControllerComponent) AITick() {
	gl := utility.GetLevel().InputReceivers[0].GetLocation()
	a.AIMoveTo(gl)
}

func (a *AIControllerComponent) AIMoveTo(goal utility.Vector) {
	start := a.parent.GetLocation()
	sp := start.Div(a.AIGridSize).Floor()
	gp := goal.Div(a.AIGridSize).Floor()

	pr := a.pathfinding.Run(sp, gp, a.IsPointLocationValid)
	switch c := len(pr); {
	case c >= 2:
		// to do
	case c == 1:
		a.target.AddInput(goal.Sub(start), 1)
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
	tr := utility.GetLevel().Trace(rc, utility.ZeroVector(), nil)
	return tr.IsHit
}
