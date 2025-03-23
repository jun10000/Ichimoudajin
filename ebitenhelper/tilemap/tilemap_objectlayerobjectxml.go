package tilemap

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

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

func (o *tileMapObjectLayerObjectXML) NewActor() (any, error) {
	location := utility.NewVector(o.LocationX, o.LocationY)
	rotation := float64(0)
	scale := utility.DefaultScale()

	switch o.Class {
	case "Pawn":
		return actor.NewPawn(location, rotation, scale), nil
	case "AIPawn":
		return actor.NewAIPawn(location, rotation, scale), nil
	case "EnemySpawner":
		return actor.NewEnemySpawner(), nil
	default:
		return nil, fmt.Errorf("found unknown class in %s: %s", o.Name, o.Class)
	}
}

func (o *tileMapObjectLayerObjectXML) CreateActor() (any, error) {
	ret, err := o.NewActor()
	if err != nil {
		return nil, err
	}

	retv := reflect.ValueOf(ret).Elem()
	for _, property := range o.Properties {
		if m := retv.MethodByName("Set" + property.Name); m.IsValid() {
			mtype := m.Type()
			if mtype.NumIn() != 1 {
				log.Printf("Found invalid method in %s: Set%s\n", o.Name, property.Name)
				continue
			}

			switch mtype.In(0) {
			case reflect.TypeOf(bool(false)):
				v, err := strconv.ParseBool(property.Value)
				if err != nil {
					return nil, err
				}
				m.Call([]reflect.Value{reflect.ValueOf(v)})
			case reflect.TypeOf(int(0)):
				v, err := strconv.Atoi(property.Value)
				if err != nil {
					return nil, err
				}
				m.Call([]reflect.Value{reflect.ValueOf(v)})
			case reflect.TypeOf(float64(0)):
				v, err := strconv.ParseFloat(property.Value, 64)
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
				log.Printf("Found unknown argument type in Set%s: %s\n", property.Name, mtype.In(0))
			}
		} else if f := retv.FieldByName(property.Name); f.CanSet() {
			switch f.Type() {
			case reflect.TypeOf(bool(false)):
				v, err := strconv.ParseBool(property.Value)
				if err != nil {
					return nil, err
				}
				f.Set(reflect.ValueOf(v))
			case reflect.TypeOf(int(0)):
				v, err := strconv.Atoi(property.Value)
				if err != nil {
					return nil, err
				}
				f.Set(reflect.ValueOf(v))
			case reflect.TypeOf(float64(0)):
				v, err := strconv.ParseFloat(property.Value, 64)
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
				log.Printf("Found unknown field type in %s: %s\n", property.Name, f.Type())
			}
		} else {
			log.Printf("Found unknown property in %s: %s\n", o.Name, property.Name)
		}
	}

	return ret, nil
}
