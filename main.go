package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/tilemap"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type GameInstance struct{ *utility.GameInstanceBase }

func (g *GameInstance) ReceiveKeyInput(key ebiten.Key, state utility.PressState) {
	g.GameInstanceBase.ReceiveKeyInput(key, state)

	switch state {
	case utility.PressStatePressed:
		switch key {
		case ebiten.KeyF11:
			ebiten.SetFullscreen(!ebiten.IsFullscreen())
		case ebiten.KeyEscape:
			utility.Exit(0)
		}
	}
}

func main() {
	utility.SetWindowTitle("Ichimoudajin")
	utility.SetScreenSize(32*40, 32*22)
	utility.PlayGame(&GameInstance{}, tilemap.NewLevelByTiledMap("stage1"))
}
