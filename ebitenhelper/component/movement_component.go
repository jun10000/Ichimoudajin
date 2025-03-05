package component

import (
	"math"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type MovementComponent struct {
	Accel              float64
	Decel              float64
	MaxSpeed           float64
	MaxReflectionCount int
	DebugTextOffset    utility.Vector

	parent     utility.Collider
	velocity   utility.Vector
	inputAccel utility.Vector
}

func NewMovementComponent(parent utility.Collider) *MovementComponent {
	return &MovementComponent{
		Accel:              8000,
		Decel:              8000,
		MaxSpeed:           200,
		MaxReflectionCount: 3,
		DebugTextOffset:    utility.NewVector(3, -12),
		parent:             parent,
	}
}

func (c *MovementComponent) AddInput(normal utility.Vector, scale float64) {
	c.inputAccel = c.inputAccel.Add(normal.Normalize().MulF(scale))
}

func (c *MovementComponent) AddLocation(offset utility.Vector) (rOffset utility.Vector, rNormal utility.Vector, rIsHit bool) {
	bounds := c.parent.GetMainColliderBounds()
	excepts := make(utility.Set[utility.Collider])
	excepts.Add(c.parent)
	o, n, ok := utility.GetLevel().Trace(bounds, offset, excepts)
	c.parent.SetLocation(c.parent.GetLocation().Add(o))
	return o, n, ok
}

func (c *MovementComponent) Tick() {
	// Update movement from input
	if !c.inputAccel.IsZero() {
		av := c.inputAccel.Normalize().MulF(c.Accel * utility.TickDuration)
		cr := c.inputAccel.Normalize().CrossZ(c.velocity)
		dv := c.inputAccel.Rotate(math.Pi / 2).Normalize().MulF(c.Decel * utility.TickDuration).ClampMax(math.Abs(cr))
		if cr < 0 {
			dv = dv.Negate()
		}

		c.velocity = c.velocity.Add(av).Add(dv).ClampMax(c.MaxSpeed)
		c.parent.SetRotation(utility.NewVector(0, 1).CrossingAngle(c.inputAccel))
	} else {
		dv := c.velocity.Normalize().MulF(utility.ClampFloat(c.Decel*utility.TickDuration, 0, c.velocity.Length()))
		c.velocity = c.velocity.Sub(dv)
	}
	c.inputAccel = utility.ZeroVector()

	// Collision test
	vn := c.velocity.Normalize()
	vl := c.velocity.Length()
	rl := vl * utility.TickDuration
	for i := 0; i <= c.MaxReflectionCount; i++ {
		ro := vn.MulF(rl)
		tro, trn, trok := c.AddLocation(ro)
		if !trok {
			break
		}

		rl = ro.Sub(tro).Length()
		vn = vn.Reflect(trn, 0)
	}
	c.velocity = vn.MulF(vl)

	// Draw parent location
	if utility.DebugIsShowMoverLocation {
		l := c.parent.GetLocation()
		o := c.DebugTextOffset
		utility.DrawDebugText(l.Add(o), l.String())
	}
}
