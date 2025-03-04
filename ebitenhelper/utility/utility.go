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
			tl.Add(ss.Mul(NewVector(-1, -1))),
			tl.Add(ss.Mul(NewVector(0, -1))),
			tl.Add(ss.Mul(NewVector(1, -1))),
			tl.Add(ss.Mul(NewVector(-1, 0))),
			tl.Add(ss.Mul(NewVector(1, 0))),
			tl.Add(ss.Mul(NewVector(-1, 1))),
			tl.Add(ss.Mul(NewVector(0, 1))),
			tl.Add(ss.Mul(NewVector(1, 1))),
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
		ClampFloat(circle.OrgX, rectangle.MinX, rectangle.MaxX),
		ClampFloat(circle.OrgY, rectangle.MinY, rectangle.MaxY))
	r := NewVector(circle.OrgX-p.X, circle.OrgY-p.Y)

	if r.Length() > circle.Radius {
		return false, ZeroVector()
	}

	return true, r.Normalize()
}

func IntersectCircleToCircle(circle1 CircleF, circle2 CircleF) (result bool, normal Vector) {
	d := NewVector(circle1.OrgX-circle2.OrgX, circle1.OrgY-circle2.OrgY)
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
		case *RectangleF:
			return IntersectRectangleToRectangle(v1, *v2)
		case *CircleF:
			r, n := IntersectCircleToRectangle(*v2, v1)
			return r, n.Negate()
		}
	case CircleF:
		switch v2 := dst.(type) {
		case RectangleF:
			return IntersectCircleToRectangle(v1, v2)
		case CircleF:
			return IntersectCircleToCircle(v1, v2)
		case *RectangleF:
			return IntersectCircleToRectangle(v1, *v2)
		case *CircleF:
			return IntersectCircleToCircle(v1, *v2)
		}
	case *RectangleF:
		switch v2 := dst.(type) {
		case RectangleF:
			return IntersectRectangleToRectangle(*v1, v2)
		case CircleF:
			r, n := IntersectCircleToRectangle(v2, *v1)
			return r, n.Negate()
		case *RectangleF:
			return IntersectRectangleToRectangle(*v1, *v2)
		case *CircleF:
			r, n := IntersectCircleToRectangle(*v2, *v1)
			return r, n.Negate()
		}
	case *CircleF:
		switch v2 := dst.(type) {
		case RectangleF:
			return IntersectCircleToRectangle(*v1, v2)
		case CircleF:
			return IntersectCircleToCircle(*v1, v2)
		case *RectangleF:
			return IntersectCircleToRectangle(*v1, *v2)
		case *CircleF:
			return IntersectCircleToCircle(*v1, *v2)
		}
	}

	log.Println("Detected unsupported intersect")
	return false, ZeroVector()
}
