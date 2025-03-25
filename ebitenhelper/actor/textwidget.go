package actor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type TextWidget struct {
	*component.WidgetComponent

	Text     string
	FontFace *text.GoTextFace
	Color    utility.RGB
}

func (g ActorGeneratorStruct) NewTextWidget(location utility.Vector, rotation float64, scale utility.Vector, size utility.Vector, name string, extra any) *TextWidget {
	a := &TextWidget{}
	a.WidgetComponent = component.NewWidgetComponent(location, size, name)

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

	return a
}

// NewLCDTextWidget is another version of NewTextWidget
func (g ActorGeneratorStruct) NewLCDTextWidget(location utility.Vector, rotation float64, scale utility.Vector, size utility.Vector, name string, extra any) *TextWidget {
	a := g.NewTextWidget(location, rotation, scale, size, name, extra)
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
