package main

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/actor"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

func NewStage1() *utility.Level {
	// Level
	level := utility.NewLevel("stage1")
	level.IsLooping = true
	actor.AddActorsToLevelFromMapFile(level, "stage1.tmx")

	// Player
	playerImage := utility.GetImageFile("images/ぴぽやキャラチップ32出力素材/現代系/女_スーツ1.png")
	player := actor.NewPawn()
	player.SetLocation(utility.NewVector(600, 300))
	player.Image = playerImage
	level.Add(player)

	// Enemy
	enemyImage := utility.GetImageFile("images/ぴぽやキャラチップ32出力素材/現代系/男_スーツ1.png")
	enemyLocations := []utility.Vector{
		utility.NewVector(700, 300),
		utility.NewVector(800, 300),
		utility.NewVector(900, 300),
		utility.NewVector(1000, 300),
		utility.NewVector(1100, 300),
		utility.NewVector(1200, 300),
		utility.NewVector(500, 500),
		utility.NewVector(600, 500),
		utility.NewVector(700, 500),
		utility.NewVector(800, 500),
	}
	for _, el := range enemyLocations {
		enemy := actor.NewAIPawn()
		enemy.SetLocation(el)
		enemy.Image = enemyImage
		enemy.MaxSpeed = 150
		level.Add(enemy)
	}

	return level
}

func main() {
	utility.SetWindowTitle("Ichimoudajin")
	utility.SetScreenSize(32*40, 32*22)
	utility.PlayGame(NewStage1())
}
