package utility

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jun10000/Ichimoudajin/assets"
)

const (
	TickCount    = 60
	TickDuration = 1.0 / float64(TickCount)

	ZOrderDefault = 0
	ZOrderEffect  = 1
	ZOrderMax     = ZOrderEffect
)

var (
	MovementMaxReflectionCount = 1
	AIValidOffset              = 0.5
	AIMaxTaskCount             = 1
	GamepadDeadZone            = 0.2

	ColorRed        = RGB{0xff, 0x00, 0x00}
	ColorOrange     = RGB{0xff, 0x80, 0x00}
	ColorYellow     = RGB{0xff, 0xff, 0x00}
	ColorLightGreen = RGB{0x80, 0xff, 0x00}
	ColorGreen      = RGB{0x00, 0xff, 0x00}
	ColorLightBlue  = RGB{0x00, 0x80, 0xff}
	ColorBlue       = RGB{0x00, 0x00, 0xff}
	ColorPurple     = RGB{0x80, 0x00, 0xff}
	ColorWhite      = RGB{0xff, 0xff, 0xff}
	ColorGray       = RGB{0x80, 0x80, 0x80}
	ColorBlack      = RGB{0x00, 0x00, 0x00}

	InitialStaticColliderCap  = 128
	InitialMovableColliderCap = 32
	InitialInputReceiverCap   = 1
	InitialPlayerCap          = 1
	InitialAITickerCap        = 32
	InitialTickerCap          = 32
	InitialDrawerCap          = 128
	InitialTrashCap           = 32
	InitialPFResultCap        = 128
)

var (
	DebugInitialDrawsCap = 32

	DebugIsShowLocation     = false
	DebugLocationTextOffset = NewVector(3, -12)

	DebugIsShowTraceDistance = false
	DebugTraceDistanceColors = map[int]color.Color{
		0: ColorRed,
	}

	DebugIsShowTraceResult = false
	DebugTraceResultLength = 30.0

	DebugIsShowAIPath = false
	DebugAIPathColor  = ColorGreen
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

var ebitenImages = NewSmap[string, *ebiten.Image]()

func GetImageFromFile(filename string) (*ebiten.Image, error) {
	if img, ok := ebitenImages.Load(filename); ok {
		return img, nil
	}

	img, _, err := ebitenutil.NewImageFromFileSystem(assets.Assets, filename)
	if err != nil {
		return nil, err
	}

	ebitenImages.Store(filename, img)
	return img, nil
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
	pressedKeys            []ebiten.Key
	releasedKeys           []ebiten.Key
	pressingKeys           []ebiten.Key
	pressedMouseButtons    []ebiten.MouseButton
	releasedMouseButtons   []ebiten.MouseButton
	pressingMouseButtons   Set[ebiten.MouseButton]
	gamepadIDs             []ebiten.GamepadID
	pressedGamepadButtons  map[ebiten.GamepadID][]ebiten.StandardGamepadButton
	releasedGamepadButtons map[ebiten.GamepadID][]ebiten.StandardGamepadButton
	pressingGamepadButtons map[ebiten.GamepadID][]ebiten.StandardGamepadButton
	gamepadAxisValues      map[GamepadAxisKey]float64
}

func NewGame() *Game {
	return &Game{
		pressedMouseButtons:    make([]ebiten.MouseButton, 0, ebiten.MouseButtonMax+1),
		releasedMouseButtons:   make([]ebiten.MouseButton, 0, ebiten.MouseButtonMax+1),
		pressingMouseButtons:   make(Set[ebiten.MouseButton]),
		pressedGamepadButtons:  make(map[ebiten.GamepadID][]ebiten.StandardGamepadButton),
		releasedGamepadButtons: make(map[ebiten.GamepadID][]ebiten.StandardGamepadButton),
		pressingGamepadButtons: make(map[ebiten.GamepadID][]ebiten.StandardGamepadButton),
		gamepadAxisValues:      make(map[GamepadAxisKey]float64),
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
	for _, b := range g.pressedMouseButtons {
		g.pressingMouseButtons.Add(b)
	}
	g.pressedMouseButtons = g.pressedMouseButtons[:0]
	g.releasedMouseButtons = g.releasedMouseButtons[:0]
	for b := ebiten.MouseButton0; b <= ebiten.MouseButtonMax; b++ {
		if inpututil.IsMouseButtonJustPressed(b) {
			g.pressedMouseButtons = append(g.pressedMouseButtons, b)
		}
		if inpututil.IsMouseButtonJustReleased(b) {
			g.releasedMouseButtons = append(g.releasedMouseButtons, b)
			g.pressingMouseButtons.Remove(b)
		}
	}
	g.gamepadIDs = inpututil.AppendJustConnectedGamepadIDs(g.gamepadIDs)
	for _, id := range g.gamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) ||
			!ebiten.IsStandardGamepadLayoutAvailable(id) {
			g.gamepadIDs = RemoveSliceItem(g.gamepadIDs, id)
		}
	}
	for _, id := range g.gamepadIDs {
		g.pressedGamepadButtons[id] = inpututil.AppendJustPressedStandardGamepadButtons(id, g.pressedGamepadButtons[id][:0])
		g.releasedGamepadButtons[id] = inpututil.AppendJustReleasedStandardGamepadButtons(id, g.releasedGamepadButtons[id][:0])
		g.pressingGamepadButtons[id] = inpututil.AppendPressedStandardGamepadButtons(id, g.pressingGamepadButtons[id][:0])
		for a := ebiten.StandardGamepadAxis(0); a <= ebiten.StandardGamepadAxisMax; a++ {
			k := NewGamepadAxisKey(id, a)
			v := ebiten.StandardGamepadAxisValue(id, a)
			g.gamepadAxisValues[k] = v
		}
	}

	lv := GetLevel()
	cp := GetCursorPosition()

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

		for _, b := range g.pressedMouseButtons {
			r.ReceiveMouseButtonInput(b, PressStatePressed, cp)
		}
		for _, b := range g.releasedMouseButtons {
			r.ReceiveMouseButtonInput(b, PressStateReleased, cp)
		}
		for b := range g.pressingMouseButtons {
			r.ReceiveMouseButtonInput(b, PressStatePressing, cp)
		}

		for _, id := range g.gamepadIDs {
			for _, b := range g.pressedGamepadButtons[id] {
				r.ReceiveGamepadButtonInput(id, b, PressStatePressed)
			}
			for _, b := range g.releasedGamepadButtons[id] {
				r.ReceiveGamepadButtonInput(id, b, PressStateReleased)
			}
			for _, b := range g.pressingGamepadButtons[id] {
				r.ReceiveGamepadButtonInput(id, b, PressStatePressing)
			}
			for a := ebiten.StandardGamepadAxis(0); a <= ebiten.StandardGamepadAxisMax; a++ {
				k := NewGamepadAxisKey(id, a)
				r.ReceiveGamepadAxisInput(id, a, g.gamepadAxisValues[k])
			}
		}
	}

	for _, t := range lv.AITickers {
		t.AITick()
	}

	for _, t := range lv.Tickers {
		t.Tick()
	}

	lv.EmptyTrashes()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	lv := GetLevel()

	for _, ds := range lv.Drawers {
		for _, d := range ds {
			d.Draw(screen)
		}
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
