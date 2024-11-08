package main

import (
	"log"

	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/actor"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

func NewStage1() *utility.Level {
	mapdata, err := assets.GetMapData("stage1.tmx")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mapdata.MapSize)

	image_player, err := assets.GetImage("images/ぴぽやキャラチップ32出力素材/現代系/女_スーツ1.png")
	if err != nil {
		log.Fatal(err)
	}

	level := utility.NewLevel()

	player := actor.NewPawn()
	player.Location = utility.NewVector(600, 300)
	player.Scale = utility.NewVector(2, 2)
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
