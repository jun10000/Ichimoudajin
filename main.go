package main

import (
	"log"

	"github.com/jun10000/Ichimoudajin/assets"
	"github.com/jun10000/Ichimoudajin/ebitenhelper"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/actor"
	"github.com/jun10000/Ichimoudajin/ebitenhelper/utility"
)

func NewStage1() *utility.Level {
	mapdata, err := assets.GetMapData("stage1.tmx")
	if err != nil {
		log.Fatal(err)
	}

	image_player, err := assets.GetImage("images/ぴぽやキャラチップ32出力素材/現代系/女_スーツ1.png")
	if err != nil {
		log.Fatal(err)
	}

	level := utility.NewLevel()

	for _, l := range mapdata.Layers {
		for ci, c := range l.Cells {
			if c.Tileset == nil {
				continue
			}

			a := actor.NewActor()
			a.Location = utility.NewVector(
				float64((ci%mapdata.MapSize.X)*mapdata.TileSize.X),
				float64(ci/mapdata.MapSize.X*mapdata.TileSize.Y))
			a.Image.Source = utility.GetSubImage(c.Tileset.Image,
				utility.NewPoint(
					c.TileIndex%c.Tileset.ColumnCount*mapdata.TileSize.X,
					c.TileIndex/c.Tileset.ColumnCount*mapdata.TileSize.Y),
				mapdata.TileSize)
			level.Add(a)
		}
	}

	player := actor.NewPawn()
	player.Location = utility.NewVector(600, 300)
	player.Animation.Source = image_player
	level.Add(player)

	return level
}

func main() {
	g := ebitenhelper.NewGame()
	g.WindowTitle = "Ichimoudajin"
	g.ScreenWidth = 32 * 40
	g.ScreenHeight = 32 * 22
	err := g.Play(NewStage1())
	if err != nil {
		log.Fatal(err)
	}
}
