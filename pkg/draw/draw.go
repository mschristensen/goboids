package draw

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mschristensen/goboids/pkg/boids"
	"github.com/mschristensen/goboids/pkg/sprites"
	"github.com/pkg/errors"
	"golang.org/x/image/colornames"
)

// DrawFrame draws the given world on the window.
func DrawFrame(window *pixelgl.Window, world *boids.World) error {
	window.Clear(colornames.Aliceblue)
	sprite, err := sprites.NewGopher("normal")
	if err != nil {
		return errors.Wrap(err, "new gopher failed")
	}
	for _, boid := range world.Boids {
		// TODO probably draws from top left, so position and scale should consider this
		sprite.Draw(window, pixel.IM.Moved(pixel.Vec{
			X: boid.Position.X,
			Y: boid.Position.Y,
		}).Scaled(pixel.Vec{
			X: boid.Position.X,
			Y: boid.Position.Y,
		}, 0.5))
	}
	return nil
}
