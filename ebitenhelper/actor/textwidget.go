package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type TextWidget struct {
	*component.DrawCom
	*component.WidgetCom

	Text     string
	FontFace *text.GoTextFace
	Color    utility.RGB
}

func (g ActorGeneratorStruct) NewTextWidget(location utility.Vector, rotation float64, scale utility.Vector, size utility.Vector, name string, extra any, isVisible bool) *TextWidget {
	a := &TextWidget{}
	a.DrawCom = component.NewDrawCom()
	a.WidgetCom = component.NewWidgetCom(location, size, name)

	if e, ok := extra.(*ExtraTextInfo); ok {
		a.Text = e.Text
		a.FontFace = &text.GoTextFace{
			Source: nil,
			Size:   e.Size,
		}
		a.Color = e.Color
	} else {
		a.Text = "Text"
		a.FontFace = &text.GoTextFace{
			Source: nil,
			Size:   16,
		}
		a.Color = utility.ColorWhite
	}

	a.SetVisibility(isVisible)

	return a
}

// NewLCDTextWidget is another version of NewTextWidget
func (g ActorGeneratorStruct) NewLCDTextWidget(location utility.Vector, rotation float64, scale utility.Vector, size utility.Vector, name string, extra any, isVisible bool) *TextWidget {
	a := g.NewTextWidget(location, rotation, scale, size, name, extra, isVisible)
	a.FontFace.Source = utility.GetFontFromFileP("fonts/LCDPHONE.ttf")
	return a
}

func (w *TextWidget) Draw(screen *ebiten.Image) {
	l := w.GetLocation()
	op := &text.DrawOptions{}
	op.GeoM.Translate(l.X, l.Y)
	op.ColorScale.ScaleWithColor(w.Color.ToRGBA(0xff))
	text.Draw(screen, w.Text, w.FontFace, op)
}
