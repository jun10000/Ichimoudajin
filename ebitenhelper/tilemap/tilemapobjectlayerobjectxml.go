package tilemap

import (
	"fmt"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/actor"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type tileMapObjectLayerObjectPropertyXML struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type tileMapObjectLayerObjectTextXML struct {
	Value     string  `xml:",chardata"`
	PixelSize float64 `xml:"pixelsize,attr"`
	Color     string  `xml:"color,attr"`
	HAlign    *string `xml:"halign,attr"`
	VAlign    *string `xml:"valign,attr"`
}

func (t *tileMapObjectLayerObjectTextXML) CreateExtraTextInfo() *actor.NewActorTextOptions {
	op := actor.NewNewActorTextOptions()
	op.Text = t.Value
	op.Size = t.PixelSize

	c, err := utility.HexStringToColor(t.Color, utility.ColorTransparent)
	utility.PanicIfError(err)
	op.Color = c

	if t.HAlign != nil {
		switch *t.HAlign {
		case "center":
			op.AlignH = text.AlignCenter
		case "right":
			op.AlignH = text.AlignEnd
		}
	}

	if t.VAlign != nil {
		switch *t.VAlign {
		case "center":
			op.AlignV = text.AlignCenter
		case "bottom":
			op.AlignV = text.AlignEnd
		}
	}

	return op
}

type tileMapObjectLayerObjectXML struct {
	Name       string                                `xml:"name,attr"`
	Class      string                                `xml:"type,attr"`
	LocationX  float64                               `xml:"x,attr"`
	LocationY  float64                               `xml:"y,attr"`
	SizeX      float64                               `xml:"width,attr"`
	SizeY      float64                               `xml:"height,attr"`
	Visible    *int                                  `xml:"visible,attr"`
	Properties []tileMapObjectLayerObjectPropertyXML `xml:"properties>property"`
	Text       *tileMapObjectLayerObjectTextXML      `xml:"text"`
}

func (o *tileMapObjectLayerObjectXML) NewActor() (utility.Actor, error) {
	op := actor.NewNewActorOptions()
	op.Name = o.Name
	op.Location = utility.NewVector(o.LocationX, o.LocationY)
	op.Size = utility.NewVector(o.SizeX, o.SizeY)
	if o.Visible != nil {
		op.IsVisible = (*o.Visible != 0)
	}
	if o.Text != nil {
		op.Text = o.Text.CreateExtraTextInfo()
	}

	return actor.ActorGenerator.NewActorByTypeName(o.Class, op)
}

func (o *tileMapObjectLayerObjectXML) CreateActor() (utility.Actor, error) {
	ret, err := o.NewActor()
	if err != nil {
		return nil, err
	}

	retv := reflect.ValueOf(ret).Elem()
	for _, property := range o.Properties {
		if m := retv.MethodByName("Set" + property.Name); m.IsValid() {
			// Search method
			mtype := m.Type()
			if mtype.NumIn() != 1 {
				return nil, fmt.Errorf("found invalid method in %s: Set%s", o.Name, property.Name)
			}

			v, err := utility.ConvertFromString(property.Value, mtype.In(0))
			if err != nil {
				return nil, err
			}

			m.Call([]reflect.Value{reflect.ValueOf(v)})
		} else if f := retv.FieldByName(property.Name); f.CanSet() {
			// Search field
			v, err := utility.ConvertFromString(property.Value, f.Type())
			if err != nil {
				return nil, err
			}

			f.Set(reflect.ValueOf(v))
		} else {
			// Search failed
			return nil, fmt.Errorf("found unknown property in %s: %s", o.Name, property.Name)
		}
	}

	return ret, nil
}
