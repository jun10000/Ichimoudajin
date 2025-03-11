package actor

import (
	"encoding/xml"
	"log"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

type tileMapObjectLayerObjectPropertyXML struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type tileMapTilesetImageXML struct {
	Source string `xml:"source,attr"`
}

type tileMapTileLayerPropertyXML struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type tileMapTileLayerDataXML struct {
	Inner string `xml:",innerxml"`
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

type tileMapTilesetXML struct {
	FirstGID  int                    `xml:"firstgid,attr"`
	TileCount int                    `xml:"tilecount,attr"`
	Columns   int                    `xml:"columns,attr"`
	Image     tileMapTilesetImageXML `xml:"image"`
}

type tileMapTileLayerXML struct {
	Name       string                        `xml:"name,attr"`
	Properties []tileMapTileLayerPropertyXML `xml:"properties>property"`
	Data       tileMapTileLayerDataXML       `xml:"data"`
}

type tileMapObjectLayerXML struct {
	Name    string                        `xml:"name,attr"`
	Objects []tileMapObjectLayerObjectXML `xml:"object"`
}

type tileMapXML struct {
	Version      string                  `xml:"version,attr"`
	Width        int                     `xml:"width,attr"`
	Height       int                     `xml:"height,attr"`
	TileWidth    int                     `xml:"tilewidth,attr"`
	TileHeight   int                     `xml:"tileheight,attr"`
	Tilesets     []tileMapTilesetXML     `xml:"tileset"`
	TileLayers   []tileMapTileLayerXML   `xml:"layer"`
	ObjectLayers []tileMapObjectLayerXML `xml:"objectgroup"`
}

type TileMapTileLayerCell struct {
	Tileset   *TileMapTileset
	TileIndex int
}

type TileMapObjectLayerObject struct {
	Name  string
	Actor any
}

type TileMapTileset struct {
	Image       *ebiten.Image
	ColumnCount int
	StartIndex  int
	LastIndex   int
}

type TileMapTileLayer struct {
	Name        string
	IsCollision bool
	Cells       []TileMapTileLayerCell
}

type TileMapObjectLayer struct {
	Name    string
	Objects []TileMapObjectLayerObject
}

type TileMap struct {
	MapSize      utility.Point
	TileSize     utility.Point
	Tilesets     []TileMapTileset
	TileLayers   []TileMapTileLayer
	ObjectLayers []TileMapObjectLayer
}

func (o *tileMapObjectLayerObjectXML) CreateActor() (any, error) {
	switch o.Class {
	case "Pawn":
		ret := NewPawn()
		ret.SetLocation(utility.NewVector(o.LocationX, o.LocationY))
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
				ret.Image = utility.GetImageFile(property.Value)
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
			}
		}

		return ret, nil
	case "AIPawn":
		ret := NewAIPawn()
		ret.SetLocation(utility.NewVector(o.LocationX, o.LocationY))
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
				ret.Image = utility.GetImageFile(property.Value)
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
			}
		}

		return ret, nil
	}

	log.Println("Found unsupported Tiled map object class: " + o.Class)
	return nil, nil
}

