package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type TextBlock struct {
	*component.ActorCom
	*component.DrawCom
	utility.Location

	Size     utility.Vector
	Text     string
	FontFace *text.GoTextFace
	Color    utility.RGB
	AlignH   text.Align
	AlignV   text.Align
}

func (g ActorGeneratorStruct) NewTextBlock(options *NewActorOptions) *TextBlock {
	a := &TextBlock{}
	a.ActorCom = component.NewActorCom(options.Name)
	a.DrawCom = component.NewDrawCom(options.IsVisible)
	a.Location = utility.NewLocation(options.Location)

	a.Size = options.Size
	a.Text = options.Text.Text
	a.FontFace = &text.GoTextFace{
		Source: nil,
		Size:   options.Text.Size,
	}
	a.Color = options.Text.Color
	a.AlignH = options.Text.AlignH
	a.AlignV = options.Text.AlignV

	return a
}

// NewLCDTextWidget is another version of NewTextWidget
func (g ActorGeneratorStruct) NewLCDTextWidget(options *NewActorOptions) *TextBlock {
	a := g.NewTextBlock(options)
	a.FontFace.Source = utility.GetFontFromFileP("fonts/LCDPHONE.ttf")
	return a
}

func (a *TextBlock) ZOrder() int {
	return utility.ZOrderWidget
}

func (a *TextBlock) Draw(screen *ebiten.Image) {
	l := a.GetLocation()
	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(a.Color)
	op.PrimaryAlign = a.AlignH
	op.SecondaryAlign = a.AlignV

	x := l.X
	switch a.AlignH {
	case text.AlignCenter:
		x += a.Size.X / 2
	case text.AlignEnd:
		x += a.Size.X
	}

	y := l.Y
	switch a.AlignV {
	case text.AlignCenter:
		y += a.Size.Y / 2
	case text.AlignEnd:
		y += a.Size.Y
	}

	op.GeoM.Translate(x, y)
	text.Draw(screen, a.Text, a.FontFace, op)
}
