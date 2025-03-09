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

func (c *MovementComponent) AddLocation(offset utility.Vector) (rOnHitDistance int, rOffset utility.Vector, rNormal *utility.Vector, rIsHit bool) {
	bounds := c.parent.GetMainColliderBounds()
	excepts := make(utility.Set[utility.Collider])
	excepts.Add(c.parent)
	thd, to, tn, th := utility.Trace(utility.GetLevel().Colliders, bounds, offset, excepts)

	if th {
		if thd == 0 { // Force back location
			c.addLocationForce(*tn)
		} else if (thd - 1) > utility.MovementInvalidDistance {
			tol, ton := to.Decompose()
			lo := ton.MulF(tol - float64(utility.MovementInvalidDistance))
			c.addLocationForce(lo)
		}
	} else {
		c.addLocationForce(to)
	}

	return thd, to, tn, th
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
		_, tro, trn, trh := c.AddLocation(ro)
		if !trh {
			break
		}

		rl = ro.Sub(tro).Length()
		vn = vn.Reflect(*trn, 0)
	}
	c.velocity = vn.MulF(vl)

	// Debug
	utility.DrawDebugLocation(c.parent.GetLocation())
}
