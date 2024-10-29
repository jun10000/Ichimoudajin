package gamemap

import (
	"github.com/jun10000/Ichimoudajin/utility"
)

func NewStage1() *utility.Level {
	level := utility.NewLevel()

	acter_rect := utility.NewActor("rectangle100x100.png")
	acter_rect.Location = utility.NewVector(100, 200)
	level.Add(acter_rect)

	pawn_tri := utility.NewPawn("triangle100x100.png")
	pawn_tri.Location = utility.NewVector(300, 400)
	level.Add(pawn_tri)

	return level
}
