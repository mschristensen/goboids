package boids

type Velocity struct {
	x, y float64
}

func (a *Velocity) Add(b *Velocity) {
	a.x += b.x
	a.y += b.y
}

type Position struct {
	x, y int
}

func (a *Position) Add(b *Position) {
	a.x += b.x
	a.y += b.y
}

func (a *Position) Sub(b *Position) {
	a.x -= b.x
	a.y -= b.y
}

func (a *Position) Div(b *Position) {
	a.x /= b.x
	a.y /= b.y
}
