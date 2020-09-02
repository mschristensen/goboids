package boids

type Boid struct {
	ID          int
	Position    *Vector
	Velocity    *Vector
	VisualRange int
}

// Neighbours returns the boids in the flock that are within the visual range of b.
func (b *Boid) Neighbours(flock []*Boid) []*Boid {
	return nil
}
