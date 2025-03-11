package main

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/actor"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

func NewStage1() *utility.Level {
	level := utility.NewLevel("stage1", true)
	actor.AddTileMapActorsToLevel(level, "stage1.tmx")
	return level
}

func main() {
	utility.SetWindowTitle("Ichimoudajin")
	utility.SetScreenSize(32*40, 32*22)
	utility.PlayGame(NewStage1())
}
