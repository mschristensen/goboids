package boids

import (
	"math/rand"

	"github.com/faiface/pixel"
)

// World describes the boid universe.
type World struct {
	Width, Height   int
	MaxSpeed        float64
	FlockSeparation float64
	Boids           []*Boid
}

// NewWorld instatiates a new World with randomly initialised Boids.
func NewWorld(width, height int, maxSpeed, flockSeparation float64, n int) *World {
	boids := make([]*Boid, n)
	for i := 0; i < n; i++ {
		r := rand.Float64()
		boids[i] = &Boid{
			ID: i,
			Position: pixel.Vec{
				X: float64(rand.Intn(width)),
				Y: float64(rand.Intn(height)),
			},
			Velocity: pixel.Vec{
				X: r * maxSpeed,
				Y: (1 - r) * maxSpeed,
			},
			Radius:       50,
			VisualRadius: 100,
		}
	}
	return &World{
		Width:           width,
		Height:          height,
		MaxSpeed:        maxSpeed,
		FlockSeparation: flockSeparation,
		Boids:           boids,
	}
}

// TODO add rules to avoid objects & predators - include walls in this in addition to Bounds
// TODO add rules to tend towards food
// TODO add rules for perching
// TODO add rules for scattering
// TODO limit acceleration
// TODO batch draw for efficiency
func (w *World) Tick() {
	for i := range w.Boids {
		neighbours := w.Boids[i].Neighbours(w.Boids)
		v1 := w.Boids[i].Cohesion(neighbours, 0.001)
		v2 := w.Boids[i].Separation(neighbours, w.FlockSeparation, 0.005)
		v3 := w.Boids[i].Alignment(neighbours, 0.01)
		v4 := w.Boids[i].Bound(pixel.Vec{
			X: 0,
			Y: 0,
		}, pixel.Vec{
			X: float64(w.Width),
			Y: float64(w.Height),
		})

		w.Boids[i].Velocity = w.Boids[i].Velocity.Add(v1)
		w.Boids[i].Velocity = w.Boids[i].Velocity.Add(v2)
		w.Boids[i].Velocity = w.Boids[i].Velocity.Add(v3)
		w.Boids[i].Velocity = w.Boids[i].Velocity.Add(v4)
		w.Boids[i].LimitVelocity(w.MaxSpeed)
		w.Boids[i].Position = w.Boids[i].Position.Add(w.Boids[i].Velocity)
	}
}
