package utility

import (
	"image/color"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var isDebugMode = false

func init() {
	_, isDebugMode = os.LookupEnv("debug")
	if isDebugMode {
		log.Println("Running in debug mode")
	}
}

func IsDebugMode() bool {
	return isDebugMode
}

func RunDebugServer() {
	if isDebugMode {
		go func() {
			log.Println(http.ListenAndServe(":6060", nil))
		}()
	}
}

func AddDebugDraw(event func(*ebiten.Image)) {
	GetLevel().AddDebugDraw(event)
}

func DrawDebugLine(start Vector, end Vector, color color.Color) {
	if isDebugMode {
		AddDebugDraw(func(screen *ebiten.Image) {
			vector.StrokeLine(screen, float32(start.X), float32(start.Y), float32(end.X), float32(end.Y), 2, color, false)
		})
	}
}

func DrawDebugRectangle(topleft Vector, size Vector, color color.Color) {
	if isDebugMode {
		AddDebugDraw(func(screen *ebiten.Image) {
			vector.DrawFilledRect(screen, float32(topleft.X), float32(topleft.Y), float32(size.X), float32(size.Y), color, false)
		})
	}
}

func DrawDebugCircle(center Vector, radius float64, color color.Color) {
	if isDebugMode {
		AddDebugDraw(func(screen *ebiten.Image) {
			vector.DrawFilledCircle(screen, float32(center.X), float32(center.Y), float32(radius), color, false)
		})
	}
}

func DrawDebugText(topleft Vector, text string) {
	if isDebugMode {
		l := topleft.Trunc()
		AddDebugDraw(func(screen *ebiten.Image) {
			ebitenutil.DebugPrintAt(screen, text, l.X, l.Y)
		})
	}
}
