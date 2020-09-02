package boids

import "math/rand"

// World describes the boid universe.
type World struct {
	Width, Height        int
	MaxSpeedX, MaxSpeedY float64
	Boids                []*Boid
}

// Initialise sets randomly initialised Boids on the given World.
func (w *World) Initialise(n int) {
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
	w.Boids = boids
}

func (w *World) Tick() {
	for i := range w.Boids {
		v1 := w.Boids[i].Cohesion(w.Boids)
		// v2
		// v3

		w.Boids[i].Velocity.Add(v1)
		w.Boids[i].Position.Add(w.Boids[i].Velocity)
	}
}
