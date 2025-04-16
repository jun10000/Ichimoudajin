package widget

import (
	"encoding/xml"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetBaseXML struct {
	Name            string  `xml:"name,attr"`
	OriginX         float64 `xml:"ox,attr"`
	OriginY         float64 `xml:"oy,attr"`
	Offset          string  `xml:"offset,attr"`
	Margin          string  `xml:"margin,attr"`
	Padding         string  `xml:"padding,attr"`
	IsHide          bool    `xml:"hide,attr"`
	BorderWidth     float64 `xml:"bdwidth,attr"`
	BorderColor     string  `xml:"bdcolor,attr"`
	BackgroundColor string  `xml:"bgcolor,attr"`
	ForegroundColor string  `xml:"fgcolor,attr"`

	FontFiles *string  `xml:"fontfiles,attr"`
	FontSize  *float64 `xml:"fontsize,attr"`
}

func (x WidgetBaseXML) Convert() *WidgetBase {
	fontFamilies := make([]*text.GoTextFaceSource, 0, utility.InitialWidgetFontCap)
	if x.FontFiles != nil {
		for _, s := range strings.Split(*x.FontFiles, ",") {
			fontFamilies = append(fontFamilies, utility.GetFontFromFileP(s))
		}
	}

	bdColor, _ := utility.HexStringToColor(x.BorderColor, utility.ColorTransparent)
	bgColor, _ := utility.HexStringToColor(x.BackgroundColor, utility.ColorTransparent)
	fgColor, _ := utility.HexStringToColor(x.ForegroundColor, utility.ColorWhite)

	return &WidgetBase{
		fontFamilies: fontFamilies,
		fontSize:     x.FontSize,

		Name:            x.Name,
		Origin:          utility.NewVector(x.OriginX, x.OriginY),
		Offset:          utility.NewVectorFromString(x.Offset),
		Margin:          utility.NewInsetFromString(x.Margin),
		Padding:         utility.NewInsetFromString(x.Padding),
		IsHide:          x.IsHide,
		BorderWidth:     x.BorderWidth,
		BorderColor:     bdColor,
		BackgroundColor: bgColor,
		ForegroundColor: fgColor,
	}
}

type WidgetContainerBaseXML struct {
	WidgetBaseXML
	HBoxes  []WidgetHBoxXML   `xml:"hbox"`
	VBoxes  []WidgetVBoxXML   `xml:"vbox"`
	Texts   []WidgetTextXML   `xml:"text"`
	Buttons []WidgetButtonXML `xml:"button"`
}

func (x WidgetContainerBaseXML) Convert() []WidgetObjecter {
	ret := make([]WidgetObjecter, 0,
		len(x.HBoxes)+len(x.VBoxes)+len(x.Texts)+len(x.Buttons))
	for _, c := range x.HBoxes {
		ret = append(ret, c.Convert())
	}
	for _, c := range x.VBoxes {
		ret = append(ret, c.Convert())
	}
	for _, c := range x.Texts {
		ret = append(ret, c.Convert())
	}
	for _, c := range x.Buttons {
		ret = append(ret, c.Convert())
	}

	return ret
}

type WidgetHBoxXML struct {
	WidgetContainerBaseXML
}

func (x WidgetHBoxXML) Convert() *WidgetHBox {
	return &WidgetHBox{
		WidgetContainerBase: &WidgetContainerBase{
			WidgetBase: x.WidgetBaseXML.Convert(),
			Children:   x.WidgetContainerBaseXML.Convert(),
		},
	}
}

type WidgetVBoxXML struct {
	WidgetContainerBaseXML
}

func (x WidgetVBoxXML) Convert() *WidgetVBox {
	return &WidgetVBox{
		WidgetContainerBase: &WidgetContainerBase{
			WidgetBase: x.WidgetBaseXML.Convert(),
			Children:   x.WidgetContainerBaseXML.Convert(),
		},
	}
}

type WidgetTextXML struct {
	WidgetBaseXML
	Text string `xml:",chardata"`
}

func (x WidgetTextXML) Convert() *WidgetText {
	return &WidgetText{
		WidgetBase: x.WidgetBaseXML.Convert(),
		Text:       x.Text,
	}
}

type WidgetButtonXML struct {
	WidgetBaseXML
	Text string `xml:",chardata"`
}

func (x WidgetButtonXML) Convert() *WidgetButton {
	return &WidgetButton{
		WidgetText: &WidgetText{
			WidgetBase: x.WidgetBaseXML.Convert(),
			Text:       x.Text,
		},
	}
}

type WidgetXML struct {
	WidgetContainerBaseXML
	Version int `xml:"version,attr"`
}

func (x WidgetXML) ToActor(name string) *Widget {
	if x.Version != 1 {
		log.Printf("Loaded WidgetXML file version is %d, not 1", x.Version)
	}

	a := &Widget{
		ActorCom: component.NewActorCom(name),
		DrawCom:  component.NewDrawCom(!x.IsHide),
		WidgetContainerBase: &WidgetContainerBase{
			WidgetBase: x.WidgetBaseXML.Convert(),
			Children:   x.WidgetContainerBaseXML.Convert(),
		},
	}
	a.Init(*a.WidgetBase)
	return a
}

func NewWidgetByFile(name string) (*Widget, error) {
	xmlByteData, err := assets.Assets.ReadFile(name + ".xml")
	if err != nil {
		return nil, err
	}

	var xmlData WidgetXML
	err = xml.Unmarshal(xmlByteData, &xmlData)
	if err != nil {
		return nil, err
	}

	return xmlData.ToActor(name), nil
}
