package boids

import (
	"math"

	"github.com/faiface/pixel"
)

func (b *Boid) Cohesion(flock []*Boid, w float64) pixel.Vec {
	v := pixel.Vec{}
	if len(flock) == 0 {
		return v
	}
	for _, boid := range flock {
		if boid.ID != b.ID {
			v = v.Add(boid.Position)
		}
	}
	n := float64(len(flock) - 1)
	if n <= 0 {
		n = 1
	}
	return v.Scaled(1 / n).Sub(b.Position).Scaled(w)
}

func (b *Boid) Separation(flock []*Boid, w float64) pixel.Vec {
	v := pixel.Vec{}
	if len(flock) == 0 {
		return v
	}
	for _, boid := range flock {
		if boid.ID != b.ID {
			if boid.Position.Sub(b.Position).Len() < b.SeparationDistance {
				v = v.Sub(boid.Position.Sub(b.Position))
			}
		}
	}
	return v.Scaled(w)
}

func (b *Boid) Alignment(flock []*Boid, w float64) pixel.Vec {
	v := pixel.Vec{}
	if len(flock) == 0 {
		return v
	}
	for _, boid := range flock {
		if boid.ID != b.ID {
			v = v.Add(boid.Velocity)
		}
	}
	n := float64(len(flock) - 1)
	if n <= 0 {
		n = 1
	}
	return v.Scaled(1 / n).Sub(b.Velocity).Scaled(w)
}

// Hunt returns a vector towards the nearest boid in the flock.
func (b *Boid) Hunt(flock []*Boid, w float64) pixel.Vec {
	v := pixel.Vec{}
	if len(flock) == 0 {
		return v
	}
	var target *Boid
	minDistance := math.MaxFloat64
	for i, boid := range flock {
		if boid.ID != b.ID {
			distance := math.Sqrt(
				math.Pow(boid.Position.X-b.Position.X, 2) - math.Pow(boid.Position.Y-b.Position.Y, 2),
			)
			if distance < minDistance {
				minDistance = distance
				target = flock[i]
			}
		}
	}
	if target == nil {
		return v
	}
	if minDistance < target.Radius {
		target.Alive = false
	}
	return b.TendToLocation(target.Position, w)
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

func (b *Boid) TendToLocation(vec pixel.Vec, w float64) pixel.Vec {
	return vec.Sub(b.Position).Scaled(w)
}

func (b *Boid) AvoidLocation(vec pixel.Vec, w float64) pixel.Vec {
	return b.TendToLocation(vec, -w)
}
