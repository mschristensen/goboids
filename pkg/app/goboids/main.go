package goboids

import (
	"fmt"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mschristensen/goboids/pkg/boids"
	"github.com/mschristensen/goboids/pkg/draw"
	"github.com/mschristensen/goboids/pkg/sprites"
	"github.com/pkg/errors"
)

const (
	WorldWidth      = 1400
	WorldHeight     = 800
	MaxSpeed        = 5.0
	FlockSeparation = 80
	NumBoids        = 15
)

func createWindow(title string, width, height float64) (*pixelgl.Window, error) {
	window, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "new window failed")
	}
	window.SetSmooth(true)
	return window, nil
}

func Run() {
	window, err := createWindow("GoBoids", WorldWidth, WorldHeight)
	if err != nil {
		panic(errors.Wrap(err, "create window failed"))
	}
	strip, err := sprites.GophersStrip()
	if err != nil {
		panic(errors.Wrap(err, "get gophers strip failed"))
	}
	drawer := draw.NewDrawer(strip)
	world := boids.NewWorld(WorldWidth, WorldHeight, MaxSpeed, FlockSeparation, NumBoids)
	frames := 0
	second := time.Tick(time.Second)
	for !window.Closed() {
		err = drawer.DrawFrame(window, world)
		if err != nil {
			panic(errors.Wrap(err, "draw boids failed"))
		}
		world.Tick()
		window.Update()
		frames++
		select {
		case <-second:
			window.SetTitle(fmt.Sprintf("%s | FPS: %d", "GoBoids", frames))
			frames = 0
		default:
		}
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}
