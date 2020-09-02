package boids

func (b *Boid) Cohesion(flock []*Boid) *Vector {
	v := &Vector{}
	for _, boid := range flock {
		if boid.ID != b.ID {
			v.Add(boid.Position)
		}
	}
	v.Div(&Vector{
		X: float64(len(flock) - 1),
		Y: float64(len(flock) - 1),
	})
	v.Sub(b.Position)
	v.Div(&Vector{
		X: 100,
		Y: 100,
	})
	return v
}

func (b *Boid) Separation(flock []*Boid) *Vector {
	v := &Vector{}
	for _, boid := range flock {
		if boid.ID != b.ID {
			if Sub(boid.Position, b.Position).Magnitude() < 100 {
				v.Sub(Sub(boid.Position, b.Position))
			}
		}
	}
	return v
}
