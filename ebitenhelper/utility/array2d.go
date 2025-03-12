package utility

import "log"

type Array2D[T any] struct {
	slice  []T
	width  int
	height int
}

func NewArray2D[T any](width int, height int) *Array2D[T] {
	return &Array2D[T]{
		slice:  make([]T, width*height),
		width:  width,
		height: height,
	}
}

func (a *Array2D[T]) getIndex(x, y int) int {
	return y*a.width + x
}

func (a *Array2D[T]) getXY(index int) (X, Y int) {
	return index % a.width, index / a.width
}

func (a *Array2D[T]) Width() int {
	return a.width
}

func (a *Array2D[T]) Height() int {
	return a.height
}

func (a *Array2D[T]) Get(x, y int) T {
	if x < 0 || x >= a.width || y < 0 || y >= a.height {
		log.Panicln("Array2D detects out of range index")
	}

	return a.slice[a.getIndex(x, y)]
}

func (a *Array2D[T]) Set(x, y int, value T) {
	if x < 0 || x >= a.width || y < 0 || y >= a.height {
		log.Panicln("Array2D detects out of range index")
	}

	a.slice[a.getIndex(x, y)] = value
}

func (a *Array2D[T]) Range() func(yield func(Point, T) bool) {
	return func(yield func(Point, T) bool) {
		for i, v := range a.slice {
			x, y := a.getXY(i)
			if !yield(NewPoint(x, y), v) {
				return
			}
		}
	}
}
