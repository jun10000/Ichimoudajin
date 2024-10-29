package utility

import (
	_ "image/png"
	"log"
	"os"

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

func NewGame() *Game {
	return &Game{
		WindowTitle:  "Game",
		ScreenWidth:  1280,
		ScreenHeight: 720,
	}
}

func (g *Game) Update() error {
	g.Temp_PressedKeys = inpututil.AppendJustPressedKeys(g.Temp_PressedKeys[:0])
	g.Temp_ReleasedKeys = inpututil.AppendJustReleasedKeys(g.Temp_ReleasedKeys[:0])
	g.Temp_PressingKeys = inpututil.AppendPressedKeys(g.Temp_PressingKeys[:0])

	for _, k := range g.Temp_PressedKeys {
		g.ReceivePressedKey(k)
	}
	for _, k := range g.Temp_ReleasedKeys {
		g.ReceiveReleasedKey(k)
	}
	for _, k := range g.Temp_PressingKeys {
		g.ReceivePressingKey(k)
	}

	g.Tick()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, d := range g.CurrentLevel.Drawers {
		d.Draw(screen)
	}
}

func (g *Game) Layout(width int, height int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func (g *Game) ReceivePressedKey(k ebiten.Key) {
	switch k {
	case ebiten.KeyEscape:
		os.Exit(0)
	}

	for _, r := range g.CurrentLevel.KeyReceivers {
		r.ReceivePressedKey(k)
	}
}

func (g *Game) ReceiveReleasedKey(k ebiten.Key) {
	for _, r := range g.CurrentLevel.KeyReceivers {
		r.ReceiveReleasedKey(k)
	}
}

func (g *Game) ReceivePressingKey(k ebiten.Key) {
	for _, r := range g.CurrentLevel.KeyReceivers {
		r.ReceivePressingKey(k)
	}
}

func (g *Game) Tick() {
	for _, t := range g.CurrentLevel.Tickers {
		t.Tick()
	}
}

func (g *Game) Play(firstlevel *Level) {
	if firstlevel == nil {
		log.Fatal("Specified first level is invalid")
	}
	g.CurrentLevel = firstlevel

	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle(g.WindowTitle)

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
