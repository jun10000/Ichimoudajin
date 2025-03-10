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

type TileMap struct {
	MapSize  utility.Point
	TileSize utility.Point
	Tilesets []TileMapTileset
	Layers   []TileMapTileLayer
}

func (m *TileMap) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	// Read and check XML data
	mxml := &tileMapXML{}
	err := decoder.DecodeElement(mxml, &start)
	if err != nil {
		return err
	}
	if mxml.Version != "1.10" {
		log.Println("Loaded map version is not 1.10")
		log.Println("This may cause problem")
	}

	// Begin creating MapInfo
	result := TileMap{
		MapSize:  utility.NewPoint(mxml.Width, mxml.Height),
		TileSize: utility.NewPoint(mxml.TileWidth, mxml.TileHeight),
	}

	// Add MapTilesets
	for _, v := range mxml.Tilesets {
		image := utility.GetImageFile(v.Image.Source)
		tileset := TileMapTileset{
			Image:       image,
			ColumnCount: v.Columns,
			StartIndex:  v.FirstGID,
			LastIndex:   v.FirstGID + v.TileCount - 1,
		}
		result.Tilesets = append(result.Tilesets, tileset)
	}

	// Add MapLayers
	for _, v := range mxml.TileLayers {
		layer := TileMapTileLayer{
			Name:        v.Name,
			IsCollision: (v.Name == "Collision"),
		}

		// Add Properties
		// for _, p := range v.Properties {
		// }

		// Add MapCells
		cellstrings := strings.ReplaceAll(v.Data.Inner, "\r", "")
		cellstrings = strings.ReplaceAll(cellstrings, "\n", "")
		cellstrings = strings.ReplaceAll(cellstrings, " ", "")
		for _, cellstring := range strings.Split(cellstrings, ",") {
			c := TileMapTileLayerCell{
				TileIndex: -1,
			}
			cellvalue, _ := strconv.Atoi(cellstring)
			for _, t := range result.Tilesets {
				if t.StartIndex <= cellvalue && cellvalue <= t.LastIndex {
					c.Tileset = &t
					c.TileIndex = cellvalue - t.StartIndex
					break
				}
			}

			layer.Cells = append(layer.Cells, c)
		}

		result.Layers = append(result.Layers, layer)
	}

	// Finished creating MapInfo
	*m = result
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

		for _, l := range m.Layers {
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
