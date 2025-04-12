package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetCommonFields struct {
	font *text.GoTextFace

	Name     string
	Origin   utility.Vector
	Position utility.Vector
	IsHide   bool
}

func (f *WidgetCommonFields) Init(inherits WidgetCommonFields) {
	if f.font == nil {
		f.font = inherits.font
	}
}

func (f *WidgetCommonFields) GetFont() *text.GoTextFace {
	return f.font
}

func (f *WidgetCommonFields) SetFont(font *text.GoTextFace) {
	f.font = font
}

type WidgetContainerFields struct {
	*WidgetCommonFields
	Children []WidgetObjecter
}

func (f *WidgetContainerFields) Init(inherits WidgetCommonFields) {
	if f.font == nil {
		f.font = inherits.font
	} else {
		inherits.font = f.font
	}

	for _, o := range f.Children {
		o.Init(inherits)
	}
}

func (f *WidgetContainerFields) SetFont(font *text.GoTextFace) {
	oldFont := f.font
	f.font = font
	for _, o := range f.Children {
		if o.GetFont() == oldFont {
			o.SetFont(font)
		}
	}
}

type WidgetHBox struct {
	*WidgetContainerFields
}

func (w *WidgetHBox) MinSize() utility.Vector {
	ret := utility.ZeroVector()

	for _, o := range w.Children {
		s := o.MinSize()
		ret.X += s.X
		if s.Y > ret.Y {
			ret.Y = s.Y
		}
	}

	return ret
}

func (w *WidgetHBox) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	preferredArea.MinX += w.Position.X
	preferredArea.MinY += w.Position.Y

	for _, o := range w.Children {
		s := o.MinSize()
		preferredArea.MaxX = preferredArea.MinX + s.X
		preferredArea.MaxY = preferredArea.MinY + s.Y
		o.Draw(screen, preferredArea)
		preferredArea.MinX = preferredArea.MaxX
	}
}

type WidgetVBox struct {
	*WidgetContainerFields
}

func (w *WidgetVBox) MinSize() utility.Vector {
	return utility.ZeroVector()
}

func (w *WidgetVBox) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
}

type WidgetText struct {
	*WidgetCommonFields
	Text string
}

func (w *WidgetText) MinSize() utility.Vector {
	if w.font == nil {
		return utility.ZeroVector()
	}

	x, y := text.Measure(w.Text, w.font, 0)
	return utility.NewVector(x, y)
}

func (w *WidgetText) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {
	if w.IsHide || w.font == nil {
		return
	}

	innerSize := w.MinSize()
	outerSize := preferredArea.Size()
	origin := w.Origin.DivF(100)
	offset := origin.Mul(outerSize).Sub(origin.Mul(innerSize)).Add(w.Position)

	op := &text.DrawOptions{}
	op.GeoM.Translate(preferredArea.MinX+offset.X, preferredArea.MinY+offset.Y)

	text.Draw(screen, w.Text, w.font, op)
}

type WidgetButton struct {
	*WidgetCommonFields
	Text string
}

func (w *WidgetButton) MinSize() utility.Vector {
	return utility.ZeroVector()
}

func (w *WidgetButton) Draw(screen *ebiten.Image, preferredArea utility.RectangleF) {

}

type Widget struct {
	*component.ActorCom
	*component.DrawCom
	*WidgetContainerFields
}

func (a *Widget) Draw(screen *ebiten.Image) {
	ri := screen.Bounds()
	r := utility.NewRectangleF(float64(ri.Min.X), float64(ri.Min.Y), float64(ri.Max.X), float64(ri.Max.Y))
	for _, o := range a.Children {
		o.Draw(screen, *r)
	}
}

func (a *Widget) ZOrder() int {
	return utility.ZOrderWidget
}
