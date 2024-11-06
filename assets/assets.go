package assets

import (
	"embed"
	"encoding/xml"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed *
var assets embed.FS

type MapTilesetImageData struct {
	Source string `xml:"source,attr"`
}

type MapTilesetData struct {
	StartGID int                 `xml:"firstgid,attr"`
	Images   MapTilesetImageData `xml:"image"`
}

type MapLayerInnerData struct {
	Data []byte `xml:",innerxml"`
}

type MapLayerData struct {
	Name      string              `xml:"name,attr"`
	DataArray []MapLayerInnerData `xml:"data"`
}

type MapData struct {
	Version    string           `xml:"version,attr"`
	Width      int              `xml:"width,attr"`
	Height     int              `xml:"height,attr"`
	TileWidth  int              `xml:"tilewidth,attr"`
	TileHeight int              `xml:"tileheight,attr"`
	Tilesets   []MapTilesetData `xml:"tileset"`
	Layers     []MapLayerData   `xml:"layer"`
}

func GetImage(imagefile string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFileSystem(assets, "images/"+imagefile)
	if err != nil {
		log.Fatal(err)
	}

	return image
}

func GetMapData(mapfile string) *MapData {
	data, err := assets.ReadFile(mapfile)
	if err != nil {
		log.Fatal(err)
	}

	data2 := MapData{}
	err = xml.Unmarshal(data, &data2)
	if err != nil {
		log.Fatal(err)
	}

	return &data2
}
