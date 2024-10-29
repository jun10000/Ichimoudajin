package utility

import "math"

type MovementComponent struct {
	parent     *Pawn
	AccelValue float64
	DecelValue float64
	MaxSpeed   float64

	CurrentVelocity Vector

	Temp_Accel Vector
}

func NewMovementComponent(pawn *Pawn) *MovementComponent {
	return &MovementComponent{
		parent:     pawn,
		AccelValue: 8000,
		DecelValue: 8000,
		MaxSpeed:   400,
	}
}

func (c *MovementComponent) AddInput(normal Vector, scale float64) {
	accel := normal.Normalize().MulF(scale)
	c.Temp_Accel = c.Temp_Accel.Add(accel)
}

func (c *MovementComponent) Tick() {
	if c.Temp_Accel.X != 0 || c.Temp_Accel.Y != 0 {
		c.CurrentVelocity = c.CurrentVelocity.Add(c.Temp_Accel.MulF(c.AccelValue * TickDuration))
		if c.CurrentVelocity.GetLength() > c.MaxSpeed {
			c.CurrentVelocity = c.CurrentVelocity.Normalize().MulF(c.MaxSpeed)
		}
	} else {
		decelspeed := c.CurrentVelocity.Normalize().MulF(c.DecelValue * TickDuration)
		if math.Abs(decelspeed.X) > math.Abs(c.CurrentVelocity.X) {
			decelspeed.X = c.CurrentVelocity.X
		}
		if math.Abs(decelspeed.Y) > math.Abs(c.CurrentVelocity.Y) {
			decelspeed.Y = c.CurrentVelocity.Y
		}
		c.CurrentVelocity = c.CurrentVelocity.Sub(decelspeed)
	}

	c.parent.Location = c.parent.Location.Add(c.CurrentVelocity.MulF(TickDuration))
	c.Temp_Accel = ZeroVector()
}
