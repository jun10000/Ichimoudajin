package main

import (
	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/utility"
)

func NewStage1() *utility.Level {
	level := utility.NewLevel()

	env_background := utility.NewLandscape()
	env_background.Image = assets.GetImage("terracotta-tiles-941741_640.jpg")
	env_background.TileScale = utility.NewVector(0.25, 0.25)
	level.Add(env_background)

	player := utility.NewPawn()
	player.Image = assets.GetImage("ぴぽやキャラチップ32出力素材/現代系/女_スーツ1.png")
	player.Location = utility.NewVector(600, 300)
	level.Add(player)

	return level
}

func main() {
	g := utility.NewGame()
	g.WindowTitle = "Ichimoudajin"
	g.Play(NewStage1())
}
