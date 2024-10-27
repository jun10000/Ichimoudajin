package gamemap

import (
	"github.com/jun10000/Ichimoudajin/utility"
)

var Stage1 *utility.Level

func init() {
	rect := utility.NewActor("rectangle100x100.png")
	rect.Location = utility.NewVector(100, 200)
	tri := utility.NewActor("triangle100x100.png")
	tri.Location = utility.NewVector(300, 400)

	Stage1 = utility.NewLevel()
	Stage1.Actors = []*utility.Actor{rect, tri}
	Stage1.Pawns = []*utility.Actor{tri}
}
