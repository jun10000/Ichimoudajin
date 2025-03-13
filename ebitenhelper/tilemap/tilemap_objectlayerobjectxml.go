package tilemap

import (
	"log"

	"github.com/jun10000/Ichimoudajin/ebitenhelper/actor"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type tileMapObjectLayerObjectXML struct {
	Name       string                                `xml:"name,attr"`
	Class      string                                `xml:"type,attr"`
	LocationX  float64                               `xml:"x,attr"`
	LocationY  float64                               `xml:"y,attr"`
	SizeX      float64                               `xml:"width,attr"`
	SizeY      float64                               `xml:"height,attr"`
	Properties []tileMapObjectLayerObjectPropertyXML `xml:"properties>property"`
}

func (o *tileMapObjectLayerObjectXML) CreatePawn() (*actor.Pawn, error) {
	rettf := utility.NewTransform(
		utility.NewVector(o.LocationX, o.LocationY),
		0,
		utility.NewVector(1, 1))
	ret := actor.NewPawn(rettf)
	ret.FrameSize.X = int(o.SizeX)
	ret.FrameSize.Y = int(o.SizeY)

	for _, property := range o.Properties {
		switch property.Name {
		case "Accel":
			err := utility.StringToFloat(property.Value, &ret.Accel)
			if err != nil {
				return nil, err
			}
		case "Decel":
			err := utility.StringToFloat(property.Value, &ret.Decel)
			if err != nil {
				return nil, err
			}
		case "FPS":
			err := utility.StringToInt(property.Value, &ret.FPS)
			if err != nil {
				return nil, err
			}
		case "FrameCount":
			err := utility.StringToInt(property.Value, &ret.FrameCount)
			if err != nil {
				return nil, err
			}
		case "FrameDirectionMap":
			clear(ret.FrameDirectionMap)
			for _, v := range property.Value {
				ret.FrameDirectionMap = append(ret.FrameDirectionMap, utility.RuneToInt(v))
			}
		case "Image":
			img, err := utility.GetImageFromFile(property.Value)
			if err != nil {
				return nil, err
			}
			ret.Image = img
		case "MaxSpeed":
			err := utility.StringToFloat(property.Value, &ret.MaxSpeed)
			if err != nil {
				return nil, err
			}
		case "RotationDeg":
			var deg float64
			err := utility.StringToFloat(property.Value, &deg)
			if err != nil {
				return nil, err
			}
			ret.SetRotation(utility.DegreeToRadian(deg))
		case "ScaleX":
			s := ret.GetScale()
			err := utility.StringToFloat(property.Value, &s.X)
			if err != nil {
				return nil, err
			}
			ret.SetScale(s)
		case "ScaleY":
			s := ret.GetScale()
			err := utility.StringToFloat(property.Value, &s.Y)
			if err != nil {
				return nil, err
			}
			ret.SetScale(s)
		default:
			log.Printf("Found unknown Tiled object (%s) property: %s = %s\n",
				o.Name, property.Name, property.Value)
		}
	}

	return ret, nil
}

func (o *tileMapObjectLayerObjectXML) CreateAIPawn() (*actor.AIPawn, error) {
	rettf := utility.NewTransform(
		utility.NewVector(o.LocationX, o.LocationY),
		0,
		utility.NewVector(1, 1))
	ret := actor.NewAIPawn(rettf)
	ret.FrameSize.X = int(o.SizeX)
	ret.FrameSize.Y = int(o.SizeY)

	for _, property := range o.Properties {
		switch property.Name {
		case "Accel":
			err := utility.StringToFloat(property.Value, &ret.Accel)
			if err != nil {
				return nil, err
			}
		case "Decel":
			err := utility.StringToFloat(property.Value, &ret.Decel)
			if err != nil {
				return nil, err
			}
		case "FPS":
			err := utility.StringToInt(property.Value, &ret.FPS)
			if err != nil {
				return nil, err
			}
		case "FrameCount":
			err := utility.StringToInt(property.Value, &ret.FrameCount)
			if err != nil {
				return nil, err
			}
		case "FrameDirectionMap":
			clear(ret.FrameDirectionMap)
			for _, v := range property.Value {
				ret.FrameDirectionMap = append(ret.FrameDirectionMap, utility.RuneToInt(v))
			}
		case "Image":
			img, err := utility.GetImageFromFile(property.Value)
			if err != nil {
				return nil, err
			}
			ret.Image = img
		case "MaxSpeed":
			err := utility.StringToFloat(property.Value, &ret.MaxSpeed)
			if err != nil {
				return nil, err
			}
		case "RotationDeg":
			var deg float64
			err := utility.StringToFloat(property.Value, &deg)
			if err != nil {
				return nil, err
			}
			ret.SetRotation(utility.DegreeToRadian(deg))
		case "ScaleX":
			s := ret.GetScale()
			err := utility.StringToFloat(property.Value, &s.X)
			if err != nil {
				return nil, err
			}
			ret.SetScale(s)
		case "ScaleY":
			s := ret.GetScale()
			err := utility.StringToFloat(property.Value, &s.Y)
			if err != nil {
				return nil, err
			}
			ret.SetScale(s)
		default:
			log.Printf("Found unknown Tiled object (%s) property: %s = %s\n",
				o.Name, property.Name, property.Value)
		}
	}

	return ret, nil
}

func (o *tileMapObjectLayerObjectXML) CreateActor() (any, error) {
	switch o.Class {
	case "Pawn":
		return o.CreatePawn()
	case "AIPawn":
		return o.CreateAIPawn()
	default:
		log.Println("Found unsupported Tiled map object class: " + o.Class)
		return nil, nil
	}
}
