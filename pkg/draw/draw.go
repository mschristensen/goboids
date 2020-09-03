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
		pos := pixel.Vec{
			X: boid.Position.X,
			Y: boid.Position.Y,
		}
		vel := pixel.Vec{
			X: boid.Velocity.X,
			Y: boid.Velocity.Y,
		}
		theta := vel.Angle()
		if theta < 0 {
			theta += 2 * math.Pi
		}
		sprite.Draw(
			window,
			pixel.IM.Moved(pos).Scaled(pos, 0.4).Rotated(
				pos,
				// gopher's head is upright so offset by -90deg to align head with x axis
				theta-(math.Pi/2),
			))
	}
	return nil
}
