package boids

func (b *Boid) Cohesion(flock []*Boid) *Vector {
	centre := &Vector{}
	for _, boid := range flock {
		if boid.ID != b.ID {
			centre.Add(boid.Position)
		}
	}
	centre.Div(&Vector{
		X: float64(len(flock) - 1),
		Y: float64(len(flock) - 1),
	})
	centre.Sub(b.Position)
	centre.Div(&Vector{
		X: 100,
		Y: 100,
	})
	return centre
}
