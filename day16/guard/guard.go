package guard

type Pos struct {
	X, Y int
}

type guard struct {
	// guard have a position on the grid
	pos        Pos
	vel        Pos
	maxX, maxY int
}

func Newguard(pos, vel Pos, mx, my int) *guard {
	return &guard{pos: pos, vel: vel, maxX: mx, maxY: my}
}

func (g *guard) Move() {
	np := Pos{X: g.pos.X + g.vel.X, Y: g.pos.Y + g.vel.Y}
	// Adjust if oob
	if np.X >= g.maxX {
		np.X -= g.maxX
	}
	if np.Y >= g.maxY {
		np.Y -= g.maxY
	}
	g.pos.X, g.pos.Y = np.X, np.Y
}

func (g *guard) Pos() Pos {
	return g.pos
}
