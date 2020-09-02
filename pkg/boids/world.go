package boids

import "math/rand"

// World describes the boid universe.
type World struct {
	Width, Height        int
	MaxSpeedX, MaxSpeedY float64
	Boids                []*Boid
}

// NewWorld instatiates a new World with randomly initialised Boids.
func NewWorld(width, height int, maxSpeedX, maxSpeedY float64, n int) *World {
	boids := make([]*Boid, n)
	for i := 0; i < n; i++ {
		boids[i] = &Boid{
			ID: i,
			Position: &Vector{
				X: float64(rand.Intn(width)),
				Y: float64(rand.Intn(height)),
			},
			Velocity: &Vector{
				X: rand.Float64() * maxSpeedX,
				Y: rand.Float64() * maxSpeedY,
			},
			VisualRange: 100,
		}
	}
	return &World{
		Width:     width,
		Height:    height,
		MaxSpeedX: maxSpeedX,
		MaxSpeedY: maxSpeedY,
		Boids:     boids,
	}
}

func (w *World) Tick() {
	for i := range w.Boids {
		v1 := w.Boids[i].Cohesion(w.Boids)
		// v2
		// v3

		w.Boids[i].Velocity.Add(v1)
		w.Boids[i].Position.Add(w.Boids[i].Velocity)
		w.Boids[i].Position.Modulo(&Vector{
			X: float64(w.Width),
			Y: float64(w.Height),
		})
	}
}
