package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/tilemap"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/widget"
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

func NewStage1() *utility.Level {
	lv := tilemap.NewLevelByTiledMap("stage1")

	mainWidget, err := widget.NewWidgetByFile("mainwidget")
	utility.PanicIfError(err)

	lv.Add(mainWidget)
	return lv
}

func main() {
	utility.SetAssetFS(assets.Assets)
	utility.SetWindowTitle("Ichimoudajin 0.0.2")
	utility.SetScreenSize(32*40, 32*22)
	utility.PlayGame(&GameInstance{}, NewStage1())
}
