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

// Drawer enables batch drawing sprites from a given strip.
type Drawer struct {
	Batch       *pixel.Batch
	SpriteStrip *sprites.Strip
}

// NewDrawer creates a new Drawer for the given sprite strip.
func NewDrawer(strip *sprites.Strip) *Drawer {
	return &Drawer{
		Batch:       pixel.NewBatch(&pixel.TrianglesData{}, strip.Asset),
		SpriteStrip: strip,
	}
}

// DrawFrame draws the world on the window.
func (d *Drawer) DrawFrame(window *pixelgl.Window, world *boids.World) error {
	window.Clear(colornames.Aliceblue)
	d.Batch.Clear()
	sprite, err := d.SpriteStrip.NewSprite(sprites.Gophers["normal"])
	if err != nil {
		return errors.Wrap(err, "new gopher failed")
	}
	for _, boid := range world.Boids {
		theta := boid.Velocity.Angle()
		if theta < 0 {
			theta += 2 * math.Pi
		}
		imd := imdraw.New(nil)

		// Render visual radius of boid
		imd.Color = color.RGBA{
			colornames.Blue.R,
			colornames.Blue.G,
			colornames.Blue.B,
			0x33,
		}
		imd.Push(boid.Position)
		imd.Circle(boid.VisualRadius, 0)
		imd.Draw(d.Batch)

		// Render boid itself as a Gopher
		sprite.Draw(
			d.Batch,
			pixel.IM.Moved(boid.Position).Scaled(
				boid.Position,
				// size boid according to radius, using the original size of sprite
				boid.Radius/float64(d.SpriteStrip.SpriteWidth),
			).Rotated(
				boid.Position,
				// gopher's head is upright so offset by -90deg to align head with x axis
				theta-(math.Pi/2),
			),
		)
	}
	d.Batch.Draw(window)
	return nil
}
