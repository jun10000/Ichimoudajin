package tilemap

import (
	"encoding/xml"
	"log"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/actor"
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

type tileMapPropertyXML struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
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
	TiledVersion string                  `xml:"tiledversion,attr"`
	Class        string                  `xml:"class,attr"`
	Width        int                     `xml:"width,attr"`
	Height       int                     `xml:"height,attr"`
	TileWidth    int                     `xml:"tilewidth,attr"`
	TileHeight   int                     `xml:"tileheight,attr"`
	Properties   []tileMapPropertyXML    `xml:"properties>property"`
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
	IsLooping    bool
	Tilesets     []TileMapTileset
	TileLayers   []TileMapTileLayer
	ObjectLayers []TileMapObjectLayer
}

func (m *TileMap) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	// Read XML data
	data := &tileMapXML{}
	err := decoder.DecodeElement(data, &start)
	if err != nil {
		return err
	}

	// Validate map data
	if data.Version != "1.10" {
		log.Println("Loaded map version is not 1.10")
		log.Println("This may cause problem")
	}
	if data.Class != "EbitenhelperMap" {
		log.Println("Loaded map's class is not EbitenhelperMap")
	}

	// Start building TileMap
	ret := TileMap{
		MapSize:  utility.NewPoint(data.Width, data.Height),
		TileSize: utility.NewPoint(data.TileWidth, data.TileHeight),
	}

	// Apply TileMap properties
	for _, dataProperty := range data.Properties {
		switch dataProperty.Name {
		case "IsLooping":
			err := utility.StringToBool(dataProperty.Value, &ret.IsLooping)
			if err != nil {
				return err
			}
		}
	}

	// Add Tilesets
	for _, dataTileset := range data.Tilesets {
		tileset := TileMapTileset{
			ColumnCount: dataTileset.Columns,
			StartIndex:  dataTileset.FirstGID,
			LastIndex:   dataTileset.FirstGID + dataTileset.TileCount - 1,
		}

		img, err := utility.GetImageFromFile(dataTileset.Image.Source)
		if err != nil {
			return err
		}
		tileset.Image = img

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
		landscape := actor.NewActor()
		landscape.Image = ebiten.NewImage(m.MapSize.X*m.TileSize.X, m.MapSize.Y*m.TileSize.Y)
		if !yield(landscape) {
			return
		}

		collisionMap := NewTileCollisionMap(m.MapSize)

		for _, layer := range m.TileLayers {
			if layer.IsCollision {
				for cellIndex, cell := range layer.Cells {
					if cell.TileIndex < 0 {
						continue
					}

					clx := cellIndex % m.MapSize.X
					cly := cellIndex / m.MapSize.X
					collisionMap.Set(clx, cly, true)
				}
			} else {
				for cellIndex, cell := range layer.Cells {
					if cell.Tileset == nil {
						continue
					}

					tlx := cell.TileIndex % cell.Tileset.ColumnCount * m.TileSize.X
					tly := cell.TileIndex / cell.Tileset.ColumnCount * m.TileSize.Y
					clx := cellIndex % m.MapSize.X
					cly := cellIndex / m.MapSize.X
					mlx := float64(clx * m.TileSize.X)
					mly := float64(cly * m.TileSize.Y)
					img := utility.GetSubImage(cell.Tileset.Image, utility.NewPoint(tlx, tly), m.TileSize)

					o := &ebiten.DrawImageOptions{}
					o.GeoM.Translate(mlx, mly)
					landscape.Image.DrawImage(img, o)
				}
			}
		}

		for a := range collisionMap.ToBlockingAreas(m.TileSize.ToVector()) {
			if !yield(a) {
				return
			}
		}

		for _, layer := range m.ObjectLayers {
			for _, object := range layer.Objects {
				if !yield(object.Actor) {
					return
				}
			}
		}
	}
}

func GetTiledFileName(levelName string) string {
	return levelName + ".tmx"
}

func GetTileMap(levelName string) *TileMap {
	xmlData, err := assets.Assets.ReadFile(GetTiledFileName(levelName))
	utility.PanicIfError(err)

	ret := &TileMap{}
	err = xml.Unmarshal(xmlData, ret)
	utility.PanicIfError(err)

	return ret
}

func AddTileMapActorsToLevel(level *utility.Level) {
	for a := range GetTileMap(level.Name).ToActors() {
		level.Add(a)
	}
}

func NewLevelByTiledMap(levelName string) *utility.Level {
	m := GetTileMap(levelName)
	l := utility.NewLevel(levelName, m.IsLooping)
	AddTileMapActorsToLevel(l)
	return l
}
