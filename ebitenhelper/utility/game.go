package utility

import (
	"errors"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var currentGameInstance *Game
var currentLevel *Level

func GetGameInstance() *Game {
	return currentGameInstance
}

func GetLevel() *Level {
	return currentLevel
}

func SetLevel(level *Level) error {
	if level == nil {
		return errors.New("loaded level is empty")
	}
	currentLevel = level
	return nil
}

type Game struct {
	WindowTitle  string
	ScreenWidth  int
	ScreenHeight int

	pressedKeys  []ebiten.Key
	releasedKeys []ebiten.Key
	pressingKeys []ebiten.Key
	drawEvents   []func(screen *ebiten.Image)
}

func NewGame() *Game {
	return &Game{
		WindowTitle:  "Game",
		ScreenWidth:  1280,
		ScreenHeight: 720,
	}
}

func (g *Game) AddDrawEvent(event func(*ebiten.Image)) {
	g.drawEvents = append(g.drawEvents, event)
}

func (g *Game) Update() error {
	g.pressedKeys = inpututil.AppendJustPressedKeys(g.pressedKeys[:0])
	g.releasedKeys = inpututil.AppendJustReleasedKeys(g.releasedKeys[:0])
	g.pressingKeys = inpututil.AppendPressedKeys(g.pressingKeys[:0])

	for _, k := range g.pressedKeys {
		g.ReceivePressedKey(k)
	}
	for _, k := range g.releasedKeys {
		g.ReceiveReleasedKey(k)
	}
	for _, k := range g.pressingKeys {
		g.ReceivePressingKey(k)
	}

	g.Tick()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, d := range GetLevel().Drawers {
		d.Draw(screen)
	}

	for _, d := range g.drawEvents {
		d(screen)
	}
	g.drawEvents = g.drawEvents[:0]
}

func (g *Game) Layout(width int, height int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}

func (g *Game) ReceivePressedKey(k ebiten.Key) {
	switch k {
	case ebiten.KeyEscape:
		os.Exit(0)
	}

	for _, r := range GetLevel().KeyReceivers {
		r.ReceivePressedKey(k)
	}
}

func (g *Game) ReceiveReleasedKey(k ebiten.Key) {
	for _, r := range GetLevel().KeyReceivers {
		r.ReceiveReleasedKey(k)
	}
}

func (g *Game) ReceivePressingKey(k ebiten.Key) {
	for _, r := range GetLevel().KeyReceivers {
		r.ReceivePressingKey(k)
	}
}

func (g *Game) Tick() {
	for _, t := range GetLevel().Tickers {
		t.Tick()
	}
}

func (g *Game) Play(firstlevel *Level) error {
	err := SetLevel(firstlevel)
	if err != nil {
		return err
	}

	currentGameInstance = g
	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle(g.WindowTitle)

	err = ebiten.RunGame(g)
	if err != nil {
		return err
	}

	return nil
}
