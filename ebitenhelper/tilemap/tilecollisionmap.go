package tilemap

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/actor"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

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

func (t *TileCollisionMap) ToRectangles() utility.Set[*utility.Rectangle] {
	// Create and unite rectangles by X axis
	table := utility.NewArray2D[*utility.Rectangle](t.size.X, t.size.Y)
	for y := range t.size.Y {
		var neighbour *utility.Rectangle
		for x := range t.size.X {
			if t.Get(x, y) {
				if neighbour == nil {
					neighbour = utility.NewRectangle(x, y, x, y)
				} else {
					neighbour.MaxX++
				}
				table.Set(x, y, neighbour)
			} else {
				neighbour = nil
			}
		}
	}

	// Unite rectangles by Y axis
	for y := range t.size.Y - 1 {
		x := 0
		for x < t.size.X {
			rself := table.Get(x, y)
			if rself == nil {
				x++
				continue
			}

			rtarget := table.Get(x, y+1)
			if rtarget == nil || rself.MinX != rtarget.MinX || rself.MaxX != rtarget.MaxX {
				x = rself.MaxX + 1
				continue
			}

			rself.MaxY++
			for xcur := rself.MinX; xcur <= rself.MaxX; xcur++ {
				table.Set(xcur, y+1, rself)
			}

			x = rself.MaxX + 1
		}
	}

	// Exclude duplicated rectangles
	ret := make(utility.Set[*utility.Rectangle])
	for _, r := range table.Range() {
		if r == nil {
			continue
		}
		ret.Add(r)
	}

	return ret
}

func (t *TileCollisionMap) ToBlockingAreas(tileSize utility.Vector) func(yield func(*actor.BlockingArea) bool) {
	return func(yield func(*actor.BlockingArea) bool) {
		for r := range t.ToRectangles() {
			lx := float64(r.MinX) * tileSize.X
			ly := float64(r.MinY) * tileSize.Y
			sx := float64(r.MaxX-r.MinX+1) * tileSize.X
			sy := float64(r.MaxY-r.MinY+1) * tileSize.Y

			a := actor.NewBlockingArea()
			a.SetLocation(utility.NewVector(lx, ly))
			a.SetSize(utility.NewVector(sx, sy))
			if !yield(a) {
				return
			}
		}
	}
}
