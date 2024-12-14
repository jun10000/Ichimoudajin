//go:build !debug

package utility

import (
	"image/color"
)

func RunDebugServer()                                           {}
func DrawDebugLine(start Vector, end Vector, color color.Color) {}
