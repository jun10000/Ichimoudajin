package utility

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	WindowTitle  string
	ScreenWidth  int
	ScreenHeight int
	CurrentLevel *Level

	Temp_PressedKeys  []ebiten.Key
	Temp_ReleasedKeys []ebiten.Key
	Temp_PressingKeys []ebiten.Key
}

func NewGame(firstlevel *Level) *Game {
	if firstlevel == nil {
		log.Fatal("Specified first level is invalid")
	}

	return &Game{
		WindowTitle:  "Game",
		ScreenWidth:  1280,
		ScreenHeight: 720,
		CurrentLevel: firstlevel,
	}
}

func (g *Game) Update() error {
	g.Temp_PressedKeys = inpututil.AppendJustPressedKeys(g.Temp_PressedKeys[:0])
	g.Temp_ReleasedKeys = inpututil.AppendJustReleasedKeys(g.Temp_ReleasedKeys[:0])
	g.Temp_PressingKeys = inpututil.AppendPressedKeys(g.Temp_PressingKeys[:0])

	for _, p := range g.CurrentLevel.Pawns {
		for _, k := range g.Temp_PressedKeys {
			p.Event_KeyPressed(k)
		}
		for _, k := range g.Temp_ReleasedKeys {
			p.Event_KeyReleased(k)
		}
		for _, k := range g.Temp_PressingKeys {
			p.Event_KeyPressing(k)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, a := range g.CurrentLevel.Actors {
		a.Draw(screen)
	}
}

func (g *Game) Layout(width int, height int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func (g *Game) Play() {
	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle(g.WindowTitle)

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
