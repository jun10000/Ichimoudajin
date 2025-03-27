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

func (g ActorGeneratorStruct) NewTextWidget(name string, location utility.Vector, rotation float64, scale utility.Vector, size utility.Vector, isVisible bool, extra *ExtraTextInfo) *TextWidget {
	a := &TextWidget{}
	a.ActorCom = component.NewActorCom(name)
	a.DrawCom = component.NewDrawCom()
	a.WidgetCom = component.NewWidgetCom(location, size)

	a.Text = extra.Text
	a.FontFace = &text.GoTextFace{
		Source: nil,
		Size:   extra.Size,
	}
	a.Color = extra.Color

	a.SetVisibility(isVisible)

	return a
}

// NewLCDTextWidget is another version of NewTextWidget
func (g ActorGeneratorStruct) NewLCDTextWidget(name string, location utility.Vector, rotation float64, scale utility.Vector, size utility.Vector, isVisible bool, extra *ExtraTextInfo) *TextWidget {
	a := g.NewTextWidget(name, location, rotation, scale, size, isVisible, extra)
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
