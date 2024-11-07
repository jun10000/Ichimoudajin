package assets

import (
	"embed"
	"encoding/xml"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jun10000/Ichimoudajin/ebitenhelper"
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
	FirstGID int                 `xml:"firstgid,attr"`
	Image    mapTilesetImage_xml `xml:"image"`
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

type MapTileset struct {
	Image      *ebiten.Image
	StartIndex int
}

type MapLayer struct {
	Name string
	Data []string
}

type MapInfo struct {
	MapSize  ebitenhelper.Point
	TileSize ebitenhelper.Point
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
		MapSize:  ebitenhelper.NewPoint(mxml.Width, mxml.Height),
		TileSize: ebitenhelper.NewPoint(mxml.TileWidth, mxml.TileHeight),
	}

	for _, v := range mxml.Tilesets {
		image, err := GetImage(v.Image.Source)
		if err != nil {
			return err
		}

		tileset := MapTileset{
			Image:      image,
			StartIndex: v.FirstGID,
		}
		result.Tilesets = append(result.Tilesets, tileset)
	}

	for _, v := range mxml.Layers {
		ds := strings.ReplaceAll(v.Data.Inner, "\r", "")
		ds = strings.ReplaceAll(ds, "\n", "")
		ds = strings.ReplaceAll(ds, " ", "")
		layer := MapLayer{
			Name: v.Name,
			Data: strings.Split(ds, ","),
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
