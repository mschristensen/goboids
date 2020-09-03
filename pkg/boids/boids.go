package boids

import "github.com/faiface/pixel"

// Boid describes a single boid.
type Boid struct {
	ID          int
	Position    pixel.Vec
	Velocity    pixel.Vec
	VisualRange int
}

// Neighbours returns the boids in the flock that are within the visual range of b.
func (b *Boid) Neighbours(flock []*Boid) []*Boid {
	return nil
}
