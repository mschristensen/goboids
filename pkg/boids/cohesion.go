package boids

func (b *Boid) Cohesion(flock []*Boid) *Position {
	centre := &Position{}
	for _, boid := range flock {
		if boid.id != b.id {
			centre.Add(boid.position)
		}
	}
	centre.Div(&Position{
		x: len(flock) - 1,
		y: len(flock) - 1,
	})
	centre.Sub(b.position)
	centre.Div(&Position{
		x: 100,
		y: 100,
	})
	return centre
}
