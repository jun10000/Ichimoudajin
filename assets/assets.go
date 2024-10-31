package assets

import (
	"embed"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed *
var assets embed.FS

func GetImage(imagefile string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFileSystem(assets, "images/"+imagefile)
	if err != nil {
		log.Fatal(err)
	}

	return image
}
