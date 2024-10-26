package gamemap

import (
	"github.com/jun10000/Ichimoudajin/utility"
)

var Stage1 = &utility.Level{}

func init() {
	Stage1.Actors = []*utility.Actor{
		utility.NewActor("rectangle100x100.png", 100, 200, 0, 1, 1),
	}
	Stage1.Pawns = []*utility.Actor{
		utility.NewActor("triangle100x100.png", 300, 400, 0, 1, 1),
	}
}
