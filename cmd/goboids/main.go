package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/mschristensen/goboids/pkg/app/goboids"
)

func main() {
	pixelgl.Run(goboids.Run)
}
