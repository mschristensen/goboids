package sprites

import (
	"image"

	"strings"

	"github.com/faiface/pixel"
	"github.com/pkg/errors"
)

// Gophers maps Gopher names to positions in the sprite strip.
var Gophers = map[string]pixel.Vec{
	"normal": {
		X: 0,
		Y: 4,
	},
}

// GophersStrip returns the sprite strip for the gophers.
func GophersStrip() (*Strip, error) {
	assetName := "gophers.png"
	imgData, err := Assets().FindString(assetName)
	if err != nil {
		return nil, errors.Wrapf(err, "find asset '%s' failed", assetName)
	}
	img, _, err := image.Decode(strings.NewReader(imgData))
	if err != nil {
		return nil, errors.Wrap(err, "decode image failed")
	}
	return &Strip{
		Asset:        pixel.PictureDataFromImage(img),
		Width:        672,
		Height:       480,
		SpriteWidth:  96,
		SpriteHeight: 96,
	}, nil
}
