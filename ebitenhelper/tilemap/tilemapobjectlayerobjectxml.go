package tilemap

import (
	"fmt"
	"reflect"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/actor"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type tileMapObjectLayerObjectPropertyXML struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type tileMapObjectLayerObjectTextXML struct {
	PixelSize float64 `xml:"pixelsize,attr"`
	Color     string  `xml:"color,attr"`
	Value     string  `xml:",chardata"`
}

func (t *tileMapObjectLayerObjectTextXML) CreateExtraTextInfo() *actor.ExtraTextInfo {
	c, err := utility.HexColorToRGB(t.Color)
	utility.PanicIfError(err)

	return &actor.ExtraTextInfo{
		Size:  t.PixelSize,
		Text:  t.Value,
		Color: c,
	}
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
	l := utility.NewVector(o.LocationX, o.LocationY)
	r := float64(0)
	s := utility.DefaultScale()
	sz := utility.NewVector(o.SizeX, o.SizeY)
	v := true
	if o.Visible != nil {
		v = (*o.Visible != 0)
	}

	var extra any
	if o.Text != nil {
		extra = o.Text.CreateExtraTextInfo()
	}

	return actor.ActorGenerator.NewActorByName(o.Class, l, r, s, sz, o.Name, extra, v)
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
