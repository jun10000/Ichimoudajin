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

	return a
}

// NewLCDTextWidget is another version of NewTextWidget
func (g ActorGeneratorStruct) NewLCDTextWidget(options *NewActorOptions) *TextBlock {
	a := g.NewTextBlock(options)
	a.FontFace.Source = utility.GetFontFromFileP("fonts/LCDPHONE.ttf")
	return a
}

func (w *TextBlock) ZOrder() int {
	return utility.ZOrderWidget
}

func (w *TextBlock) Draw(screen *ebiten.Image) {
	l := w.GetLocation()
	op := &text.DrawOptions{}
	op.GeoM.Translate(l.X, l.Y)
	op.ColorScale.ScaleWithColor(w.Color)
	text.Draw(screen, w.Text, w.FontFace, op)
}
