package boids

func (b *Boid) Cohesion(flock []*Boid, w *Vector) *Vector {
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
	v.Mul(w)
	return v
}

func (b *Boid) Separation(flock []*Boid, w *Vector) *Vector {
	v := &Vector{}
	for _, boid := range flock {
		if boid.ID != b.ID {
			if Sub(boid.Position, b.Position).Magnitude() < 100 {
				v.Sub(Sub(boid.Position, b.Position))
			}
		}
	}
	v.Mul(w)
	return v
}

func (b *Boid) Alignment(flock []*Boid, w *Vector) *Vector {
	v := &Vector{}
	for _, boid := range flock {
		if boid.ID != b.ID {
			v.Add(boid.Velocity)
		}
	}
	v.Div(&Vector{
		X: float64(len(flock) - 1),
		Y: float64(len(flock) - 1),
	})
	v.Sub(b.Velocity)
	v.Mul(w)
	return v
}

func (b *Boid) Bound(min, max *Vector) *Vector {
	v := &Vector{}
	if b.Position.X < min.X {
		v.X = 10
	} else if b.Position.X > max.X {
		v.X = -10
	}
	if b.Position.Y < min.Y {
		v.Y = 10
	} else if b.Position.Y > max.Y {
		v.Y = -10
	}
	return v
}

func (b *Boid) LimitVelocity(limit float64) {
	if b.Velocity.Magnitude() > limit {
		b.Velocity.Div(&Vector{
			X: b.Velocity.Magnitude(),
			Y: b.Velocity.Magnitude(),
		})
		b.Velocity.Mul(&Vector{
			X: limit,
			Y: limit,
		})
	}
}