func (m *TileMap) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	// Read XML data
	data := &tileMapXML{}
	err := decoder.DecodeElement(data, &start)
	if err != nil {
		return err
	}

	// Check map version
	if data.Version != "1.10" {
		log.Println("Loaded map version is not 1.10")
		log.Println("This may cause problem")
	}

	// Start building TileMap
	ret := TileMap{
		MapSize:  utility.NewPoint(data.Width, data.Height),
		TileSize: utility.NewPoint(data.TileWidth, data.TileHeight),
	}

	// Add Tilesets
	for _, dataTileset := range data.Tilesets {
		tileset := TileMapTileset{
			Image:       utility.GetImageFile(dataTileset.Image.Source),
			ColumnCount: dataTileset.Columns,
			StartIndex:  dataTileset.FirstGID,
			LastIndex:   dataTileset.FirstGID + dataTileset.TileCount - 1,
		}
		ret.Tilesets = append(ret.Tilesets, tileset)
	}

	// Add TileLayers
	for _, dataTileLayer := range data.TileLayers {
		tileLayer := TileMapTileLayer{
			Name:        dataTileLayer.Name,
			IsCollision: (dataTileLayer.Name == "Collision"),
		}

		// Add TileLayer.Properties
		// for _, dataTileLayerProperty := range dataTileLayer.Properties {
		// }

		// Add TileLayer.Cells
		css := utility.RemoveAllStrings(dataTileLayer.Data.Inner, "\r", "\n", " ")
		for _, cs := range strings.Split(css, ",") {
			cv, err := strconv.Atoi(cs)
			if err != nil {
				return err
			}

			tileLayerCell := TileMapTileLayerCell{
				TileIndex: -1,
			}
			for _, tileset := range ret.Tilesets {
				if tileset.StartIndex <= cv && cv <= tileset.LastIndex {
					tileLayerCell.Tileset = &tileset
					tileLayerCell.TileIndex = cv - tileset.StartIndex
					break
				}
			}

			tileLayer.Cells = append(tileLayer.Cells, tileLayerCell)
		}

		ret.TileLayers = append(ret.TileLayers, tileLayer)
	}

	// Add ObjectLayers
	for _, dataObjectLayer := range data.ObjectLayers {
		objectLayer := TileMapObjectLayer{
			Name: dataObjectLayer.Name,
		}

		// Add ObjectLayer.Objects
		for _, dataObjectLayerObject := range dataObjectLayer.Objects {
			objectLayerObject := TileMapObjectLayerObject{
				Name: dataObjectLayerObject.Name,
			}
			a, err := dataObjectLayerObject.CreateActor()
			if err != nil {
				return err
			}

			if a != nil {
				objectLayerObject.Actor = a
				objectLayer.Objects = append(objectLayer.Objects, objectLayerObject)
			}
		}

		ret.ObjectLayers = append(ret.ObjectLayers, objectLayer)
	}

	// Finish building TileMap
	*m = ret
	return nil
}

func (m *TileMap) ToActors() func(yield func(any) bool) {
	return func(yield func(any) bool) {
		mapimage := ebiten.NewImage(m.MapSize.X*m.TileSize.X, m.MapSize.Y*m.TileSize.Y)
		mapactor := NewActor()
		mapactor.Image = mapimage
		if !yield(mapactor) {
			return
		}

		for _, l := range m.TileLayers {
			if l.IsCollision {
				for ci, c := range l.Cells {
					if c.TileIndex < 0 {
						continue
					}

					b := NewBlockingArea()
					b.SetLocation(utility.NewVector(
						float64((ci%m.MapSize.X)*m.TileSize.X),
						float64(ci/m.MapSize.X*m.TileSize.Y)))
					b.Size = m.TileSize.ToVector()
					if !yield(b) {
						return
					}
				}
			} else {
				for ci, c := range l.Cells {
					if c.Tileset == nil {
						continue
					}

					o := &ebiten.DrawImageOptions{}
					o.GeoM.Translate(
						float64((ci%m.MapSize.X)*m.TileSize.X),
						float64(ci/m.MapSize.X*m.TileSize.Y))
					mapimage.DrawImage(utility.GetSubImage(
						c.Tileset.Image,
						utility.NewPoint(
							c.TileIndex%c.Tileset.ColumnCount*m.TileSize.X,
							c.TileIndex/c.Tileset.ColumnCount*m.TileSize.Y),
						m.TileSize), o)
				}
			}
		}

		for _, ol := range m.ObjectLayers {
			for _, o := range ol.Objects {
				if !yield(o.Actor) {
					return
				}
			}
		}
	}
}

func GetTileMap(filename string) *TileMap {
	xmlData, err := assets.Assets.ReadFile(filename)
	utility.PanicIfError(err)

	ret := &TileMap{}
	err = xml.Unmarshal(xmlData, ret)
	utility.PanicIfError(err)

	return ret
}

func AddTileMapActorsToLevel(level *utility.Level, filename string) {
	for a := range GetTileMap(filename).ToActors() {
		level.Add(a)
	}
}
