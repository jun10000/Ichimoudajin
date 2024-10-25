package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/gamemap"
	"github.com/jun10000/Ichimoudajin/utility"
)

type Game struct {
	ScreenWidth  int
	ScreenHeight int
	WindowTitle  string
	FirstLevel   *utility.Level
	CurrentLevel *utility.Level
}

func PlayGame(screen_w int, screen_h int, title string, firstlevel *utility.Level) {
	g := &Game{screen_w, screen_h, title, firstlevel, firstlevel}
	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle(g.WindowTitle)

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	for _, a := range g.CurrentLevel.Actors {
		a.Draw(s)
	}
}

func (g *Game) Layout(w int, h int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func main() {
	PlayGame(1280, 720, "Ichimoudajin", gamemap.Stage1)
}
