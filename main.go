package main

import (
	"log"

	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper"
)

func NewStage1() *ebitenhelper.Level {
	mapdata, err := assets.GetMapData("stage1.tmx")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mapdata.MapSize)

	image_player, err := assets.GetImage("images/ぴぽやキャラチップ32出力素材/現代系/女_スーツ1.png")
	if err != nil {
		log.Fatal(err)
	}

	level := ebitenhelper.NewLevel()

	player := ebitenhelper.NewPawn()
	player.Location = ebitenhelper.NewVector(600, 300)
	player.Scale = ebitenhelper.NewVector(2, 2)
	player.Animation.Source = image_player
	level.Add(player)

	return level
}

func main() {
	g := ebitenhelper.NewGame()
	g.WindowTitle = "Ichimoudajin"
	err := g.Play(NewStage1())
	if err != nil {
		log.Fatal(err)
	}
}
