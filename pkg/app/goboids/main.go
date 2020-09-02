package goboids

import (
	_ "image/png"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mschristensen/goboids/pkg/boids"
	"github.com/mschristensen/goboids/pkg/draw"
	"github.com/pkg/errors"
)

const WorldWidth = 1024
const WorldHeight = 768
const MaxSpeedX = 10.0
const MaxSpeedY = 10.0
const NumBoids = 20

func Run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "GoBoids",
		Bounds: pixel.R(0, 0, WorldWidth, WorldHeight),
		VSync:  true,
	})
	if err != nil {
		panic(errors.Wrap(err, "new window failed"))
	}
	world := &boids.World{
		Width:     WorldWidth,
		Height:    WorldHeight,
		MaxSpeedX: MaxSpeedX,
		MaxSpeedY: MaxSpeedY,
	}
	world.Initialise(NumBoids)
	for !win.Closed() {
		err = draw.DrawFrame(win, world)
		if err != nil {
			panic(errors.Wrap(err, "draw boids failed"))
		}
		world.Tick()
		win.Update()
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}
