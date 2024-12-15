package utility

import (
	"image/color"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var isDebug = false

func init() {
	_, isDebug = os.LookupEnv("debug")
	if isDebug {
		log.Println("Running in debug mode")
	}
}

func RunDebugServer() {
	if isDebug {
		go func() {
			log.Println(http.ListenAndServe(":6060", nil))
		}()
	}
}

func DrawDebugLine(start Vector, end Vector, color color.Color) {
	if isDebug {
		GetGameInstance().AddDrawEvent(func(screen *ebiten.Image) {
			vector.StrokeLine(screen, float32(start.X), float32(start.Y), float32(end.X), float32(end.Y), 2, color, false)
		})
	}
}
