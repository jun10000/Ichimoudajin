package main

import (
	"log"

	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper"
)

func NewStage1() *ebitenhelper.Level {
	level := ebitenhelper.NewLevel()

	mapdata := assets.GetMapData("stage1.tmx")
	log.Println(mapdata.Version)

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
