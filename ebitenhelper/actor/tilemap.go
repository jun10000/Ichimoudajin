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

type mapTilesetImageXML struct {
	Source string `xml:"source,attr"`
}

type mapLayerPropertyXML struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type mapLayerDataXML struct {
	Inner string `xml:",innerxml"`
}

type mapTilesetXML struct {
	FirstGID  int                `xml:"firstgid,attr"`
	TileCount int                `xml:"tilecount,attr"`
	Columns   int                `xml:"columns,attr"`
	Image     mapTilesetImageXML `xml:"image"`
}

type mapLayerXML struct {
	Name       string                `xml:"name,attr"`
	Properties []mapLayerPropertyXML `xml:"properties>property"`
	Data       mapLayerDataXML       `xml:"data"`
}

type mapInfoXML struct {
	Version    string          `xml:"version,attr"`
	Width      int             `xml:"width,attr"`
	Height     int             `xml:"height,attr"`
	TileWidth  int             `xml:"tilewidth,attr"`
	TileHeight int             `xml:"tileheight,attr"`
	Tilesets   []mapTilesetXML `xml:"tileset"`
	Layers     []mapLayerXML   `xml:"layer"`
}

type MapCell struct {
	Tileset   *MapTileset
	TileIndex int
}

type MapTileset struct {
	Image       *ebiten.Image
	ColumnCount int
	StartIndex  int
	LastIndex   int
}

type MapLayer struct {
	Name        string
	IsCollision bool
	Cells       []MapCell
}

type MapInfo struct {
	MapSize  utility.Point
	TileSize utility.Point
	Tilesets []MapTileset
	Layers   []MapLayer
}

func (m *MapInfo) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	// Read and check XML data
	mxml := &mapInfoXML{}
	err := decoder.DecodeElement(mxml, &start)
	if err != nil {
		return err
	}
	if mxml.Version != "1.10" {
		log.Println("Loaded map version is not 1.10")
		log.Println("This may cause problem")
	}

	// Begin creating MapInfo
	result := MapInfo{
		MapSize:  utility.NewPoint(mxml.Width, mxml.Height),
		TileSize: utility.NewPoint(mxml.TileWidth, mxml.TileHeight),
	}

	// Add MapTilesets
	for _, v := range mxml.Tilesets {
		image := utility.GetImageFile(v.Image.Source)
		tileset := MapTileset{
			Image:       image,
			ColumnCount: v.Columns,
			StartIndex:  v.FirstGID,
			LastIndex:   v.FirstGID + v.TileCount - 1,
		}
		result.Tilesets = append(result.Tilesets, tileset)
	}

	// Add MapLayers
	for _, v := range mxml.Layers {
		layer := MapLayer{
			Name:        v.Name,
			IsCollision: false,
		}

		// Add Properties
		for _, p := range v.Properties {
			switch p.Name {
			case "IsCollision":
				v, _ := strconv.ParseBool(p.Value)
				layer.IsCollision = v
			}
		}

		// Add MapCells
		cellstrings := strings.ReplaceAll(v.Data.Inner, "\r", "")
		cellstrings = strings.ReplaceAll(cellstrings, "\n", "")
		cellstrings = strings.ReplaceAll(cellstrings, " ", "")
		for _, cellstring := range strings.Split(cellstrings, ",") {
			c := MapCell{
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

func (m *MapInfo) GetActors() func(yield func(any) bool) {
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

func getMapInfo(filename string) (*MapInfo, error) {
	data, err := assets.Assets.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data2 := &MapInfo{}
	err = xml.Unmarshal(data, data2)
	if err != nil {
		return nil, err
	}

	return data2, nil
}

func getActorsFromMapFile(filename string) (func(yield func(any) bool), error) {
	mi, err := getMapInfo(filename)
	if err != nil {
		return nil, err
	}
	return mi.GetActors(), nil
}

func AddActorsToLevelFromMapFile(level *utility.Level, filename string) {
	as, err := getActorsFromMapFile(filename)
	utility.PanicIfError(err)
	for a := range as {
		level.Add(a)
	}
}
