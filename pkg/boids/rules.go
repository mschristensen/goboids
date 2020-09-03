package boids

import "github.com/faiface/pixel"

func (b *Boid) Cohesion(flock []*Boid, w float64) pixel.Vec {
	v := pixel.Vec{}
	for _, boid := range flock {
		if boid.ID != b.ID {
			v = v.Add(boid.Position)
		}
	}
	return v.Scaled(1 / float64(len(flock)-1)).Sub(b.Position).Scaled(w)
}

func (b *Boid) Separation(flock []*Boid, dist, w float64) pixel.Vec {
	v := pixel.Vec{}
	for _, boid := range flock {
		if boid.ID != b.ID {
			if boid.Position.Sub(b.Position).Len() < dist {
				v = v.Sub(boid.Position.Sub(b.Position))
			}
		}
	}
	return v.Scaled(w)
}

func (b *Boid) Alignment(flock []*Boid, w float64) pixel.Vec {
	v := pixel.Vec{}
	for _, boid := range flock {
		if boid.ID != b.ID {
			v = v.Add(boid.Velocity)
		}
	}
	return v.Scaled(1 / float64(len(flock)-1)).Sub(b.Velocity).Scaled(w)
}

func (b *Boid) Bound(min, max pixel.Vec) pixel.Vec {
	v := pixel.Vec{}
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
	if b.Velocity.Len() > limit {
		b.Velocity = b.Velocity.Scaled(1 / b.Velocity.Len()).Scaled(limit)
	}
}
