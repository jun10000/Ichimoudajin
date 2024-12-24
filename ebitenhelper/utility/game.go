package utility

import (
	"errors"

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
	WindowTitle string
	ScreenSize  Point

	pressedKeys  []ebiten.Key
	releasedKeys []ebiten.Key
	pressingKeys []ebiten.Key
	gamepadIDs   []ebiten.GamepadID
	drawEvents   []func(screen *ebiten.Image)
}

func NewGame() *Game {
	return &Game{
		WindowTitle: "Game",
		ScreenSize:  NewPoint(1280, 720),
	}
}

func (g *Game) AddDrawEvent(event func(*ebiten.Image)) {
	g.drawEvents = append(g.drawEvents, event)
}

func (g *Game) GetGamepadIDs() []ebiten.GamepadID {
	g.gamepadIDs = inpututil.AppendJustConnectedGamepadIDs(g.gamepadIDs)
	for _, id := range g.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) ||
			!ebiten.IsStandardGamepadLayoutAvailable(id) {
			g.gamepadIDs = RemoveSliceItem(g.gamepadIDs, id)
		}
	}

	return g.gamepadIDs
}

func (g *Game) Update() error {
	g.pressedKeys = inpututil.AppendJustPressedKeys(g.pressedKeys[:0])
	for _, k := range g.pressedKeys {
		g.ReceiveKeyInput(k, PressStatePressed)
	}

	g.releasedKeys = inpututil.AppendJustReleasedKeys(g.releasedKeys[:0])
	for _, k := range g.releasedKeys {
		g.ReceiveKeyInput(k, PressStateReleased)
	}

	g.pressingKeys = inpututil.AppendPressedKeys(g.pressingKeys[:0])
	for _, k := range g.pressingKeys {
		g.ReceiveKeyInput(k, PressStatePressing)
	}

	for _, id := range g.GetGamepadIDs() {
		for b := ebiten.StandardGamepadButton(0); b <= ebiten.StandardGamepadButtonMax; b++ {
			if inpututil.IsStandardGamepadButtonJustPressed(id, b) {
				g.ReceiveButtonInput(id, b, PressStatePressed)
			}
			if inpututil.IsStandardGamepadButtonJustReleased(id, b) {
				g.ReceiveButtonInput(id, b, PressStateReleased)
			}
		}
		for a := ebiten.StandardGamepadAxis(0); a <= ebiten.StandardGamepadAxisMax; a++ {
			g.ReceiveAxisInput(id, a, ebiten.StandardGamepadAxisValue(id, a))
		}
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
	return g.ScreenSize.X, g.ScreenSize.Y
}

func (g *Game) ReceiveKeyInput(key ebiten.Key, state PressState) {
	for _, r := range GetLevel().InputReceivers {
		r.ReceiveKeyInput(key, state)
	}
}

func (g *Game) ReceiveButtonInput(id ebiten.GamepadID, button ebiten.StandardGamepadButton, state PressState) {
	for _, r := range GetLevel().InputReceivers {
		r.ReceiveButtonInput(id, button, state)
	}
}

func (g *Game) ReceiveAxisInput(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis, value float64) {
	for _, r := range GetLevel().InputReceivers {
		r.ReceiveAxisInput(id, axis, value)
	}
}

func (g *Game) Tick() {
	for _, t := range GetLevel().Tickers {
		t.Tick()
	}
}

func (g *Game) Play(firstlevel *Level) error {
	RunDebugServer()

	err := SetLevel(firstlevel)
	if err != nil {
		return err
	}

	currentGameInstance = g
	ebiten.SetWindowSize(g.ScreenSize.X, g.ScreenSize.Y)
	ebiten.SetWindowTitle(g.WindowTitle)

	err = ebiten.RunGame(g)
	if err != nil {
		return err
	}

	return nil
}
