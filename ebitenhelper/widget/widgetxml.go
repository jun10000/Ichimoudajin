package widget

import (
	"encoding/xml"
	"log"

	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/component"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type WidgetCommonAttributesXML struct {
	Name    string  `xml:"name,attr"`
	OriginX float64 `xml:"originx,attr"`
	OriginY float64 `xml:"originy,attr"`
	X       float64 `xml:"x,attr"`
	Y       float64 `xml:"y,attr"`
	IsHide  bool    `xml:"hide,attr"`
}

func (x WidgetCommonAttributesXML) Convert() *WidgetCommonFields {
	return &WidgetCommonFields{
		Name:     x.Name,
		Origin:   utility.NewVector(x.OriginX, x.OriginY),
		Position: utility.NewVector(x.X, x.Y),
		IsHide:   x.IsHide,
	}
}

type WidgetContainerElementsXML struct {
	HBoxes  []WidgetHBoxXML   `xml:"hbox"`
	VBoxes  []WidgetVBoxXML   `xml:"vbox"`
	Texts   []WidgetTextXML   `xml:"text"`
	Buttons []WidgetButtonXML `xml:"button"`
}

func (x WidgetContainerElementsXML) Convert() []WidgetObject {
	ret := make([]WidgetObject, 0,
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
	WidgetCommonAttributesXML
	WidgetContainerElementsXML
}

func (x WidgetHBoxXML) Convert() *WidgetHBox {
	return &WidgetHBox{
		WidgetCommonFields: x.WidgetCommonAttributesXML.Convert(),
		Children:           x.WidgetContainerElementsXML.Convert(),
	}
}

type WidgetVBoxXML struct {
	WidgetCommonAttributesXML
	WidgetContainerElementsXML
}

func (x WidgetVBoxXML) Convert() *WidgetVBox {
	return &WidgetVBox{
		WidgetCommonFields: x.WidgetCommonAttributesXML.Convert(),
		Children:           x.WidgetContainerElementsXML.Convert(),
	}
}

type WidgetTextXML struct {
	WidgetCommonAttributesXML
	Text string `xml:",chardata"`
}

func (x WidgetTextXML) Convert() *WidgetText {
	return &WidgetText{
		WidgetCommonFields: x.WidgetCommonAttributesXML.Convert(),
		Text:               x.Text,
	}
}

type WidgetButtonXML struct {
	WidgetCommonAttributesXML
	Text string `xml:",chardata"`
}

func (x WidgetButtonXML) Convert() *WidgetButton {
	return &WidgetButton{
		WidgetCommonFields: x.WidgetCommonAttributesXML.Convert(),
		Text:               x.Text,
	}
}

type WidgetXML struct {
	Version int  `xml:"version,attr"`
	IsHide  bool `xml:"hide,attr"`
	WidgetContainerElementsXML
}

func (x WidgetXML) ToActor(name string) *Widget {
	if x.Version != 1 {
		log.Printf("Loaded WidgetXML file version is %d, not 1", x.Version)
	}

	return &Widget{
		ActorCom: component.NewActorCom(name),
		DrawCom:  component.NewDrawCom(!x.IsHide),
		Children: x.Convert(),
	}
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
