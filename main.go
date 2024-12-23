package main

import (
	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/actor"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

func NewStage1() *utility.Level {
	mapdata, err := assets.GetMapData("stage1.tmx")
	utility.ExitIfError(err)

	image_player, err := assets.GetImage("images/ぴぽやキャラチップ32出力素材/現代系/女_スーツ1.png")
	utility.ExitIfError(err)

	image_blocker, err := assets.GetImage("images/ぴぽやキャラチップ32出力素材/現代系/男_スーツ1.png")
	utility.ExitIfError(err)

	level := utility.NewLevel()
	level.AddRange(mapdata.GetActors())

	player := actor.NewPawn()
	player.SetLocation(utility.NewVector(600, 300))
	player.Image = image_player
	level.Add(player)

	blocker := actor.NewAnimatedActor()
	blocker.SetLocation(utility.NewVector(500, 300))
	blocker.Image = image_blocker
	level.Add(blocker)

	level.IsLooping = true
	return level
}

func main() {
	g := utility.NewGame()
	g.WindowTitle = "Ichimoudajin"
	g.ScreenSize = utility.NewPoint(32*40, 32*22)

	utility.ExitIfError(g.Play(NewStage1()))
}
