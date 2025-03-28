package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type TextWidget struct {
	*component.ActorCom
	*component.DrawCom
	*component.WidgetCom

	Text     string
	FontFace *text.GoTextFace
	Color    utility.RGB
}

func (g ActorGeneratorStruct) NewTextWidget(options *NewActorOptions) *TextWidget {
	a := &TextWidget{}
	a.ActorCom = component.NewActorCom(options.Name)
	a.DrawCom = component.NewDrawCom()
	a.WidgetCom = component.NewWidgetCom(options.Location, options.Size)

	a.Text = options.Text.Text
	a.FontFace = &text.GoTextFace{
		Source: nil,
		Size:   options.Text.Size,
	}
	a.Color = options.Text.Color

	a.SetVisibility(options.IsVisible)

	return a
}

// NewLCDTextWidget is another version of NewTextWidget
func (g ActorGeneratorStruct) NewLCDTextWidget(options *NewActorOptions) *TextWidget {
	a := g.NewTextWidget(options)
	a.FontFace.Source = utility.GetFontFromFileP("fonts/LCDPHONE.ttf")
	return a
}

func (w *TextWidget) Draw(screen *ebiten.Image) {
	l := w.GetLocation()
	op := &text.DrawOptions{}
	op.GeoM.Translate(l.X, l.Y)
	op.ColorScale.ScaleWithColor(w.Color)
	text.Draw(screen, w.Text, w.FontFace, op)
}
