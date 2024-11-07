package ebitenhelper

import (
	"errors"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	WindowTitle  string
	ScreenWidth  int
	ScreenHeight int

	currentLevel *Level

	t_PressedKeys  []ebiten.Key
	t_ReleasedKeys []ebiten.Key
	t_PressingKeys []ebiten.Key
}

func NewGame() *Game {
	return &Game{
		WindowTitle:  "Game",
		ScreenWidth:  1280,
		ScreenHeight: 720,
	}
}

func (g *Game) Update() error {
	g.t_PressedKeys = inpututil.AppendJustPressedKeys(g.t_PressedKeys[:0])
	g.t_ReleasedKeys = inpututil.AppendJustReleasedKeys(g.t_ReleasedKeys[:0])
	g.t_PressingKeys = inpututil.AppendPressedKeys(g.t_PressingKeys[:0])

	for _, k := range g.t_PressedKeys {
		g.ReceivePressedKey(k)
	}
	for _, k := range g.t_ReleasedKeys {
		g.ReceiveReleasedKey(k)
	}
	for _, k := range g.t_PressingKeys {
		g.ReceivePressingKey(k)
	}

	g.Tick()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, d := range g.currentLevel.Drawers {
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

	for _, r := range g.currentLevel.KeyReceivers {
		r.ReceivePressedKey(k)
	}
}

func (g *Game) ReceiveReleasedKey(k ebiten.Key) {
	for _, r := range g.currentLevel.KeyReceivers {
		r.ReceiveReleasedKey(k)
	}
}

func (g *Game) ReceivePressingKey(k ebiten.Key) {
	for _, r := range g.currentLevel.KeyReceivers {
		r.ReceivePressingKey(k)
	}
}

func (g *Game) Tick() {
	for _, t := range g.currentLevel.Tickers {
		t.Tick()
	}
}

func (g *Game) LoadLevel(level *Level) error {
	if level == nil {
		return errors.New("loaded level is empty")
	}
	g.currentLevel = level
	return nil
}

func (g *Game) Play(firstlevel *Level) error {
	err := g.LoadLevel(firstlevel)
	if err != nil {
		return err
	}

	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle(g.WindowTitle)

	err = ebiten.RunGame(g)
	if err != nil {
		return err
	}

	return nil
}
