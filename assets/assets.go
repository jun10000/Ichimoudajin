package assets

import (
	"embed"
	"encoding/xml"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

//go:embed *
var assets embed.FS

type mapTilesetImage_xml struct {
	Source string `xml:"source,attr"`
}

type mapLayerData_xml struct {
	Inner string `xml:",innerxml"`
}

type mapTileset_xml struct {
	FirstGID  int                 `xml:"firstgid,attr"`
	TileCount int                 `xml:"tilecount,attr"`
	Columns   int                 `xml:"columns,attr"`
	Image     mapTilesetImage_xml `xml:"image"`
}

type mapLayer_xml struct {
	Name string           `xml:"name,attr"`
	Data mapLayerData_xml `xml:"data"`
}

type mapInfo_xml struct {
	Version    string           `xml:"version,attr"`
	Width      int              `xml:"width,attr"`
	Height     int              `xml:"height,attr"`
	TileWidth  int              `xml:"tilewidth,attr"`
	TileHeight int              `xml:"tileheight,attr"`
	Tilesets   []mapTileset_xml `xml:"tileset"`
	Layers     []mapLayer_xml   `xml:"layer"`
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
	Name  string
	Cells []MapCell
}

type MapInfo struct {
	MapSize  utility.Point
	TileSize utility.Point
	Tilesets []MapTileset
	Layers   []MapLayer
}

func (m *MapInfo) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	mxml := &mapInfo_xml{}
	err := decoder.DecodeElement(mxml, &start)
	if err != nil {
		return err
	}
	if mxml.Version != "1.10" {
		log.Println("Loaded map version is not 1.10")
		log.Println("This may cause problem")
	}

	result := MapInfo{
		MapSize:  utility.NewPoint(mxml.Width, mxml.Height),
		TileSize: utility.NewPoint(mxml.TileWidth, mxml.TileHeight),
	}

	for _, v := range mxml.Tilesets {
		image, err := GetImage(v.Image.Source)
		if err != nil {
			return err
		}

		tileset := MapTileset{
			Image:       image,
			ColumnCount: v.Columns,
			StartIndex:  v.FirstGID,
			LastIndex:   v.FirstGID + v.TileCount - 1,
		}
		result.Tilesets = append(result.Tilesets, tileset)
	}

	for _, v := range mxml.Layers {
		layer := MapLayer{
			Name: v.Name,
		}

		cellstrings := strings.ReplaceAll(v.Data.Inner, "\r", "")
		cellstrings = strings.ReplaceAll(cellstrings, "\n", "")
		cellstrings = strings.ReplaceAll(cellstrings, " ", "")

		for _, cellstring := range strings.Split(cellstrings, ",") {
			c := MapCell{}
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

	*m = result
	return nil
}

func GetMapData(mapfile string) (*MapInfo, error) {
	data, err := assets.ReadFile(mapfile)
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

func GetImage(imagefile string) (*ebiten.Image, error) {
	image, _, err := ebitenutil.NewImageFromFileSystem(assets, imagefile)
	if err != nil {
		return nil, err
	}

	return image, nil
}
