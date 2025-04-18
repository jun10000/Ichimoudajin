package widget

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetBase struct {
	fontFamilies []*text.GoTextFaceSource
	fontSize     *float64

	Name            string
	Origin          utility.Vector
	Offset          utility.Vector
	Margin          utility.Inset
	Padding         utility.Inset
	IsHide          bool
	BorderWidth     float64
	BorderColor     color.Color
	BackgroundColor color.Color
	ForegroundColor color.Color
}

func (w *WidgetBase) Init(inherits WidgetBase) {
	if len(w.fontFamilies) == 0 {
		w.fontFamilies = inherits.fontFamilies
	}
	if w.fontSize == nil {
		w.fontSize = inherits.fontSize
	}
}

func (w *WidgetBase) GetFontFamilies() []*text.GoTextFaceSource {
	return w.fontFamilies
}

func (w *WidgetBase) GetFontSize() *float64 {
	return w.fontSize
}

func (w *WidgetBase) SetFontFamilies(fontFamilies []*text.GoTextFaceSource) {
	w.fontFamilies = fontFamilies
}

func (w *WidgetBase) SetFontSize(fontSize *float64) {
	w.fontSize = fontSize
}

func (w *WidgetBase) GetName() string {
	return w.Name
}

func (w *WidgetBase) GetWidgetObject(name string) WidgetObjecter {
	return nil
}

func (w *WidgetBase) MinSize(screenSize *utility.Vector) utility.Vector {
	sy := screenSize.Y
	x := (w.Margin.Left + w.Margin.Right + w.BorderWidth*2 + w.Padding.Left + w.Padding.Right) * sy
	y := (w.Margin.Top + w.Margin.Bottom + w.BorderWidth*2 + w.Padding.Top + w.Padding.Bottom) * sy
	return utility.NewVector(x, y)
}

func (w *WidgetBase) GetTextFace() text.Face {
	var s float64
	if w.fontSize != nil {
		s = *w.fontSize
	}

	fs := make([]text.Face, 0, utility.InitialWidgetFontCap)
	for _, f := range w.fontFamilies {
		fs = append(fs, &text.GoTextFace{
			Source: f,
			Size:   s,
		})
	}

	mf, err := text.NewMultiFace(fs...)
	utility.PanicIfError(err)
	return mf
}

func (w *WidgetBase) GetAlignedArea(screenArea *utility.RectangleF, outerArea *utility.RectangleF, innerSize utility.Vector) (innerArea *utility.RectangleF) {
	screenHeight := screenArea.Size().Y
	outerPos := outerArea.TopLeft()
	outerSize := outerArea.Size()
	retPos := w.Origin.Mul(outerSize).Sub(w.Origin.Mul(innerSize)).Add(w.Offset.MulF(screenHeight)).Add(outerPos)
	return utility.NewRectangleF(retPos.X, retPos.Y, retPos.X+innerSize.X, retPos.Y+innerSize.Y)
}

func (w *WidgetBase) DrawBackground(screen *ebiten.Image, preferredArea utility.RectangleF) {
	s := utility.NewRectangleFFromGoRect(screen.Bounds())
	sy := s.Size().Y
	preferredArea.MinX += (w.Margin.Left + w.BorderWidth/2) * sy
	preferredArea.MinY += (w.Margin.Top + w.BorderWidth/2) * sy
	preferredArea.MaxX -= (w.Margin.Right + w.BorderWidth/2) * sy
	preferredArea.MaxY -= (w.Margin.Bottom + w.BorderWidth/2) * sy
	utility.DrawRectangle(screen, preferredArea.TopLeft(), preferredArea.Size(), float32(w.BorderWidth*sy), w.BorderColor, w.BackgroundColor, true)
}

func (w *WidgetBase) BackgroundToForegroundArea(screenSize *utility.Vector, out *utility.RectangleF) {
	sy := screenSize.Y
	out.MinX += (w.Margin.Left + w.BorderWidth + w.Padding.Left) * sy
	out.MinY += (w.Margin.Top + w.BorderWidth + w.Padding.Top) * sy
	out.MaxX -= (w.Margin.Right + w.BorderWidth + w.Padding.Right) * sy
	out.MaxY -= (w.Margin.Bottom + w.BorderWidth + w.Padding.Bottom) * sy
}
