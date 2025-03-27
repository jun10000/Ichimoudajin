package utility

type CallTimer struct {
	isRunning     bool
	runningFunc   func()
	tickGoalIndex int
}

func NewCallTimer() *CallTimer {
	return &CallTimer{}
}

func (c *CallTimer) Tick() {
	if !c.isRunning {
		return
	}

	if GetTickIndex() >= c.tickGoalIndex {
		c.runningFunc()
		c.isRunning = false
	}
}

func (c *CallTimer) StartCallTimer(f func(), seconds float32) {
	if c.isRunning {
		return
	}

	c.isRunning = true
	c.runningFunc = f
	c.tickGoalIndex = GetTickIndex() + int(seconds*TickCount)
}

func (c *CallTimer) StopCallTimer() {
	c.isRunning = false
}
