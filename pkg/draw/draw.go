package draw

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
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
		imd := imdraw.New(nil)
		imd.Color = color.RGBA{
			colornames.Blue.R,
			colornames.Blue.G,
			colornames.Blue.B,
			0x33,
		}
		imd.Push(boid.Position)
		imd.Circle(boid.VisualRadius, 0)
		imd.Draw(window)
		sprite.Draw(
			window,
			pixel.IM.Moved(boid.Position).Scaled(
				boid.Position,
				// size boid according to radius, using the original size of sprite
				boid.Radius/float64(sprites.GophersStrip.SpriteWidth),
			).Rotated(
				boid.Position,
				// gopher's head is upright so offset by -90deg to align head with x axis
				theta-(math.Pi/2),
			),
		)
	}
	return nil
}
