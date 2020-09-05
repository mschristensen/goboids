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
	Asset        pixel.Picture
	Width        float64
	Height       float64
	SpriteWidth  float64
	SpriteHeight float64
}

// Bounds returns the bounds of the sprite at the given position in the sprite given
// sprite strip. The coordinate (0, 0) corresponds to the bottom left sprite.
func (strip *Strip) Bounds(loc pixel.Vec) (*pixel.Rect, error) {
	if loc.X > strip.Width {
		return nil, errors.Errorf("x cannot be > 6, got %d", loc.X)
	}
	if loc.X < 0 {
		return nil, errors.Errorf("x cannot be < 0, got %d", loc.X)
	}
	if loc.Y > strip.Height {
		return nil, errors.Errorf("y cannot be > 4, got %d", loc.Y)
	}
	if loc.Y < 0 {
		return nil, errors.Errorf("y cannot be < 0, got %d", loc.Y)
	}
	bounds := pixel.R(
		float64(loc.X)*float64(strip.SpriteWidth),
		float64(loc.Y)*float64(strip.SpriteHeight),
		(float64(loc.X)+1)*float64(strip.SpriteWidth),
		(float64(loc.Y)+1)*float64(strip.SpriteHeight),
	)
	return &bounds, nil
}

// NewSprite returns a Sprite from the given location in the Strip.
func (strip *Strip) NewSprite(loc pixel.Vec) (*pixel.Sprite, error) {
	bounds, err := strip.Bounds(loc)
	if err != nil {
		return nil, errors.Wrap(err, "get bounds failed")
	}
	return pixel.NewSprite(strip.Asset, *bounds), nil
}
