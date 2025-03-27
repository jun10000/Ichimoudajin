package component

import (
	"math"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type MovementCom struct {
	parent     utility.MovableCollider
	velocity   utility.Vector
	inputAccel utility.Vector

	Accel    float64
	Decel    float64
	MaxSpeed float64
}

func NewMovementCom(parent utility.MovableCollider) *MovementCom {
	return &MovementCom{
		parent: parent,

		Accel:    8000,
		Decel:    8000,
		MaxSpeed: 200,
	}
}

func (c *MovementCom) AddInput(normal utility.Vector, scale float64) {
	c.inputAccel = c.inputAccel.Add(normal.Normalize().MulF(scale))
}

func (c *MovementCom) addLocationForce(offset utility.Vector) {
	c.parent.SetLocation(c.parent.GetLocation().Add(offset))
}

func (c *MovementCom) AddLocation(offset utility.Vector) *utility.TraceResult[utility.Collider] {
	bounds := c.parent.GetFirstBounds()
	if bounds == nil {
		c.addLocationForce(offset)
		return utility.NewTraceResult[utility.Collider](offset)
	}

	excepts := make(utility.Set[utility.Collider])
	excepts.Add(c.parent)
	r := utility.Trace(utility.GetLevel().Colliders, bounds, offset, excepts)

	if r.IsHit { // Force back location
		if r.IsFirstHit {
			c.addLocationForce(*r.HitNormal)
		} else {
			c.addLocationForce(r.TraceOffset.Add(*r.HitNormal))
		}

		c.parent.ReceiveHit(r)
	} else {
		c.addLocationForce(r.InputOffset)
	}

	return r
}

func (c *MovementCom) Tick() {
	// Update movement from input
	if !c.inputAccel.IsZero() {
		ia := c.inputAccel.Normalize()
		av := ia.MulF(c.Accel * utility.TickDuration)
		cr := ia.CrossZ(c.velocity)
		dv := ia.Rotate(math.Pi / 2).MulF(c.Decel * utility.TickDuration).ClampMax(math.Abs(cr))
		if cr < 0 {
			dv = dv.Negate()
		}

		c.velocity = c.velocity.Add(av).Add(dv).ClampMax(c.MaxSpeed)
		c.parent.SetRotation(utility.DownVectorPtr().CrossingAngle(ia))
	} else {
		vl, vn := c.velocity.Decompose()
		dv := vn.MulF(utility.ClampFloat(c.Decel*utility.TickDuration, 0, vl))
		c.velocity = c.velocity.Sub(dv)
	}
	c.inputAccel = utility.ZeroVector()

	// Collision test
	velocityL, velocityN := c.velocity.Decompose()
	remD := velocityL * utility.TickDuration
	for range utility.MovementMaxReflectionCount + 1 {
		tr := c.AddLocation(velocityN.MulF(remD))
		if !tr.IsHit {
			break
		}

		remD -= float64(tr.TraceoffsetD)
		if remD <= 0 {
			break
		}

		velocityN = velocityN.Reflect(*tr.HitNormal, 0)
	}
	c.velocity = velocityN.MulF(velocityL)

	// Debug
	utility.DrawDebugLocation(c.parent.GetLocation())
}
