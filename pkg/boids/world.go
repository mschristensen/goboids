package boids

import "math/rand"

// World describes the boid universe.
type World struct {
	Width, Height        int
	MaxSpeedX, MaxSpeedY float64
}

// Initialise returns a set of Boids randomly initialised in the given World.
func (w *World) Initialise(n int) []*Boid {
	boids := make([]*Boid, n)
	for i := 0; i < n; i++ {
		boids[i] = &Boid{
			ID: i,
			Position: &Vector{
				X: float64(rand.Intn(w.Width)),
				Y: float64(rand.Intn(w.Height)),
			},
			Velocity: &Vector{
				X: rand.Float64() * w.MaxSpeedX,
				Y: rand.Float64() * w.MaxSpeedY,
			},
			VisualRange: 100,
		}
	}
	return boids
}
