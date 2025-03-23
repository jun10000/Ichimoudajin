package main

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/tilemap"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

func main() {
	utility.SetWindowTitle("Ichimoudajin")
	utility.SetScreenSize(32*40, 32*22)
	utility.PlayGame(tilemap.NewLevelByTiledMap("stage1"))
}
