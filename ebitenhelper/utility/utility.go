package utility

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TickCount          int     = 60
	TickDuration       float64 = 1.0 / float64(TickCount)
	TraceSafeDistance  int     = 3
	AIValidOffset      float64 = 0.5
	AIInitialResultCap int     = 1024
	AIRequestCap       int     = 1

	IsShowDebugMoverLocation bool = false
	IsShowDebugTraceDistance bool = false
	IsShowDebugAIPath        bool = true
)

var (
	ColorGray   = color.RGBA{R: 128, G: 128, B: 128}
	ColorRed    = color.RGBA{R: 255, G: 8}
	ColorYellow = color.RGBA{R: 255, G: 255}
	ColorGreen  = color.RGBA{G: 255}
	ColorBlue   = color.RGBA{G: 128, B: 255}
)

type PressState int

const (
	PressStatePressed PressState = iota
	PressStateReleased
	PressStatePressing
)

func NewRectangle(location Point, size Point) image.Rectangle {
	return image.Rect(location.X, location.Y, location.X+size.X, location.Y+size.Y)
}

func GetSubImage(parentimage *ebiten.Image, location Point, size Point) *ebiten.Image {
	if parentimage == nil {
		return nil
	}

	return parentimage.SubImage(NewRectangle(location, size)).(*ebiten.Image)
}

func DrawImage(dst *ebiten.Image, src *ebiten.Image, transform Transformer) {
	if dst == nil || src == nil {
		return
	}

	tl := transform.GetLocation()
	tr := transform.GetRotation()
	ts := transform.GetScale()

	ls := []Vector{tl}
	if GetLevel().IsLooping {
		ss := GetGameInstance().ScreenSize.ToVector()
		ls = append(ls,
			tl.Add(ss.Mul(NewVector(-1, -1))),
			tl.Add(ss.Mul(NewVector(0, -1))),
			tl.Add(ss.Mul(NewVector(1, -1))),
			tl.Add(ss.Mul(NewVector(-1, 0))),
			tl.Add(ss.Mul(NewVector(1, 0))),
			tl.Add(ss.Mul(NewVector(-1, 1))),
			tl.Add(ss.Mul(NewVector(0, 1))),
			tl.Add(ss.Mul(NewVector(1, 1))),
		)
	}

	for _, l := range ls {
		o := &ebiten.DrawImageOptions{}
		o.GeoM.Scale(ts.X, ts.Y)
		o.GeoM.Rotate(tr)
		o.GeoM.Translate(l.X, l.Y)

		dst.DrawImage(src, o)
	}
}

func ExitIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ClampFloat(value float64, min float64, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}

	return value
}

func ClampLocation(location Vector) Vector {
	r := location
	lv := GetLevel()
	gi := GetGameInstance()

	if lv == nil || !lv.IsLooping || gi == nil {
		return r
	}

	ss := gi.ScreenSize.ToVector()
	r = r.Mod(ss)
	if r.X < 0 {
		r.X += ss.X
	}
	if r.Y < 0 {
		r.Y += ss.Y
	}

	return r
}

func RemoveSliceItem[T comparable](slice []T, item T) []T {
	r := make([]T, len(slice))
	for _, v := range slice {
		if v == item {
			continue
		}

		r = append(r, v)
	}

	return r
}

func IntersectRectangleToRectangle(rectangle1 RectangleF, rectangle2 RectangleF) (result bool, normal Vector) {
	xleft := rectangle1.MaxX - rectangle2.MinX
	xright := rectangle2.MaxX - rectangle1.MinX
	ytop := rectangle1.MaxY - rectangle2.MinY
	ybottom := rectangle2.MaxY - rectangle1.MinY

	if xleft < 0 || xright < 0 || ytop < 0 || ybottom < 0 {
		return false, ZeroVector()
	}

	isright := xleft > xright
	isbottom := ytop > ybottom
	isy := math.Min(xleft, xright) > math.Min(ytop, ybottom)

	if isy {
		if isbottom {
			return true, NewVector(0, 1)
		} else {
			return true, NewVector(0, -1)
		}
	} else {
		if isright {
			return true, NewVector(1, 0)
		} else {
			return true, NewVector(-1, 0)
		}
	}
}

func IntersectCircleToRectangle(circle CircleF, rectangle RectangleF) (result bool, normal Vector) {
	p := NewVector(
		ClampFloat(circle.Origin.X, rectangle.MinX, rectangle.MaxX),
		ClampFloat(circle.Origin.Y, rectangle.MinY, rectangle.MaxY))
	r := circle.Origin.Sub(p)

	if r.Length() > circle.Radius {
		return false, ZeroVector()
	}

	return true, r.Normalize()
}

func IntersectCircleToCircle(circle1 CircleF, circle2 CircleF) (result bool, normal Vector) {
	d := circle1.Origin.Sub(circle2.Origin)
	if d.Length() > circle1.Radius+circle2.Radius {
		return false, ZeroVector()
	}

	return true, d.Normalize()
}

/*
Intersect supports following type combinations
  - RectangleF -> RectangleF
  - RectangleF -> CircleF
  - CircleF -> RectangleF
  - CircleF -> CircleF
*/
func Intersect(src Bounder, dst Bounder) (result bool, normal Vector) {
	switch v1 := src.(type) {
	case RectangleF:
		switch v2 := dst.(type) {
		case RectangleF:
			return IntersectRectangleToRectangle(v1, v2)
		case CircleF:
			r, n := IntersectCircleToRectangle(v2, v1)
			return r, n.Negate()
		}
	case CircleF:
		switch v2 := dst.(type) {
		case RectangleF:
			return IntersectCircleToRectangle(v1, v2)
		case CircleF:
			return IntersectCircleToCircle(v1, v2)
		}
	}

	log.Println("Detected unsupported intersect")
	return false, ZeroVector()
}
