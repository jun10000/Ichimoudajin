package component

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type MovementComponent struct {
	Accel    float64
	Decel    float64
	MaxSpeed float64

	velocity   utility.Vector
	inputAccel utility.Vector
}

func NewMovementComponent() *MovementComponent {
	return &MovementComponent{
		Accel:    8000,
		Decel:    8000,
		MaxSpeed: 200,
	}
}

func (c *MovementComponent) AddInput(normal utility.Vector, scale float64) {
	c.inputAccel = c.inputAccel.Add(normal.Normalize().MulF(scale))
}

func (c *MovementComponent) Tick(mover utility.Mover) {
	// Update movement from input
	if !c.inputAccel.IsZero() {
		as := c.Accel * utility.TickDuration
		av := c.inputAccel.Normalize().MulF(as)
		c.velocity = c.velocity.Add(av).ClampMax(c.MaxSpeed)
		mover.SetRotation(utility.NewVector(0, 1).CrossingAngle(c.inputAccel))
	} else {
		ds := utility.ClampFloat(c.Decel*utility.TickDuration, 0, c.velocity.Length())
		dv := c.velocity.Normalize().MulF(ds)
		c.velocity = c.velocity.Sub(dv)
	}
	c.inputAccel = utility.ZeroVector()

	// Collision test
	trm := c.velocity.MulF(utility.TickDuration)
	for i := 0; i < 10; i++ {
		tr := ebitenhelper.GetLevel().Trace(mover.GetBounds(), trm, mover)
		mover.AddLocation(tr.Offset)
		if !tr.IsHit {
			break
		}

		c.velocity = c.velocity.Reflect(tr.Normal, 0)
		trm = tr.ROffset.Reflect(tr.Normal, 0)
	}
}
