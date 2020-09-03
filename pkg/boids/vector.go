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

func (a *Vector) Mul(b *Vector) {
	a.X *= b.X
	a.Y *= b.Y
}

func (a *Vector) Div(b *Vector) {
	a.X /= b.X
	a.Y /= b.Y
}

func (a *Vector) Modulo(b *Vector) {
	a.X = math.Mod(a.X, b.X)
	a.Y = math.Mod(a.Y, b.Y)
}

func (a *Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(a.X, 2) + math.Pow(a.Y, 2))
}

func Add(a, b *Vector) *Vector {
	return &Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func Sub(a, b *Vector) *Vector {
	return &Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func Mul(a, b *Vector) *Vector {
	return &Vector{
		X: a.X * b.X,
		Y: a.Y * b.Y,
	}
}

func Div(a, b *Vector) *Vector {
	return &Vector{
		X: a.X / b.X,
		Y: a.Y / b.Y,
	}
}

func Modulo(a, b *Vector) *Vector {
	return &Vector{
		X: math.Mod(a.X, b.X),
		Y: math.Mod(a.Y, b.Y),
	}
}
