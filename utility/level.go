package utility

type Level struct {
	Drawers      []Drawer
	KeyReceivers []KeyReceiver
	Tickers      []Ticker
}

func NewLevel() *Level {
	return &Level{}
}
