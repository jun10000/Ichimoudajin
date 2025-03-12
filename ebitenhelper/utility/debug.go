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
	if isDebugMode {
		GetLevel().AddDebugDraw(event)
	}
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

func DrawDebugLocation(location Vector) {
	if isDebugMode && DebugIsShowLocation {
		DrawDebugText(location.Add(DebugLocationTextOffset), location.String())
	}
}

func DrawDebugTraceDistance(target Bounder, distance int) {
	if isDebugMode && DebugIsShowTraceDistance {
		dc, ok := DebugTraceDistanceColors[distance]
		if !ok {
			return
		}

		switch dt := target.(type) {
		case *RectangleF:
			DrawDebugRectangle(dt.TopLeft(), dt.Size(), dc)
		case *CircleF:
			DrawDebugCircle(dt.CenterLocation(), dt.Radius, dc)
		default:
			log.Println("Drawing unknown bounder type")
		}
	}
}

func DrawDebugTraceResult(r *TraceResult, b Bounder) {
	if isDebugMode && DebugIsShowTraceResult {
		if !r.IsHit || r.IsFirstHit {
			return
		}

		ls := b.CenterLocation()
		lh := ls.Add(r.TraceOffset.MulF(DebugTraceResultLength))
		DrawDebugLine(ls, lh, DebugColorGreen)
		DrawDebugLine(lh, lh.Add(r.InputOffset.Sub(r.TraceOffset).MulF(DebugTraceResultLength)), DebugColorRed)
		DrawDebugLine(ls, ls.Add((*r.HitNormal).MulF(DebugTraceResultLength)), DebugColorBlue)
	}
}

func DrawDebugAIPath(path []Point) {
	if isDebugMode && DebugIsShowAIPath {
		l := GetLevel()
		for _, p := range path {
			DrawDebugRectangle(l.PFToRealLocation(p, false), l.AIGridSize, DebugAIPathColor)
		}
	}
}
