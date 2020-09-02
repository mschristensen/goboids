package sprites

import (
	"image"

	"strings"

	"github.com/faiface/pixel"
	"github.com/pkg/errors"
)

// GophersStrip is the sprite Strip for the gophers.
var GophersStrip = &Strip{
	Asset:        "gophers.png",
	Width:        672,
	Height:       480,
	SpriteWidth:  96,
	SpriteHeight: 96,
}

// Gophers maps Gopher names to Sprites.
var Gophers = map[string]*Sprite{
	"normal": {
		X: 0,
		Y: 4,
	},
}

// NewGopher returns a new pixel Sprite for the gopher with the specified name.
func NewGopher(name string) (*pixel.Sprite, error) {
	imgData, err := Assets().FindString(GophersStrip.Asset)
	if err != nil {
		return nil, errors.Wrapf(err, "find asset '%s' failed", GophersStrip.Asset)
	}
	img, _, err := image.Decode(strings.NewReader(imgData))
	if err != nil {
		return nil, errors.Wrap(err, "decode image failed")
	}
	gopher, ok := Gophers[name]
	if !ok {
		return nil, errors.Wrapf(err, "gopher with name '%s' not found", name)
	}
	bounds, err := GophersStrip.Bounds(gopher)
	if err != nil {
		return nil, errors.Wrap(err, "get bounds failed")
	}
	return pixel.NewSprite(pixel.PictureDataFromImage(img), *bounds), nil
}
