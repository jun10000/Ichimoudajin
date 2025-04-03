package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type DrawTextCom struct {
	*DrawCom
	parent utility.StaticRectangler

	Font       *text.GoTextFaceSource
	FontSize   float64
	Text       string
	TextColor  utility.RGB
	TextAlignH text.Align
	TextAlignV text.Align
}

func NewDrawTextCom(parent utility.StaticRectangler, isVisible bool) *DrawTextCom {
	return &DrawTextCom{
		DrawCom: NewDrawCom(isVisible),
		parent:  parent,

		FontSize:   16,
		TextAlignH: text.AlignStart,
		TextAlignV: text.AlignStart,
	}
}

func (c *DrawTextCom) Draw(screen *ebiten.Image) {
	l := c.parent.GetLocation()
	sz := c.parent.GetSize()

	lx := l.X
	switch c.TextAlignH {
	case text.AlignCenter:
		lx += sz.X / 2
	case text.AlignEnd:
		lx += sz.X
	}

	ly := l.Y
	switch c.TextAlignV {
	case text.AlignCenter:
		ly += sz.Y / 2
	case text.AlignEnd:
		ly += sz.Y
	}

	ff := &text.GoTextFace{
		Source: c.Font,
		Size:   c.FontSize,
	}

	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(c.TextColor)
	op.PrimaryAlign = c.TextAlignH
	op.SecondaryAlign = c.TextAlignV
	op.GeoM.Translate(lx, ly)

	text.Draw(screen, c.Text, ff, op)
}
