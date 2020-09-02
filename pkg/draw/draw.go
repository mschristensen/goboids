package draw

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mschristensen/goboids/pkg/boids"
	"github.com/mschristensen/goboids/pkg/sprites"
	"github.com/pkg/errors"
)

// DrawBoids draws the given list of boids on the window.
func DrawBoids(window *pixelgl.Window, boids []*boids.Boid) error {
	sprite, err := sprites.NewGopher("normal")
	if err != nil {
		return errors.Wrap(err, "new gopher failed")
	}
	for _, boid := range boids {
		sprite.Draw(window, pixel.IM.Moved(pixel.Vec{
			X: boid.Position.X,
			Y: boid.Position.Y,
		}))
	}
	return nil
}
