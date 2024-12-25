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

type GamepadAxisKey struct {
	ID   ebiten.GamepadID
	Axis ebiten.StandardGamepadAxis
}

func NewGamepadAxisKey(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis) GamepadAxisKey {
	return GamepadAxisKey{
		ID:   id,
		Axis: axis,
	}
}

type Game struct {
	WindowTitle string
	ScreenSize  Point

	pressedKeys     []ebiten.Key
	releasedKeys    []ebiten.Key
	pressingKeys    []ebiten.Key
	gamepadIDs      []ebiten.GamepadID
	pressedButtons  map[ebiten.GamepadID][]ebiten.StandardGamepadButton
	releasedButtons map[ebiten.GamepadID][]ebiten.StandardGamepadButton
	pressingButtons map[ebiten.GamepadID][]ebiten.StandardGamepadButton
	axisValues      map[GamepadAxisKey]float64
	drawEvents      []func(screen *ebiten.Image)
}

func NewGame() *Game {
	return &Game{
		WindowTitle:     "Game",
		ScreenSize:      NewPoint(1280, 720),
		pressedButtons:  map[ebiten.GamepadID][]ebiten.StandardGamepadButton{},
		releasedButtons: map[ebiten.GamepadID][]ebiten.StandardGamepadButton{},
		pressingButtons: map[ebiten.GamepadID][]ebiten.StandardGamepadButton{},
		axisValues:      map[GamepadAxisKey]float64{},
	}
}

func (g *Game) AddDrawEvent(event func(*ebiten.Image)) {
	g.drawEvents = append(g.drawEvents, event)
}

func (g *Game) Update() error {
	g.pressedKeys = inpututil.AppendJustPressedKeys(g.pressedKeys[:0])
	g.releasedKeys = inpututil.AppendJustReleasedKeys(g.releasedKeys[:0])
	g.pressingKeys = inpututil.AppendPressedKeys(g.pressingKeys[:0])
	g.gamepadIDs = inpututil.AppendJustConnectedGamepadIDs(g.gamepadIDs)
	for _, id := range g.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) ||
			!ebiten.IsStandardGamepadLayoutAvailable(id) {
			g.gamepadIDs = RemoveSliceItem(g.gamepadIDs, id)
		}
	}
	for _, id := range g.gamepadIDs {
		g.pressedButtons[id] = inpututil.AppendJustPressedStandardGamepadButtons(id, g.pressedButtons[id][:0])
		g.releasedButtons[id] = inpututil.AppendJustReleasedStandardGamepadButtons(id, g.releasedButtons[id][:0])
		g.pressingButtons[id] = inpututil.AppendPressedStandardGamepadButtons(id, g.pressingButtons[id][:0])
		for a := ebiten.StandardGamepadAxis(0); a <= ebiten.StandardGamepadAxisMax; a++ {
			k := NewGamepadAxisKey(id, a)
			v := ebiten.StandardGamepadAxisValue(id, a)
			g.axisValues[k] = v
		}
	}

	for _, r := range GetLevel().InputReceivers {
		for _, k := range g.pressedKeys {
			r.ReceiveKeyInput(k, PressStatePressed)
		}
		for _, k := range g.releasedKeys {
			r.ReceiveKeyInput(k, PressStateReleased)
		}
		for _, k := range g.pressingKeys {
			r.ReceiveKeyInput(k, PressStatePressing)
		}
		for _, id := range g.gamepadIDs {
			for _, b := range g.pressedButtons[id] {
				r.ReceiveButtonInput(id, b, PressStatePressed)
			}
			for _, b := range g.releasedButtons[id] {
				r.ReceiveButtonInput(id, b, PressStateReleased)
			}
			for _, b := range g.pressingButtons[id] {
				r.ReceiveButtonInput(id, b, PressStatePressing)
			}
			for a := ebiten.StandardGamepadAxis(0); a <= ebiten.StandardGamepadAxisMax; a++ {
				k := NewGamepadAxisKey(id, a)
				r.ReceiveAxisInput(id, a, g.axisValues[k])
			}
		}
	}

	for _, t := range GetLevel().AITickers {
		t.AITick()
	}

	for _, t := range GetLevel().Tickers {
		t.Tick()
	}

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
