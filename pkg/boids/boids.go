package boids

import (
	"math"

	"github.com/faiface/pixel"
)

// Boid describes a single boid.
type Boid struct {
	ID           int
	Position     pixel.Vec
	Velocity     pixel.Vec
	Radius       float64
	VisualRadius float64
}

// Neighbours returns the boids in the flock that are within the field of view of b.
func (b *Boid) Neighbours(flock []*Boid) []*Boid {
	var neighbours []*Boid
	for i := range flock {
		dist := math.Sqrt(
			math.Pow(flock[i].Position.X-b.Position.X, 2) + math.Pow(flock[i].Position.Y-b.Position.Y, 2),
		)
		if dist <= b.VisualRadius+flock[i].Radius {
			neighbours = append(neighbours, flock[i])
		}
	}
	return neighbours
}
