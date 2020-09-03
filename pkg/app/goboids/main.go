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

const WorldWidth = 1400
const WorldHeight = 800
const MaxSpeed = 10.0
const FlockSeparation = 50
const NumBoids = 15

func Run() {
	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "GoBoids",
		Bounds: pixel.R(0, 0, WorldWidth, WorldHeight),
		VSync:  true,
	})
	if err != nil {
		panic(errors.Wrap(err, "new window failed"))
	}
	win.SetSmooth(true)
	world := boids.NewWorld(WorldWidth, WorldHeight, MaxSpeed, FlockSeparation, NumBoids)
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
