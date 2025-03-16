package utility

import "image/color"

type RGB struct {
	R, G, B uint8
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	return
}

func (c RGB) ToRGBA(alpha uint8) color.RGBA {
	return color.RGBA{c.R, c.G, c.B, alpha}
}
