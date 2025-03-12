package actor

import "github.com/jun10000/Ichimoudajin/ebitenhelper/utility"

type TileCollisionMap struct {
	*utility.Array2D[bool]
	size utility.Point
}

func NewTileCollisionMap(size utility.Point) *TileCollisionMap {
	return &TileCollisionMap{
		Array2D: utility.NewArray2D[bool](size.X, size.Y),
		size:    size,
	}
}

func (t *TileCollisionMap) ToBlockingAreas(tileSize utility.Vector) func(yield func(*BlockingArea) bool) {
	return func(yield func(*BlockingArea) bool) {
		for x := range t.size.X {
			for y := range t.size.Y {
				if !t.Get(x, y) {
					continue
				}

				lx := float64(x) * tileSize.X
				ly := float64(y) * tileSize.Y
				sz := tileSize

				a := NewBlockingArea()
				a.SetLocation(utility.NewVector(lx, ly))
				a.Size = sz
				if !yield(a) {
					return
				}
			}
		}
	}
}
