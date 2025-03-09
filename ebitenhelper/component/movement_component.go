package component

import (
	"math"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type MovementComponent struct {
	Accel    float64
	Decel    float64
	MaxSpeed float64

	parent     utility.Collider
	velocity   utility.Vector
	inputAccel utility.Vector
}

func NewMovementComponent(parent utility.Collider) *MovementComponent {
	return &MovementComponent{
		Accel:    8000,
		Decel:    8000,
		MaxSpeed: 200,
		parent:   parent,
	}
}

func (c *MovementComponent) AddInput(normal utility.Vector, scale float64) {
	c.inputAccel = c.inputAccel.Add(normal.Normalize().MulF(scale))
}

func (c *MovementComponent) addLocationForce(offset utility.Vector) {
	c.parent.SetLocation(c.parent.GetLocation().Add(offset))
}

func (c *MovementComponent) AddLocation(offset utility.Vector) *utility.TraceResult {
	bounds := c.parent.GetMainColliderBounds()
	excepts := make(utility.Set[utility.Collider])
	excepts.Add(c.parent)
	r := utility.Trace(utility.GetLevel().Colliders, bounds, offset, excepts)

	if r.IsHit {
		if r.HitOffsetD == 0 { // Force back location
			c.addLocationForce(*r.HitNormal)
		} else if r.TraceoffsetD > utility.MovementInvalidDistance {
			tol, ton := r.TraceOffset.Decompose()
			lo := ton.MulF(tol - float64(utility.MovementInvalidDistance))
			c.addLocationForce(lo)
		}
	} else {
		c.addLocationForce(r.InputOffset)
	}

	return r
}

func (c *MovementComponent) Tick() {
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
	vl, vn := c.velocity.Decompose()
	rl := vl * utility.TickDuration
	for range utility.MovementMaxReflectionCount + 1 {
		ro := vn.MulF(rl)
		tr := c.AddLocation(ro)
		if !tr.IsHit {
			break
		}

		rl = ro.Sub(tr.TraceOffset).Length()
		vn = vn.Reflect(*tr.HitNormal, 0)
	}
	c.velocity = vn.MulF(vl)

	// Debug
	utility.DrawDebugLocation(c.parent.GetLocation())
}
