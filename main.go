package main

import (
	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/utility"
)

func NewStage1() *utility.Level {
	level := utility.NewLevel()

	env_background := utility.NewActor()
	env_background.Image = assets.GetImage("terracotta-tiles-941741_640.jpg")
	level.Add(env_background)

	pawn_tri := utility.NewPawn()
	pawn_tri.Location = utility.NewVector(300, 400)
	pawn_tri.Image = assets.GetImage("triangle100x100.png")
	level.Add(pawn_tri)

	return level
}

func main() {
	g := utility.NewGame()
	g.WindowTitle = "Ichimoudajin"
	g.Play(NewStage1())
}
