package grid

import "fmt"

const (
	WALL = '#'
	FREE = '.'
)

type Pos struct {
	x, y int
}

type grid struct {
	grid       [][]rune // The actual map
	start, end Pos
}

func NewGrid(g [][]rune) *grid {
	grid := &grid{grid: g}
	grid.start = grid.GetStart()
	grid.end = grid.GetEnd()
	return grid
}

func (g *grid) Print() string {
	out := ""
	for _, v := range g.grid {
		out += fmt.Sprintln(string(v))
	}
	return out
}

func (g *grid) GetStart() Pos {
	for y, row := range g.grid {
		for x, char := range row {
			if char == 'S' {
				return Pos{x: x, y: y}
			}
		}
	}
	return Pos{}
}

func (g *grid) GetEnd() Pos {
	for y, row := range g.grid {
		for x, char := range row {
			if char == 'E' {
				return Pos{x: x, y: y}
			}
		}
	}
	return Pos{}
}

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var DIRS = []Pos{{y: -1, x: 0}, {y: 0, x: 1}, {y: 1, x: 0}, {y: 0, x: -1}}

// A node is a junction in the maze
type Node struct {
	pos  Pos     // The starting position of our node
	next []*Node // adjacent nodes
}

func NewNode(curr Pos) *Node {
	return &Node{pos: curr, next: make([]*Node, 0)}
}

// Returns the lowest path's score, the score is calculated by adding up the sum of movements and rotation
// a movement is 1 point and a rotation is 1000 points, the best path will minimise the amount of rotation
// needed to reach the goal
func (g *grid) FindPath(curr *Node, seen map[Pos]bool, startDir int, score int) []int {
	scores := make([]int, 0)
	for dir, nPos := range DIRS {
		copySeen := make(map[Pos]bool)
		for p := range seen {
			copySeen[p] = true
		}
		if g.grid[curr.pos.y+nPos.y][curr.pos.x+nPos.x] != WALL {
			next, finished, deltaScore, lastdir := g.FindNextNode(curr.pos, dir, copySeen)
			if startDir != dir { // Account for direction change on the node
				deltaScore += 1000
			}
			if finished {
				fmt.Printf("Found a path with score %d\n", score+deltaScore)
				scores = append(scores, score+deltaScore)
			} else if next != nil {
				curr.next = append(curr.next, next)
				scores = append(scores, g.FindPath(next, copySeen, lastdir, score+deltaScore)...)
			} else if next == nil {
				// This was a dead end, attempt another direction
				continue
			}
		}
	}
	return scores
}

func (g *grid) FindNextNode(start Pos, dir int, seen map[Pos]bool) (*Node, bool, int, int) {
	curr := start
	seen[curr] = true
	score := 0
	for {
		curr.y, curr.x = curr.y+DIRS[dir].y, curr.x+DIRS[dir].x
		if _, beenHereBefore := seen[curr]; beenHereBefore {
			return nil, false, -1, -1
		} else {
			seen[curr] = true
		}
		score++
		if g.IsNode(curr) {
			return NewNode(curr), false, score, dir
		} else if g.IsWall(curr) {
			// This should be a dead end
			return nil, false, -1, -1
		} else if g.IsEnd(curr) {
			return nil, true, score, -1
		} else if yes, newDir := g.IsCorner(curr, dir); yes {
			dir = newDir
			score += 1000
		}
	}
}

func (g *grid) IsEnd(pos Pos) bool {
	return g.grid[pos.y][pos.x] == 'E'
}

func (g *grid) IsWall(pos Pos) bool {
	return g.grid[pos.y][pos.x] == '#'
}

// A position is a corner if it forces changing direction
func (g *grid) IsCorner(pos Pos, dir int) (bool, int) {

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Printf("Current position y %d, x %d\n", pos.y, pos.x)
	// 	}
	// }()

	// nPos := Pos{y: pos.y + DIRS[dir].y, x: pos.x + DIRS[dir].x}
	// if !inBound(nPos, len(g.grid), len(g.grid[0])) {
	// }

	if g.grid[pos.y+DIRS[dir].y][pos.x+DIRS[dir].x] == WALL {
		if g.grid[pos.y+DIRS[(dir+1)%4].y][pos.x+DIRS[(dir+1)%4].x] == FREE {
			return true, (dir + 1) % 4
		} else if g.grid[pos.y+DIRS[(dir+3)%4].y][pos.x+DIRS[(dir+3)%4].x] == FREE {
			return true, (dir + 3) % 4
		}
	}
	return false, 0
}

// A position is a node if it has 1 or less wall around itself
func (g *grid) IsNode(pos Pos) bool {
	count := 0
	for _, nPos := range DIRS {
		if inBound(Pos{y: pos.y + nPos.y, x: pos.x + nPos.x}, len(g.grid), len(g.grid[0])) {
			if g.grid[pos.y+nPos.y][pos.x+nPos.x] == WALL {
				count++
			}
		}

	}
	return count <= 1
}

func inBound(pos Pos, maxV, maxH int) bool {
	return (pos.y < maxV && pos.y >= 0) && (pos.x < maxH && pos.x >= 0)
}
