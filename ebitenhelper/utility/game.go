package utility

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/assets"
)

const (
	TickCount    = 60
	TickDuration = 1.0 / float64(TickCount)

	ZOrderDefault = 0
	ZOrderEffect  = 1
	ZOrderWidget  = 2
	ZOrderMax     = ZOrderWidget
)

var (
	MovementMaxReflectionCount = 1
	AIValidOffset              = 0.5
	AIMaxTaskCount             = 1
	GamepadDeadZone            = 0.2

	ColorRed         = RGB{0xff, 0x00, 0x00}
	ColorOrange      = RGB{0xff, 0x80, 0x00}
	ColorYellow      = RGB{0xff, 0xff, 0x00}
	ColorLightGreen  = RGB{0x80, 0xff, 0x00}
	ColorGreen       = RGB{0x00, 0xff, 0x00}
	ColorLightBlue   = RGB{0x00, 0x80, 0xff}
	ColorBlue        = RGB{0x00, 0x00, 0xff}
	ColorPurple      = RGB{0x80, 0x00, 0xff}
	ColorWhite       = RGB{0xff, 0xff, 0xff}
	ColorGray        = RGB{0x80, 0x80, 0x80}
	ColorBlack       = RGB{0x00, 0x00, 0x00}
	ColorTransparent = ColorBlack.ToNRGBA(0x00)

	InitialActorCap                = 128
	InitialStaticColliderCap       = 128
	InitialMovableColliderCap      = 32
	InitialInputReceivableActorCap = 1
	InitialPlayerCap               = 1
	InitialBeginPlayerCap          = 1
	InitialEndPlayerCap            = 1
	InitialAITickerCap             = 32
	InitialTickerCap               = 32
	InitialDrawerCap               = 128
	InitialTrashCap                = 32
	InitialPFResultCap             = 128
)

var (
	DebugInitialDrawsCap = 32

	DebugIsShowLocation     = false
	DebugLocationTextOffset = NewVector(3, -12)

	DebugIsShowTraceDistance = false
	DebugTraceDistanceColors = map[int]color.Color{
		0: ColorRed,
	}

	DebugIsShowTraceResult               = false
	DebugTraceResultLength               = 30.0
	DebugTraceResultOffsetColor          = ColorLightGreen
	DebugTraceResultRemainingOffsetColor = ColorWhite
	DebugTraceResultHitNormalColor       = ColorRed

	DebugIsShowAIPath = false
	DebugAIPathColor  = ColorGreen.ToNRGBA(0x30)
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

var gameInstance GameInstancer

func GetGameInstance[T GameInstancer]() T {
	return gameInstance.(T)
}

var currentLevel *Level

func GetLevel() *Level {
	return currentLevel
}

func SetLevel(level *Level) error {
	// Check level
	if level == nil {
		return errors.New("loaded level is empty")
	}

	// Execute all EndPlay
	if currentLevel != nil {
		for _, a := range currentLevel.EndPlayers {
			a.EndPlay()
		}
	}

	// Change level
	currentLevel = level

	// Create PF cache
	if IsDebugMode() {
		err := level.LoadOrBuildPFCache()
		if err != nil {
			return err
		}
	} else {
		err := level.LoadPFCache()
		if err != nil {
			return err
		}
	}

	// Execute all BeginPlay
	for _, a := range level.BeginPlayers {
		a.BeginPlay()
	}

	return nil
}

var tickIndex = 0

func GetTickIndex() int {
	return tickIndex
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

func GetImageFromFileP(filename string) *ebiten.Image {
	img, err := GetImageFromFile(filename)
	PanicIfError(err)
	return img
}

var ebitenFonts = NewSmap[string, *text.GoTextFaceSource]()

func GetFontFromFile(filename string) (*text.GoTextFaceSource, error) {
	if ff, ok := ebitenFonts.Load(filename); ok {
		return ff, nil
	}

	f, err := assets.Assets.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	ff, err := text.NewGoTextFaceSource(f)
	if err != nil {
		return nil, err
	}

	ebitenFonts.Store(filename, ff)
	return ff, nil
}

func GetFontFromFileP(filename string) *text.GoTextFaceSource {
	ff, err := GetFontFromFile(filename)
	PanicIfError(err)
	return ff
}

type GameInstanceBase struct{}

func (g *GameInstanceBase) ReceiveKeyInput(key ebiten.Key, state PressState) {
}
func (g *GameInstanceBase) ReceiveMouseButtonInput(button ebiten.MouseButton, state PressState, pos Point) {
}
func (g *GameInstanceBase) ReceiveGamepadButtonInput(id ebiten.GamepadID, button ebiten.StandardGamepadButton, state PressState) {
}
func (g *GameInstanceBase) ReceiveGamepadAxisInput(id ebiten.GamepadID, axis ebiten.StandardGamepadAxis, value float64) {
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

func PlayGame(instance GameInstancer, firstlevel *Level) {
	RunDebugServer()

	gameInstance = instance
	PanicIfError(SetLevel(firstlevel))
	PanicIfError(ebiten.RunGame(NewGame()))
}

func (g *Game) callInputReceiverEvent(receiver InputReceiver, mousePosition Point) {
	for _, k := range g.pressedKeys {
		receiver.ReceiveKeyInput(k, PressStatePressed)
	}
	for _, k := range g.releasedKeys {
		receiver.ReceiveKeyInput(k, PressStateReleased)
	}
	for _, k := range g.pressingKeys {
		receiver.ReceiveKeyInput(k, PressStatePressing)
	}

	for _, b := range g.pressedMouseButtons {
		receiver.ReceiveMouseButtonInput(b, PressStatePressed, mousePosition)
	}
	for _, b := range g.releasedMouseButtons {
		receiver.ReceiveMouseButtonInput(b, PressStateReleased, mousePosition)
	}
	for b := range g.pressingMouseButtons {
		receiver.ReceiveMouseButtonInput(b, PressStatePressing, mousePosition)
	}

	for _, id := range g.gamepadIDs {
		for _, b := range g.pressedGamepadButtons[id] {
			receiver.ReceiveGamepadButtonInput(id, b, PressStatePressed)
		}
		for _, b := range g.releasedGamepadButtons[id] {
			receiver.ReceiveGamepadButtonInput(id, b, PressStateReleased)
		}
		for _, b := range g.pressingGamepadButtons[id] {
			receiver.ReceiveGamepadButtonInput(id, b, PressStatePressing)
		}
		for a := ebiten.StandardGamepadAxis(0); a <= ebiten.StandardGamepadAxisMax; a++ {
			k := NewGamepadAxisKey(id, a)
			v := g.gamepadAxisValues[k]
			if -GamepadDeadZone < v && v < GamepadDeadZone {
				continue
			}

			receiver.ReceiveGamepadAxisInput(id, a, v)
		}
	}
}

func (g *Game) Update() error {
	tickIndex++

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

	if gameInstance != nil {
		g.callInputReceiverEvent(gameInstance, cp)
	}
	for _, r := range lv.InputReceivableActors {
		g.callInputReceiverEvent(r, cp)
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
			if d.GetVisibility() {
				d.Draw(screen)
			}
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
