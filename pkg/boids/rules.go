package boids

import (
	"math"
	"math/rand"

	"github.com/faiface/pixel"
)

// Cohesion steers the Boid to travel towards the center of the flock.
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

// Separation steers the the Boid to maintain the minimum separation distances with
// other Boids in the flock.
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

// Alignment steers the Boid to align it's velocity direction with the
// other Boids in the flock.
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

// Bound keeps the Boid within the given bounds by modifying its velocity.
func (b *Boid) Bound(min, max pixel.Vec) pixel.Vec {
	v := pixel.Vec{}
	if b.Position.X < min.X {
		v.X = rand.Float64() * 10.0
	} else if b.Position.X > max.X {
		v.X = -rand.Float64() * 10.0
	}
	if b.Position.Y < min.Y {
		v.Y = rand.Float64() * 10.0
	} else if b.Position.Y > max.Y {
		v.Y = -rand.Float64() * 10.0
	}
	return v
}

// LimitVelocity ensures the magnitude of the Boid's velocity is below the given limit.
// If not, it is set to a vector of equivalent direction but of `limit` magnitude.
func (b *Boid) LimitVelocity(limit float64) {
	if b.Velocity.Len() > limit {
		b.Velocity = b.Velocity.Scaled(1 / b.Velocity.Len()).Scaled(limit)
	}
}

// TendToLocation returns a vector pointing towards the given location, weighted by w.
func (b *Boid) TendToLocation(vec pixel.Vec, w float64) pixel.Vec {
	return vec.Sub(b.Position).Scaled(w)
}

// AvoidLocation returns a vector pointing away from the given location, weighted by w.
func (b *Boid) AvoidLocation(vec pixel.Vec, w float64) pixel.Vec {
	return b.TendToLocation(vec, -w)
}
