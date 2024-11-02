package main

import (
	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/utility"
)

func NewStage1() *utility.Level {
	level := utility.NewLevel()

	env_background := utility.NewLandscape()
	env_background.Scale = utility.NewVector(0.25, 0.25)
	env_background.Image = assets.GetImage("terracotta-tiles-941741_640.jpg")
	level.Add(env_background)

	player := utility.NewPawn()
	player.Location = utility.NewVector(600, 300)
	player.Scale = utility.NewVector(2, 2)
	player.Animation.Image = assets.GetImage("ぴぽやキャラチップ32出力素材/現代系/女_スーツ1.png")
	level.Add(player)

	return level
}

func main() {
	g := utility.NewGame()
	g.WindowTitle = "Ichimoudajin"
	g.Play(NewStage1())
}
