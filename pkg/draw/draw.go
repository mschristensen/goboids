package draw

import (
	"math"

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
		theta := boid.Velocity.Angle()
		if theta < 0 {
			theta += 2 * math.Pi
		}
		sprite.Draw(
			window,
			pixel.IM.Moved(boid.Position).Scaled(boid.Position, 0.4).Rotated(
				boid.Position,
				// gopher's head is upright so offset by -90deg to align head with x axis
				theta-(math.Pi/2),
			))
	}
	return nil
}
