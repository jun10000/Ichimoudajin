package main

import (
	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper"
)

func NewStage1() *ebitenhelper.Level {
	level := ebitenhelper.NewLevel()

	env_background := ebitenhelper.NewLandscape()
	env_background.Scale = ebitenhelper.NewVector(0.25, 0.25)
	env_background.Image = assets.GetImage("terracotta-tiles-941741_640.jpg")
	level.Add(env_background)

	player := ebitenhelper.NewPawn()
	player.Location = ebitenhelper.NewVector(600, 300)
	player.Scale = ebitenhelper.NewVector(2, 2)
	player.Animation.Image = assets.GetImage("ぴぽやキャラチップ32出力素材/現代系/女_スーツ1.png")
	level.Add(player)

	return level
}

func main() {
	g := ebitenhelper.NewGame()
	g.WindowTitle = "Ichimoudajin"
	g.Play(NewStage1())
}
