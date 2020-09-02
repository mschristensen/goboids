package main

import (
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mschristensen/goboids/pkg/sprites"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "GoBoids",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	sprite, err := sprites.NewGopher("normal")
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Skyblue)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
