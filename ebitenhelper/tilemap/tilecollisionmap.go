package tilemap

import (
	"log"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/actor"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type TileCollisionMap struct {
	*utility.Array2D[bool]
}

func NewTileCollisionMap(size utility.Point) *TileCollisionMap {
	if size.X < 1 || size.Y < 1 {
		log.Panicln("TileCollisionMap: Invalid size is received")
	}

	return &TileCollisionMap{
		Array2D: utility.NewArray2D[bool](size.X, size.Y),
	}
}

func (t *TileCollisionMap) ToRectangles() utility.Set[*utility.Rectangle] {
	width, height := t.Width(), t.Height()

	// Create and unite rectangles by X axis
	table := utility.NewArray2D[*utility.Rectangle](width, height)
	for y := range height {
		var neighbour *utility.Rectangle
		for x := range width {
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
	for y := range height - 1 {
		x := 0
		for x < width {
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
			l := r.TopLeft().ToVector().Mul(tileSize)
			sz := r.Size().ToVector().Mul(tileSize)

			a := actor.NewBlockingArea()
			a.SetLocation(l)
			a.SetSize(sz)
			if !yield(a) {
				return
			}
		}
	}
}
