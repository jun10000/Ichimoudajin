package actor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type TextBlock struct {
	*component.ActorCom
	*component.DrawTextCom
	utility.Location
	utility.Size

	BorderWidth float32
	BorderColor color.Color
	FillColor   color.Color
}

func (g ActorGeneratorStruct) NewTextBlock(options *NewActorOptions) *TextBlock {
	a := &TextBlock{}
	a.ActorCom = component.NewActorCom(options.Name)
	a.DrawTextCom = component.NewDrawTextCom(a, options.IsVisible)
	a.Location = utility.NewLocation(options.Location)
	a.Size = utility.NewSize(options.Size)

	a.FontSize = options.Text.Size
	a.Text = options.Text.Text
	a.TextColor = options.Text.Color
	a.TextAlignH = options.Text.AlignH
	a.TextAlignV = options.Text.AlignV

	a.BorderWidth = 0
	a.BorderColor = utility.ColorTransparent
	a.FillColor = utility.ColorTransparent

	return a
}

// NewLCDTextWidget is another version of NewTextWidget
func (g ActorGeneratorStruct) NewLCDTextWidget(options *NewActorOptions) *TextBlock {
	a := g.NewTextBlock(options)
	a.Font = utility.GetFontFromFileP("fonts/LCDPHONE.ttf")
	return a
}

func (a *TextBlock) ZOrder() int {
	return utility.ZOrderWidget
}

func (a *TextBlock) Draw(screen *ebiten.Image) {
	l := a.GetLocation()
	sz := a.GetSize()

	utility.DrawRectangle(screen, l, sz, a.BorderWidth, a.BorderColor, a.FillColor, true)
	a.DrawTextCom.Draw(screen)
}
