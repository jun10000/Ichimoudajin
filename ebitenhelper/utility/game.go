package utility

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	TickCount    = 60
	TickDuration = 1.0 / float64(TickCount)
)

var (
	MovementMaxReflectionCount = 1
	MovementInvalidDistance    = 3
	AIValidOffset              = 0.5
	AIMaxTaskCount             = 1
	InitialPFResultCap         = 128
	InitialInputReceiverCap    = 1
	InitialAITickerCap         = 32
	InitialTickerCap           = 32
	InitialDrawerCap           = 128
)

var (
	DebugColorRed        = color.RGBA{R: 255, G: 8}
	DebugColorYellow     = color.RGBA{R: 255, G: 255}
	DebugColorGreen      = color.RGBA{G: 255}
	DebugColorBlue       = color.RGBA{G: 128, B: 255}
	DebugColorGray       = color.RGBA{R: 128, G: 128, B: 128}
	DebugInitialDrawsCap = 32

	DebugIsShowLocation     = false
	DebugLocationTextOffset = NewVector(3, -12)

	DebugIsShowTraceDistance = true
	DebugTraceDistanceColors = map[int]color.RGBA{
		0: DebugColorRed,
	}

	DebugIsShowTraceResult = true
	DebugTraceResultLength = 30.0

	DebugIsShowAIPath = false
	DebugAIPathColor  = DebugColorGreen
)

var windowTitle = "Game"

func GetWindowTitle() string {
	return windowTitle
}

func SetWindowTitle(title string) {
	windowTitle = title
	ebiten.SetWindowTitle(title)
}

var screenSize = NewPoint(1280, 720)

func GetScreenSize() Point {
	return screenSize
}

func SetScreenSize(width int, height int) {
	screenSize.X = width
	screenSize.Y = height
	ebiten.SetWindowSize(width, height)
}

var currentLevel *Level

func GetLevel() *Level {
	return currentLevel
}

func SetLevel(level *Level) error {
	if level == nil {
		return errors.New("loaded level is empty")
	}

	currentLevel = level
	if IsDebugMode() {
		return level.LoadOrBuildPFCache()
	} else {
		return level.LoadPFCache()
	}
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
}

func NewGame() *Game {
	return &Game{
		pressedButtons:  map[ebiten.GamepadID][]ebiten.StandardGamepadButton{},
		releasedButtons: map[ebiten.GamepadID][]ebiten.StandardGamepadButton{},
		pressingButtons: map[ebiten.GamepadID][]ebiten.StandardGamepadButton{},
		axisValues:      map[GamepadAxisKey]float64{},
	}
}

func PlayGame(firstlevel *Level) {
	RunDebugServer()

	PanicIfError(SetLevel(firstlevel))
	PanicIfError(ebiten.RunGame(NewGame()))
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

	for _, r := range lv.InputReceivers {
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

	for _, t := range lv.AITickers {
		t.AITick()
	}

	for _, t := range lv.Tickers {
		t.Tick()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	lv := GetLevel()

	for _, d := range lv.Drawers {
		d.Draw(screen)
	}

	if IsDebugMode() {
		for _, d := range lv.DebugDraws {
			d(screen)
		}
		lv.ClearDebugDraw()
	}
}

func (g *Game) Layout(width int, height int) (int, int) {
	s := GetScreenSize()
	return s.X, s.Y
}
