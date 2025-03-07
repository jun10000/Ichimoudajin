package utility

import (
	"image"
	"log"
	"math"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jun10000/Ichimoudajin/assets"
)

type PressState int

const (
	PressStatePressed PressState = iota
	PressStateReleased
	PressStatePressing
)

type BounderType int

const (
	BounderTypeRectangle BounderType = iota
	BounderTypeCircle
)

type Empty struct{}

func GetImageFile(filename string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFileSystem(assets.Assets, filename)
	PanicIfError(err)
	return image
}

func GetSubImage(parentimage *ebiten.Image, location Point, size Point) *ebiten.Image {
	if parentimage == nil {
		return nil
	}

	r := image.Rect(location.X, location.Y, location.X+size.X, location.Y+size.Y)
	return parentimage.SubImage(r).(*ebiten.Image)
}

func DrawImage(dst *ebiten.Image, src *ebiten.Image, transform Transformer) {
	if dst == nil || src == nil {
		return
	}

	tl := transform.GetLocation()
	tr := transform.GetRotation()
	ts := transform.GetScale()

	var ls []Vector
	if !GetLevel().IsLooping {
		ls = []Vector{tl}
	} else {
		ss := GetScreenSize().ToVector()
		ls = []Vector{
			tl,
			tl.AddXY(-ss.X, -ss.Y),
			tl.AddXY(0, -ss.Y),
			tl.AddXY(ss.X, -ss.Y),
			tl.AddXY(-ss.X, 0),
			tl.AddXY(ss.X, 0),
			tl.AddXY(-ss.X, ss.Y),
			tl.AddXY(0, ss.Y),
			tl.AddXY(ss.X, ss.Y),
		}
	}

	for _, l := range ls {
		o := &ebiten.DrawImageOptions{}
		o.GeoM.Scale(ts.X, ts.Y)
		o.GeoM.Rotate(tr)
		o.GeoM.Translate(l.X, l.Y)

		dst.DrawImage(src, o)
	}
}

func PanicIfError(err error) {
	if err != nil {
		log.Panic(err)
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

	if lv == nil || !lv.IsLooping {
		return r
	}

	ss := GetScreenSize().ToVector()
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
	i := slices.Index(slice, item)
	if i == -1 {
		return slice
	}

	return append(slice[:i], slice[i+1:]...)
}

func IntersectRectangleToRectangle(rectangle1 *RectangleF, rectangle2 *RectangleF) (result bool, normal *Vector) {
	xleft := rectangle1.MaxX - rectangle2.MinX
	if xleft < 0 {
		return false, nil
	}

	xright := rectangle2.MaxX - rectangle1.MinX
	if xright < 0 {
		return false, nil
	}

	ytop := rectangle1.MaxY - rectangle2.MinY
	if ytop < 0 {
		return false, nil
	}

	ybottom := rectangle2.MaxY - rectangle1.MinY
	if ybottom < 0 {
		return false, nil
	}

	if math.Min(xleft, xright) > math.Min(ytop, ybottom) {
		if ytop > ybottom {
			return true, DownVectorPtr()
		} else {
			return true, UpVectorPtr()
		}
	} else {
		if xleft > xright {
			return true, RightVectorPtr()
		} else {
			return true, LeftVectorPtr()
		}
	}
}

func IntersectCircleToRectangle(circle *CircleF, rectangle *RectangleF, isReverse bool) (result bool, normal *Vector) {
	p := NewVector(
		ClampFloat(circle.OrgX, rectangle.MinX, rectangle.MaxX),
		ClampFloat(circle.OrgY, rectangle.MinY, rectangle.MaxY))

	var r Vector
	if isReverse {
		r = NewVector(p.X-circle.OrgX, p.Y-circle.OrgY)
	} else {
		r = NewVector(circle.OrgX-p.X, circle.OrgY-p.Y)
	}
	rll := r.Length2()

	if rll > (circle.Radius * circle.Radius) {
		return false, nil
	}

	n := r.DivF(math.Sqrt(rll))
	return true, &n
}

func IntersectCircleToCircle(circle1 *CircleF, circle2 *CircleF) (result bool, normal *Vector) {
	d := NewVector(circle1.OrgX-circle2.OrgX, circle1.OrgY-circle2.OrgY)
	dll := d.Length2()
	if dll > ((circle1.Radius + circle2.Radius) * (circle1.Radius + circle2.Radius)) {
		return false, nil
	}

	n := d.DivF(math.Sqrt(dll))
	return true, &n
}
