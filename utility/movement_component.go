package utility

import (
	"math"
)

type MovementComponent struct {
	Parent   *Pawn
	Accel    float64
	Decel    float64
	MaxSpeed float64

	CurrentVelocity Vector

	t_InputAccel Vector
}

func NewMovementComponent(pawn *Pawn) *MovementComponent {
	return &MovementComponent{
		Parent:   pawn,
		Accel:    8000,
		Decel:    8000,
		MaxSpeed: 400,
	}
}

func (c *MovementComponent) AddInput(normal Vector, scale float64) {
	c.t_InputAccel = c.t_InputAccel.Add(normal.Normalize().MulF(scale))
}

func (c *MovementComponent) Tick() {
	if c.t_InputAccel.X != 0 || c.t_InputAccel.Y != 0 {
		c.CurrentVelocity = c.CurrentVelocity.Add(c.t_InputAccel.MulF(c.Accel * TickDuration))
		if c.CurrentVelocity.GetLength() > c.MaxSpeed {
			c.CurrentVelocity = c.CurrentVelocity.Normalize().MulF(c.MaxSpeed)
		}
	} else {
		decelspeed := c.CurrentVelocity.Normalize().MulF(c.Decel * TickDuration)
		if math.Abs(decelspeed.X) > math.Abs(c.CurrentVelocity.X) {
			decelspeed.X = c.CurrentVelocity.X
		}
		if math.Abs(decelspeed.Y) > math.Abs(c.CurrentVelocity.Y) {
			decelspeed.Y = c.CurrentVelocity.Y
		}
		c.CurrentVelocity = c.CurrentVelocity.Sub(decelspeed)
	}

	c.Parent.Location = c.Parent.Location.Add(c.CurrentVelocity.MulF(TickDuration))
	c.t_InputAccel = ZeroVector()
}
