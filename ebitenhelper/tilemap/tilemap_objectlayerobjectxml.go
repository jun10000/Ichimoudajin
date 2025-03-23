package tilemap

import (
	"log"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
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

func (o *tileMapObjectLayerObjectXML) CreateActor() (any, error) {
	var ret any
	switch o.Class {
	case "Pawn":
		l := utility.NewVector(o.LocationX, o.LocationY)
		ret = actor.NewPawn(l, 0, utility.DefaultScale())
	case "AIPawn":
		l := utility.NewVector(o.LocationX, o.LocationY)
		ret = actor.NewAIPawn(l, 0, utility.DefaultScale())
	default:
		log.Println("Found unsupported Tiled map object class: " + o.Class)
		return nil, nil
	}

	retv := reflect.ValueOf(ret).Elem()
	for _, property := range o.Properties {
		if m := retv.MethodByName("Set" + property.Name); m.IsValid() {
			mtype := m.Type()
			if mtype.NumIn() != 1 {
				log.Printf("Set%s method has invalid argument counts\n", property.Name)
				continue
			}

			switch mtype.In(0) {
			case reflect.TypeOf(float64(0)):
				var v float64
				err := utility.StringToFloat(property.Value, &v)
				if err != nil {
					return nil, err
				}
				m.Call([]reflect.Value{reflect.ValueOf(v)})
			case reflect.TypeOf((*ebiten.Image)(nil)):
				img, err := utility.GetImageFromFile(property.Value)
				if err != nil {
					return nil, err
				}
				m.Call([]reflect.Value{reflect.ValueOf(img)})
			default:
				log.Printf("Found unsupported argument type %s\n", mtype.In(0))
			}
		} else if f := retv.FieldByName(property.Name); f.CanSet() {
			switch f.Type() {
			case reflect.TypeOf(float64(0)):
				var v float64
				err := utility.StringToFloat(property.Value, &v)
				if err != nil {
					return nil, err
				}
				f.Set(reflect.ValueOf(v))
			case reflect.TypeOf((*ebiten.Image)(nil)):
				img, err := utility.GetImageFromFile(property.Value)
				if err != nil {
					return nil, err
				}
				f.Set(reflect.ValueOf(img))
			default:
				log.Printf("Found unsupported field type %s\n", f.Type())
			}
		} else {
			log.Printf("Found unknown property (%s) in %s\n", property.Name, o.Name)
		}
	}

	// Calculate scale
	// sz := utility.NewVector(o.SizeX, o.SizeY)
	// s := sz.Div(ret.FrameSize.ToVector())
	// ret.SetScale(s)

	return ret, nil
}
