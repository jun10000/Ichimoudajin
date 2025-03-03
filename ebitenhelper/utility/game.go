package utility

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	currentGameInstance *Game
	currentLevel        *Level

	WindowTitle = "Game"
	ScreenSize  = NewPoint(1280, 720)
)

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
	if AIIsUsePFCacheFile {
		if IsDebugMode {
			return level.LoadOrBuildPFCache()
		} else {
			return level.LoadPFCache()
		}
	}

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
		pressedButtons:  map[ebiten.GamepadID][]ebiten.StandardGamepadButton{},
		releasedButtons: map[ebiten.GamepadID][]ebiten.StandardGamepadButton{},
		pressingButtons: map[ebiten.GamepadID][]ebiten.StandardGamepadButton{},
		axisValues:      map[GamepadAxisKey]float64{},
		drawEvents:      make([]func(screen *ebiten.Image), 0, InitialDrawEventCap),
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

	lv := GetLevel()

	for r := range lv.InputReceivers {
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

	for t := range lv.AITickers {
		t.AITick()
	}

	for t := range lv.Tickers {
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
	return ScreenSize.X, ScreenSize.Y
}

func (g *Game) Play(firstlevel *Level) error {
	RunDebugServer()

	ebiten.SetWindowSize(ScreenSize.X, ScreenSize.Y)
	ebiten.SetWindowTitle(WindowTitle)
	currentGameInstance = g
	err := SetLevel(firstlevel)
	if err != nil {
		return err
	}

	err = ebiten.RunGame(g)
	if err != nil {
		return err
	}

	return nil
}
