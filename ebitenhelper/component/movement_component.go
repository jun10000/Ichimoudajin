package component

import (
	"math"

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
		av := c.inputAccel.Normalize().MulF(c.Accel * utility.TickDuration)
		cr := c.inputAccel.Normalize().Cross(c.velocity).Dot(utility.NewVector3(0, 0, 1))
		dv := c.inputAccel.Rotate(math.Pi / 2).Normalize().MulF(c.Decel * utility.TickDuration).ClampMax(math.Abs(cr))
		if cr < 0 {
			dv = dv.Negate()
		}

		c.velocity = c.velocity.Add(av).Add(dv).ClampMax(c.MaxSpeed)
		mover.SetRotation(utility.NewVector(0, 1).CrossingAngle(c.inputAccel))
	} else {
		dv := c.velocity.Normalize().MulF(utility.ClampFloat(c.Decel*utility.TickDuration, 0, c.velocity.Length()))
		c.velocity = c.velocity.Sub(dv)
	}
	c.inputAccel = utility.ZeroVector()

	// Collision test
	trm := c.velocity.MulF(utility.TickDuration)
	for i := 0; i < 10; i++ {
		tr := utility.GetLevel().Trace(mover.GetBounds(), trm, mover)
		mover.AddLocation(tr.Offset)
		if !tr.IsHit {
			break
		}

		c.velocity = c.velocity.Reflect(tr.Normal, 0)
		trm = tr.ROffset.Reflect(tr.Normal, 0)
	}
}
