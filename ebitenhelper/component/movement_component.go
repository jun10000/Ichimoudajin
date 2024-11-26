package component

import (
	"math"

	"github.com/jun10000/Ichimoudajin/ebitenhelper"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type MovementComponent struct {
	Accel    float64
	Decel    float64
	MaxSpeed float64

	CurrentVelocity utility.Vector

	t_InputAccel utility.Vector
}

func NewMovementComponent() *MovementComponent {
	return &MovementComponent{
		Accel:    8000,
		Decel:    8000,
		MaxSpeed: 200,
	}
}

func (c *MovementComponent) AddInput(normal utility.Vector, scale float64) {
	c.t_InputAccel = c.t_InputAccel.Add(normal.Normalize().MulF(scale))
}

func (c *MovementComponent) Tick(mover utility.Mover) {
	if c.t_InputAccel.X != 0 || c.t_InputAccel.Y != 0 {
		c.CurrentVelocity = c.CurrentVelocity.Add(c.t_InputAccel.MulF(c.Accel * utility.TickDuration))
		if c.CurrentVelocity.Length() > c.MaxSpeed {
			c.CurrentVelocity = c.CurrentVelocity.Normalize().MulF(c.MaxSpeed)
		}
		mover.SetRotation(utility.NewVector(0, 1).CrossingAngle(c.t_InputAccel))
	} else {
		decelspeed := c.CurrentVelocity.Normalize().MulF(c.Decel * utility.TickDuration)
		if math.Abs(decelspeed.X) > math.Abs(c.CurrentVelocity.X) {
			decelspeed.X = c.CurrentVelocity.X
		}
		if math.Abs(decelspeed.Y) > math.Abs(c.CurrentVelocity.Y) {
			decelspeed.Y = c.CurrentVelocity.Y
		}
		c.CurrentVelocity = c.CurrentVelocity.Sub(decelspeed)
	}

	traceoffset := c.CurrentVelocity.MulF(utility.TickDuration)
	traceresult := ebitenhelper.GetLevel().Trace(mover.GetBounds(), traceoffset, mover)
	if traceresult.IsHit {
		c.CurrentVelocity = c.CurrentVelocity.Reflect(traceresult.Normal, 0)
		mover.SetLocation(mover.GetLocation().Add(traceresult.Offset))
		traceoffset = c.CurrentVelocity.MulF(utility.TickDuration * traceresult.RDistanceRatio)
		traceresult = ebitenhelper.GetLevel().Trace(mover.GetBounds(), traceoffset, mover)
		if traceresult.IsHit {
			mover.SetLocation(mover.GetLocation().Add(traceresult.Offset))
		} else {
			mover.SetLocation(mover.GetLocation().Add(traceoffset))
		}
	} else {
		mover.SetLocation(mover.GetLocation().Add(traceoffset))
	}

	c.t_InputAccel = utility.ZeroVector()
}
