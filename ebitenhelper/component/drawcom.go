package component

type DrawCom struct {
	isVisible bool
}

func NewDrawCom() *DrawCom {
	return &DrawCom{
		isVisible: true,
	}
}

func (c *DrawCom) GetVisibility() bool {
	return c.isVisible
}

func (c *DrawCom) SetVisibility(isVisible bool) {
	c.isVisible = isVisible
}
