package utility

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameInstance struct {
	FirstLevel *Level

	WindowTitle  string
	ScreenWidth  int
	ScreenHeight int
	CurrentLevel *Level

	Temp_Keys []ebiten.Key
}

func PlayGame(title string, screen_w int, screen_h int, firstlevel *Level) {
	g := &GameInstance{
		FirstLevel:   firstlevel,
		WindowTitle:  title,
		ScreenWidth:  screen_w,
		ScreenHeight: screen_h,
		CurrentLevel: firstlevel,
	}
	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle(g.WindowTitle)

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *GameInstance) Update() error {
	// OnKeyPressed
	g.Temp_Keys = inpututil.AppendJustPressedKeys(g.Temp_Keys[:0])
	for _, k := range g.Temp_Keys {
		for _, p := range g.CurrentLevel.Pawns {
			p.Event_KeyPressed(k)
		}
	}

	// OnKeyReleased
	g.Temp_Keys = inpututil.AppendJustReleasedKeys(g.Temp_Keys[:0])
	for _, k := range g.Temp_Keys {
		for _, p := range g.CurrentLevel.Pawns {
			p.Event_KeyReleased(k)
		}
	}

	// OnKeyPressing
	g.Temp_Keys = inpututil.AppendPressedKeys(g.Temp_Keys[:0])
	for _, k := range g.Temp_Keys {
		for _, p := range g.CurrentLevel.Pawns {
			p.Event_KeyPressing(k)
		}
	}

	return nil
}

func (g *GameInstance) Draw(screen *ebiten.Image) {
	for _, a := range g.CurrentLevel.Actors {
		a.Draw(screen)
	}
	for _, a := range g.CurrentLevel.Pawns {
		a.Draw(screen)
	}
}

func (g *GameInstance) Layout(width int, height int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
