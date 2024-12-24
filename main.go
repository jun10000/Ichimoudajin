package main

import (
	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/actor"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

func NewStage1() *utility.Level {
	// Level
	level := utility.NewLevel()
	level.IsLooping = true

	// Map actors
	mapData, err := assets.GetMapData("stage1.tmx")
	utility.ExitIfError(err)
	level.AddRange(mapData.GetActors())

	// Player
	playerImage, err := assets.GetImage("images/ぴぽやキャラチップ32出力素材/現代系/女_スーツ1.png")
	utility.ExitIfError(err)
	player := actor.NewPawn()
	player.SetLocation(utility.NewVector(600, 300))
	player.Image = playerImage
	level.Add(player)

	// Enemy
	enemyImage, err := assets.GetImage("images/ぴぽやキャラチップ32出力素材/現代系/男_スーツ1.png")
	utility.ExitIfError(err)
	enemy := actor.NewAIPawn()
	enemy.SetLocation(utility.NewVector(500, 300))
	enemy.Image = enemyImage
	level.Add(enemy)

	return level
}

func main() {
	g := utility.NewGame()
	g.WindowTitle = "Ichimoudajin"
	g.ScreenSize = utility.NewPoint(32*40, 32*22)

	utility.ExitIfError(g.Play(NewStage1()))
}
