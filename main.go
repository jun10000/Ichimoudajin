package main

import (
	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper"
)

func NewStage1() *ebitenhelper.Level {
	level := ebitenhelper.NewLevel()

	background := ebitenhelper.NewLandscape()
	background.Image.Source = assets.GetImage("terracotta-tiles-941741_640.jpg")
	background.Image.TileScale = ebitenhelper.NewVector(0.25, 0.25)
	level.Add(background)

	player := ebitenhelper.NewPawn()
	player.Location = ebitenhelper.NewVector(600, 300)
	player.Scale = ebitenhelper.NewVector(2, 2)
	player.Animation.Source = assets.GetImage("ぴぽやキャラチップ32出力素材/現代系/女_スーツ1.png")
	level.Add(player)

	return level
}

func main() {
	g := ebitenhelper.NewGame()
	g.WindowTitle = "Ichimoudajin"
	g.Play(NewStage1())
}
