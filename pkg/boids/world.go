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
	Predators       []*Boid
}

// NewWorld instatiates a new World with randomly initialised Boids.
func NewWorld(width, height int, maxSpeed, flockSeparation float64, n int) *World {
	boids := make([]*Boid, n)
	for i := 0; i < n; i++ {
		r := rand.Float64()
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
			Radius:       50,
			VisualRadius: 100,
		}
	}
	predators := make([]*Boid, 1)
	r := rand.Float64()
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
		Radius:       80,
		VisualRadius: 130,
	}
	return &World{
		Width:           width,
		Height:          height,
		MaxSpeed:        maxSpeed,
		FlockSeparation: flockSeparation,
		Boids:           boids,
		Predators:       predators,
	}
}

// TODO add rules to avoid objects
// TODO add rules to avoid flying into walls
// TODO add rules to tend towards food
// TODO add rules for perching
// TODO add rules for scattering
// TODO limit acceleration
func (w *World) TickBoids() {
	for i := range w.Boids {
		var vectors []pixel.Vec
		neighbours := w.Boids[i].Neighbours(w.Boids)
		vectors = append(vectors, w.Boids[i].Cohesion(neighbours, 0.001))
		vectors = append(vectors, w.Boids[i].Separation(neighbours, w.FlockSeparation, 0.005))
		vectors = append(vectors, w.Boids[i].Alignment(neighbours, 0.03))
		vectors = append(vectors, w.Boids[i].Bound(pixel.Vec{
			X: 0,
			Y: 0,
		}, pixel.Vec{
			X: float64(w.Width),
			Y: float64(w.Height),
		}))
		for _, predator := range w.Boids[i].Neighbours(w.Predators) {
			vectors = append(vectors, w.Boids[i].AvoidLocation(predator.Position, 0.01))
		}

		for _, vector := range vectors {
			w.Boids[i].Velocity = w.Boids[i].Velocity.Add(vector)
		}
		w.Boids[i].LimitVelocity(w.MaxSpeed)
		w.Boids[i].Position = w.Boids[i].Position.Add(w.Boids[i].Velocity)
	}
}

func (w *World) TickPredators() {
	for i := range w.Predators {
		neighbours := w.Predators[i].Neighbours(w.Boids)
		v1 := w.Predators[i].Hunt(neighbours, 0.05)
		v2 := w.Predators[i].Bound(pixel.Vec{
			X: 0,
			Y: 0,
		}, pixel.Vec{
			X: float64(w.Width),
			Y: float64(w.Height),
		})

		w.Predators[i].Velocity = w.Predators[i].Velocity.Add(v1)
		w.Predators[i].Velocity = w.Predators[i].Velocity.Add(v2)
		w.Predators[i].LimitVelocity(w.MaxSpeed)
		w.Predators[i].Position = w.Predators[i].Position.Add(w.Predators[i].Velocity)
	}
}

func (w *World) Tick() {
	w.TickBoids()
	w.TickPredators()
}
