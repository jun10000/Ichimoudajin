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
	IsDebugMode        bool
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
		MaxReflectionCount: 1,
		DebugTextOffset:    utility.NewVector(3, -12),
		parent:             parent,
	}
}

func (c *MovementComponent) AddInput(normal utility.Vector, scale float64) {
	c.inputAccel = c.inputAccel.Add(normal.Normalize().MulF(scale))
}

func (c *MovementComponent) Tick() {
	// Update movement from input
	if !c.inputAccel.IsZero() {
		av := c.inputAccel.Normalize().MulF(c.Accel * utility.TickDuration)
		cr := c.inputAccel.Normalize().Cross(c.velocity).Dot(utility.NewVector3(0, 0, 1))
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
	ecs := []utility.Collider{c.parent}
	vn := c.velocity.Normalize()
	vl := c.velocity.Length()
	rl := vl * utility.TickDuration
	lv := utility.GetLevel()
	for i := 0; i <= c.MaxReflectionCount; i++ {
		ro := vn.MulF(rl)
		tr := lv.Trace(c.parent.GetColliderBounds(), ro, ecs, c.IsDebugMode)
		c.parent.AddLocation(tr.Offset)
		if tr.IsFirstHit && i == 0 {
			c.parent.AddLocation(vn.Negate())
			break
		}
		if !tr.IsHit {
			break
		}

		rl = tr.ROffset.Length()
		vn = vn.Reflect(tr.Normal, 0)
	}
	c.velocity = vn.MulF(vl)

	// Draw parent location
	if c.IsDebugMode {
		l := c.parent.GetLocation()
		o := c.DebugTextOffset
		utility.DrawDebugText(l.Add(o), l.String())
	}
}
