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

	return slices.Delete(slice, i, i+1)
}

func GetColliderBounds[T ColliderComparable](colliders *Smap[T, [9]Bounder], excepts Set[T]) func(yield func(Bounder) bool) {
	return func(yield func(Bounder) bool) {
		loop := GetLevel().IsLooping
		for c, bs := range colliders.Range() {
			if excepts != nil && excepts.Contains(c) {
				continue
			}

			if loop {
				for i := range 9 {
					if !yield(bs[i]) {
						return
					}
				}
			} else {
				if !yield(bs[0]) {
					return
				}
			}
		}
	}
}

func Intersect[T ColliderComparable](colliders *Smap[T, [9]Bounder], target Bounder, excepts Set[T]) (result bool, normal *Vector) {
	for b := range GetColliderBounds(colliders, excepts) {
		r, n := target.IntersectTo(b)
		if r {
			return true, n
		}
	}

	return false, nil
}

func Trace[T ColliderComparable](colliders *Smap[T, [9]Bounder], target Bounder, offset Vector, excepts Set[T]) (rOnHitDistance int, rOffset Vector, rNormal *Vector, rIsHit bool) {
	ol, on := offset.Decompose()
	oli := int(math.Trunc(ol)) + 1

	for i := 0; i <= oli; i++ {
		v := on.MulF(float64(i))
		bo := target.Offset(v.X, v.Y, nil)
		r, n := Intersect(colliders, bo, excepts)
		if r {
			DrawDebugTraceDistance(target, i)
			if i == 0 {
				return 0, ZeroVector(), n, true
			} else {
				o := on.MulF(float64(i - 1))
				return i, o, n, true
			}
		}
	}

	return oli + 1, offset, nil, false
}
