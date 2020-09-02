package boids

import "math"

type Vector struct {
	X, Y float64
}

func (a *Vector) Add(b *Vector) {
	a.X += b.X
	a.Y += b.Y
}

func (a *Vector) Sub(b *Vector) {
	a.X -= b.X
	a.Y -= b.Y
}

func (a *Vector) Div(b *Vector) {
	a.X /= b.X
	a.Y /= b.Y
}

func (a *Vector) Modulo(b *Vector) {
	a.X = math.Mod(a.X, b.X)
	a.Y = math.Mod(a.Y, b.Y)
}
