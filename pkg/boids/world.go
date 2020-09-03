package boids

import "math/rand"

// World describes the boid universe.
type World struct {
	Width, Height int
	MaxSpeed      float64
	Boids         []*Boid
}

// NewWorld instatiates a new World with randomly initialised Boids.
func NewWorld(width, height int, maxSpeed float64, n int) *World {
	boids := make([]*Boid, n)
	for i := 0; i < n; i++ {
		r := rand.Float64()
		boids[i] = &Boid{
			ID: i,
			Position: &Vector{
				X: float64(rand.Intn(width)),
				Y: float64(rand.Intn(height)),
			},
			Velocity: &Vector{
				X: r * maxSpeed,
				Y: (1 - r) * maxSpeed,
			},
			VisualRange: 100,
		}
	}
	return &World{
		Width:    width,
		Height:   height,
		MaxSpeed: maxSpeed,
		Boids:    boids,
	}
}

// TODO add rules to tend towards food
// TODO add rules to avoid predators
// TODO add rules for perching
// TODO add rules for scattering
func (w *World) Tick() {
	for _, boid := range w.Boids {
		v1 := boid.Cohesion(w.Boids, &Vector{
			X: 0.02,
			Y: 0.02,
		})
		v2 := boid.Separation(w.Boids, &Vector{
			X: 0.05,
			Y: 0.05,
		})
		v3 := boid.Alignment(w.Boids, &Vector{
			X: 0.05,
			Y: 0.05,
		})
		v4 := boid.Bound(&Vector{
			X: 0,
			Y: 0,
		}, &Vector{
			X: float64(w.Width),
			Y: float64(w.Height),
		})

		boid.Velocity.Add(v1)
		boid.Velocity.Add(v2)
		boid.Velocity.Add(v3)
		boid.Velocity.Add(v4)
		boid.LimitVelocity(w.MaxSpeed)
		boid.Position.Add(boid.Velocity)
	}
}
