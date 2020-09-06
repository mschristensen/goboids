package boids

import (
	"math/rand"

	"github.com/faiface/pixel"
)

// World describes the boid universe.
type World struct {
	Width, Height int
	Boids         []*Boid
	Predators     []*Boid
}

// NewWorld instatiates a new World with randomly initialised Boids.
func NewWorld(width, height int, n int) *World {
	boids := make([]*Boid, n)
	for i := 0; i < n; i++ {
		r := rand.Float64()
		maxSpeed := 8.0
		boids[i] = &Boid{
			ID:    i,
			Alive: true,
			Position: pixel.Vec{
				X: float64(rand.Intn(width)),
				Y: float64(rand.Intn(height)),
			},
			Velocity: pixel.Vec{
				X: r * maxSpeed,
				Y: (1 - r) * maxSpeed,
			},
			Radius:             25,
			VisualRadius:       100,
			MaxSpeed:           maxSpeed,
			SeparationDistance: 80.0,
		}
	}
	predators := make([]*Boid, 1)
	r := rand.Float64()
	maxSpeed := 5.0
	predators[0] = &Boid{
		ID:    n,
		Alive: true,
		Position: pixel.Vec{
			X: float64(rand.Intn(width)),
			Y: float64(rand.Intn(height)),
		},
		Velocity: pixel.Vec{
			X: r * maxSpeed,
			Y: (1 - r) * maxSpeed,
		},
		Radius:       40,
		VisualRadius: 150,
		MaxSpeed:     5.0,
	}
	return &World{
		Width:     width,
		Height:    height,
		Boids:     boids,
		Predators: predators,
	}
}

// TickBoids updates the World's boids positions and velocities according to the rules.
func (w *World) TickBoids() {
	for i := range w.Boids {
		var vectors []pixel.Vec
		neighbours := w.Boids[i].Neighbours(w.Boids)
		vectors = append(vectors, w.Boids[i].Cohesion(neighbours, 0.001))
		vectors = append(vectors, w.Boids[i].Separation(neighbours, 0.005))
		vectors = append(vectors, w.Boids[i].Alignment(neighbours, 0.03))
		vectors = append(vectors, w.Boids[i].Bound(pixel.Vec{
			X: w.Boids[i].Radius,
			Y: w.Boids[i].Radius,
		}, pixel.Vec{
			X: float64(w.Width) - w.Boids[i].Radius,
			Y: float64(w.Height) - w.Boids[i].Radius,
		}))
		for _, predator := range w.Boids[i].Neighbours(w.Predators) {
			vectors = append(vectors, w.Boids[i].AvoidLocation(predator.Position, 0.01))
		}

		for _, vector := range vectors {
			w.Boids[i].Velocity = w.Boids[i].Velocity.Add(vector)
		}
		w.Boids[i].LimitVelocity(w.Boids[i].MaxSpeed)
		w.Boids[i].Position = w.Boids[i].Position.Add(w.Boids[i].Velocity)
	}
}

// TickPredators updates the World's predators positions and velocities according to the rules.
func (w *World) TickPredators() {
	for i := range w.Predators {
		neighbours := w.Predators[i].Neighbours(w.Boids)
		var vectors []pixel.Vec
		vectors = append(vectors, w.Predators[i].Hunt(neighbours, 0.05))
		vectors = append(vectors, w.Predators[i].Bound(pixel.Vec{
			X: 0,
			Y: 0,
		}, pixel.Vec{
			X: float64(w.Width),
			Y: float64(w.Height),
		}))

		for _, vector := range vectors {
			w.Predators[i].Velocity = w.Predators[i].Velocity.Add(vector)
		}
		w.Predators[i].LimitVelocity(w.Predators[i].MaxSpeed)
		w.Predators[i].Position = w.Predators[i].Position.Add(w.Predators[i].Velocity)
	}
}

// Tick updates the World according to the rules.
func (w *World) Tick() {
	w.TickBoids()
	w.TickPredators()
}
