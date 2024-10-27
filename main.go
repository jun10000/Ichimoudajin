package main

import (
	"github.com/jun10000/Ichimoudajin/gamemap"
	"github.com/jun10000/Ichimoudajin/utility"
)

func main() {
	g := utility.NewGame(gamemap.Stage1)
	g.WindowTitle = "Ichimoudajin"
	g.Play()
}
