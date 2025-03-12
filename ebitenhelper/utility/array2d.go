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

func (a *Array2D[T]) Get(x, y int) T {
	if x < 0 || x >= a.width || y < 0 || y >= a.height {
		log.Panicln("Array2D detects out of range index")
	}

	return a.slice[y*a.width+x]
}

func (a *Array2D[T]) Set(x, y int, value T) {
	if x < 0 || x >= a.width || y < 0 || y >= a.height {
		log.Panicln("Array2D detects out of range index")
	}

	a.slice[y*a.width+x] = value
}
