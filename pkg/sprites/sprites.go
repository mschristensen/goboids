package sprites

import (
	"github.com/faiface/pixel"
	"github.com/gobuffalo/packr/v2"
	"github.com/pkg/errors"
)

var box *packr.Box

// Assets returns the packr box providing access to the
// static asset files which are packaged in the built binary.
func Assets() *packr.Box {
	if box == nil {
		box = packr.New("assets", "./assets")
	}
	return box
}

// Strip describes a strip of sprites.
type Strip struct {
	Asset        string
	Width        int
	Height       int
	SpriteWidth  int
	SpriteHeight int
}

// Sprite describes a single sprite in a strip.
type Sprite struct {
	X, Y int
}

// Bounds returns the bounds of the sprite at the given position in the sprite given
// sprite strip. The coordinate (0, 0) corresponds to the bottom left sprite.
func (strip *Strip) Bounds(sprite *Sprite) (*pixel.Rect, error) {
	if sprite.X > strip.Width {
		return nil, errors.Errorf("x cannot be > 6, got %d", sprite.X)
	}
	if sprite.X < 0 {
		return nil, errors.Errorf("x cannot be < 0, got %d", sprite.X)
	}
	if sprite.Y > strip.Height {
		return nil, errors.Errorf("y cannot be > 4, got %d", sprite.Y)
	}
	if sprite.Y < 0 {
		return nil, errors.Errorf("y cannot be < 0, got %d", sprite.Y)
	}
	bounds := pixel.R(
		float64(sprite.X)*96,
		float64(sprite.Y)*96,
		(float64(sprite.X)+1)*96,
		(float64(sprite.Y)+1)*96,
	)
	return &bounds, nil
}
