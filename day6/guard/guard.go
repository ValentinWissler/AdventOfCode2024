package guard

import "fmt"

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

// Directions the guard can move towards
var (
	DIR_UP    = [2]int{-1, 0}
	DIR_DOWN  = [2]int{1, 0}
	DIR_LEFT  = [2]int{0, -1}
	DIR_RIGHT = [2]int{0, 1}
)

// A guard has a position on the grid and moves in a specific direction
type Guard struct {
	dir int
	pos [2]int
}

func NewGuard(dir int, pos [2]int) *Guard {
	return &Guard{dir: dir, pos: pos}
}

// A guard can move in a specific direction
func (g *Guard) Move() {
	switch g.dir {
	case UP:
		g.pos[0], g.pos[1] = g.pos[0]+DIR_UP[0], g.pos[1]+DIR_UP[1]
	case DOWN:
		g.pos[0], g.pos[1] = g.pos[0]+DIR_DOWN[0], g.pos[1]+DIR_DOWN[1]
	case LEFT:
		g.pos[0], g.pos[1] = g.pos[0]+DIR_LEFT[0], g.pos[1]+DIR_LEFT[1]
	case RIGHT:
		g.pos[0], g.pos[1] = g.pos[0]+DIR_RIGHT[0], g.pos[1]+DIR_RIGHT[1]
	default:
		panic(fmt.Sprintf("unknown directions: %d", g.dir))
	}
}

// A guard can change direction if facing an obstacle
func (g *Guard) ChangeDir() {
	g.dir = (g.dir + 1) % 4
}

// Return the next dir should the guard change direction
func (g *Guard) NextDir() int {
	return (g.dir + 1) % 4
}

// Get Guard position
func (g *Guard) Pos() [2]int {
	return g.pos
}

// Get Guard theoratical next pos
func (g *Guard) NextPos(dir int) [2]int {
	switch dir {
	case UP:
		return [2]int{g.pos[0] + DIR_UP[0], g.pos[1] + DIR_UP[1]}
	case DOWN:
		return [2]int{g.pos[0] + DIR_DOWN[0], g.pos[1] + DIR_DOWN[1]}
	case LEFT:
		return [2]int{g.pos[0] + DIR_LEFT[0], g.pos[1] + DIR_LEFT[1]}
	case RIGHT:
		return [2]int{g.pos[0] + DIR_RIGHT[0], g.pos[1] + DIR_RIGHT[1]}
	default:
		panic(fmt.Sprintf("unknown directions: %d", dir))
	}
}

// Get Guard direction
func (g *Guard) Dir() int {
	return g.dir
}
