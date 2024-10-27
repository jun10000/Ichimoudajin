package utility

type MovementComponent struct {
	AccelSpeed float64
	DecelSpeed float64
	MaxSpeed   float64

	CurrentVelocity Vector

	Temp_InputVelocity Vector
}

func NewMovementComponent() *MovementComponent {
	return &MovementComponent{
		AccelSpeed: 5,
		DecelSpeed: 10,
		MaxSpeed:   5,
	}
}

func (c *MovementComponent) AddInput(normal Vector, scale float64) {
	n := normal.Normalize()
	v2 := n.MultiplyFloat(scale * c.AccelSpeed)
	c.Temp_InputVelocity = c.Temp_InputVelocity.Add(v2)
}

func (c *MovementComponent) Event_Tick(parent *Actor) {
	if c.Temp_InputVelocity.X == 0 && c.Temp_InputVelocity.Y == 0 {
		l := c.CurrentVelocity.GetLength() - c.DecelSpeed/60
		n := c.CurrentVelocity.Normalize()
		if l < 0 {
			l = 0
		}
		c.CurrentVelocity = n.MultiplyFloat(l)
	} else {
		c.CurrentVelocity = c.CurrentVelocity.Add(c.Temp_InputVelocity.DivideFloat(60))
	}
	parent.Location = parent.Location.Add(c.CurrentVelocity.Clamp(0, c.MaxSpeed))
	c.Temp_InputVelocity = NewVector(0, 0)
}
