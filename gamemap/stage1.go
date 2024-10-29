package gamemap

import (
	"github.com/jun10000/Ichimoudajin/utility"
)

var Stage1 *utility.Level

func init() {
	acter_rect := utility.NewActor("rectangle100x100.png")
	acter_rect.Location = utility.NewVector(100, 200)
	pawn_tri := utility.NewPawn("triangle100x100.png")
	pawn_tri.Location = utility.NewVector(300, 400)

	Stage1 = utility.NewLevel()
	Stage1.Drawers = []utility.Drawer{acter_rect, pawn_tri}
	Stage1.KeyReceivers = []utility.KeyReceiver{pawn_tri}
	Stage1.Tickers = []utility.Ticker{pawn_tri}
}
