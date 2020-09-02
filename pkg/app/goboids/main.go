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
const MaxSpeedX = 1.0
const MaxSpeedY = 1.0
const NumBoids = 10

func Run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "GoBoids",
		Bounds: pixel.R(0, 0, WorldWidth, WorldHeight),
		VSync:  true,
	})
	if err != nil {
		panic(errors.Wrap(err, "new window failed"))
	}
	world := boids.NewWorld(
		WorldWidth, WorldHeight,
		MaxSpeedX, MaxSpeedY,
		NumBoids,
	)
	throttle := time.Tick(time.Millisecond * 10)
	for !win.Closed() {
		err = draw.DrawFrame(win, world)
		if err != nil {
			panic(errors.Wrap(err, "draw boids failed"))
		}
		world.Tick()
		win.Update()
		<-throttle
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}
