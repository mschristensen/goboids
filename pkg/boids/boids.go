package boids

type Boid struct {
	id          int
	position    *Position
	velocity    *Velocity
	visualRange int
}

// Neighbours returns the boids in the flock that are within the visual range of b.
func (b *Boid) Neighbours(flock []*Boid) []*Boid {
	return nil
}
