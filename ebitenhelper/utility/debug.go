package utility

import (
	"image/color"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
			DrawLine(screen, start, end, 2, color, false)
		})
	}
}

func DrawDebugRectangle(topleft Vector, size Vector, color color.Color) {
	if isDebugMode {
		AddDebugDraw(func(screen *ebiten.Image) {
			DrawRectangle(screen, topleft, size, 0, color, color, false)
		})
	}
}

func DrawDebugCircle(center Vector, radius float32, color color.Color) {
	if isDebugMode {
		AddDebugDraw(func(screen *ebiten.Image) {
			DrawCircle(screen, center, radius, 0, color, color, false)
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
			DrawDebugCircle(dt.CenterLocation(), float32(dt.Radius), dc)
		default:
			log.Println("Drawing unknown bounder type")
		}
	}
}

func DrawDebugTraceResult[T ColliderComparable](r *TraceResult[T], b Bounder) {
	if isDebugMode && DebugIsShowTraceResult {
		if !r.IsHit || r.IsFirstHit {
			return
		}

		ls := b.CenterLocation()
		lo := ls.Add(r.TraceOffset.MulF(DebugTraceResultLength))
		lr := lo.Add(r.InputOffset.Sub(r.TraceOffset).MulF(DebugTraceResultLength))
		ln := ls.Add((*r.HitNormal).MulF(DebugTraceResultLength))
		DrawDebugLine(ls, lo, DebugTraceResultOffsetColor)
		DrawDebugLine(lo, lr, DebugTraceResultRemainingOffsetColor)
		DrawDebugLine(ls, ln, DebugTraceResultHitNormalColor)
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
