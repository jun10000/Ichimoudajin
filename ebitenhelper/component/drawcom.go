package component

type DrawCom struct {
	isVisible bool
}

func NewDrawCom(isVisible bool) *DrawCom {
	return &DrawCom{
		isVisible: isVisible,
	}
}

func (c *DrawCom) GetVisibility() bool {
	return c.isVisible
}

func (c *DrawCom) SetVisibility(isVisible bool) {
	c.isVisible = isVisible
}
