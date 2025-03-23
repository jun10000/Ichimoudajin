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

type tileMapObjectLayerObjectXML struct {
	Name       string                                `xml:"name,attr"`
	Class      string                                `xml:"type,attr"`
	LocationX  float64                               `xml:"x,attr"`
	LocationY  float64                               `xml:"y,attr"`
	SizeX      float64                               `xml:"width,attr"`
	SizeY      float64                               `xml:"height,attr"`
	Properties []tileMapObjectLayerObjectPropertyXML `xml:"properties>property"`
}

func (o *tileMapObjectLayerObjectXML) NewActor() (any, error) {
	l := utility.NewVector(o.LocationX, o.LocationY)
	r := float64(0)
	s := utility.DefaultScale()
	sz := utility.NewVector(o.SizeX, o.SizeY)
	return actor.ActorGenerator.NewActorByName(o.Class, l, r, s, sz)
}

func (o *tileMapObjectLayerObjectXML) CreateActor() (any, error) {
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
