package actor

import (
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type TextBlock struct {
	*component.ActorCom
	*component.DrawTextCom
	utility.Location
	utility.Size
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
