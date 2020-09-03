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
			VisualRange: 100,
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

// TODO add neighbours, ideally with a forward facing vision with parameterised breadth
// TODO add boid size & consider overlap in neighbours calc
// TODO add rules to tend towards food
// TODO add rules to avoid predators
// TODO add rules for perching
// TODO add rules for scattering
// TODO limit acceleration
// TODO visualise visual range etc
func (w *World) Tick() {
	for _, boid := range w.Boids {
		v1 := boid.Cohesion(w.Boids, 0.01)
		v2 := boid.Separation(w.Boids, w.FlockSeparation, 0.1)
		v3 := boid.Alignment(w.Boids, 0.01)
		v4 := boid.Bound(pixel.Vec{
			X: 0,
			Y: 0,
		}, pixel.Vec{
			X: float64(w.Width),
			Y: float64(w.Height),
		})

		boid.Velocity = boid.Velocity.Add(v1)
		boid.Velocity = boid.Velocity.Add(v2)
		boid.Velocity = boid.Velocity.Add(v3)
		boid.Velocity = boid.Velocity.Add(v4)
		boid.LimitVelocity(w.MaxSpeed)
		boid.Position = boid.Position.Add(boid.Velocity)
	}
}
