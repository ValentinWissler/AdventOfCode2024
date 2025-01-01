package grid

import "fmt"

const (
	WALL    = '#'
	FREE    = '.'
	PINK    = "\033[35m"
	GREEN   = "\033[32m"
	DEFAULT = "\033[0m"
)

type Pos struct {
	x, y int
}

type grid struct {
	grid       [][]rune // The actual map
	start, end Pos
	tree       *NodeTree
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

// Return all the nodes on the grid plus the start and the end position
func (g *grid) FindNodes() map[Pos]*Node {
	nodes := make(map[Pos]*Node, 0)
	start := NewNode(g.GetStart())
	nodes[start.pos] = start
	for y, row := range g.grid {
		for x, item := range row {
			if item == WALL {
				continue
			}
			pos := Pos{x: x, y: y}
			if g.IsNode(pos) {
				curr := NewNode(pos)
				nodes[curr.pos] = curr
			}
		}
	}
	end := NewNode(g.GetEnd())
	nodes[end.pos] = end
	return nodes
}

// Visualise the nodes on the grid
func (g *grid) PrintNodes(nodes map[Pos]*Node) string {
	out := ""
	for y, row := range g.grid {
		for x, item := range row {
			pos := Pos{x: x, y: y}
			if _, isNode := nodes[pos]; isNode {
				out += fmt.Sprintf("%s%s%s", PINK, "+", DEFAULT)
			} else {
				out += fmt.Sprint(string(item))
			}
			if x+1 == len(row) {
				out += "\n"
			}
		}
	}
	return out
}

// Connect each nodes to its neghboring nodes, each node has at least 1 neighbour
// this allows us building a weighted graph
func (g *grid) ConnectNodes(nodes map[Pos]*Node) *NodeTree {
	// Keep track of the Nodes we processed
	// seenNodes := make(map[*Node]bool)
	// For each known position of a Node
	for pos, node := range nodes {
		for dir := range DIRS {
			next, _, distance, _ := g.FindNextNode(pos, dir)
			if next != nil {
				node.addEdge(NewNodeEdge(distance, next))
				next.addEdge(NewNodeEdge(distance, node))
			}
		}
		edges := make(map[Pos]bool)
		for _, edge := range node.edge {
			edges[edge.next.pos] = true
		}
		fmt.Println(g.PrintConnectedNodes(pos, edges))
	}
	root := nodes[g.start]
	g.tree = NewNodeTree(root)
	return g.tree
}

func (g *grid) PrintConnectedNodes(source Pos, destination map[Pos]bool) string {
	out := ""
	for y, row := range g.grid {
		for x, item := range row {
			pos := Pos{x: x, y: y}
			if pos.x == source.x && pos.y == source.y {
				out += fmt.Sprintf("%s%s%s", GREEN, "+", DEFAULT)
			} else if _, isDest := destination[pos]; isDest {
				out += fmt.Sprintf("%s%s%s", PINK, "o", DEFAULT)
			} else {
				out += string(item)
			}
			if x+1 == len(row) {
				out += "\n"
			}
		}
	}
	return out
}

func (g *grid) FindNextNode(start Pos, dir int) (*Node, bool, int, int) {
	curr := start
	seen := make(map[Pos]bool)
	seen[curr] = true
	score := 0
	for {
		curr.y, curr.x = curr.y+DIRS[dir].y, curr.x+DIRS[dir].x
		// fmt.Printf("exploring %d,%d\n", curr.y, curr.x)
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

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var DIRS = []Pos{{y: -1, x: 0}, {y: 0, x: 1}, {y: 1, x: 0}, {y: 0, x: -1}}

// Returns the lowest path's score, the score is calculated by adding up the sum of movements and rotation
// a movement is 1 point and a rotation is 1000 points, the best path will minimise the amount of rotation
// needed to reach the goal
// func (g *grid) FindPath(curr *Node, seen map[Pos]bool, startDir int, score int) []int {
// 	scores := make([]int, 0)
// 	for dir, nPos := range DIRS {
// 		copySeen := make(map[Pos]bool)
// 		for p := range seen {
// 			copySeen[p] = true
// 		}
// 		if g.grid[curr.pos.y+nPos.y][curr.pos.x+nPos.x] != WALL {
// 			next, finished, deltaScore, lastdir := g.FindNextNode(curr.pos, dir, copySeen)
// 			if startDir != dir { // Account for direction change on the node
// 				deltaScore += 1000
// 			}
// 			if finished {
// 				fmt.Printf("Found a path with score %d\n", score+deltaScore)
// 				scores = append(scores, score+deltaScore)
// 			} else if next != nil {
// 				curr.next = append(curr.next, next)
// 				scores = append(scores, g.FindPath(next, copySeen, lastdir, score+deltaScore)...)
// 			} else if next == nil {
// 				// This was a dead end, attempt another direction
// 				continue
// 			}
// 		}
// 	}
// 	return scores
// }

func (g *grid) IsEnd(pos Pos) bool {
	return g.grid[pos.y][pos.x] == 'E'
}

func (g *grid) IsWall(pos Pos) bool {
	return g.grid[pos.y][pos.x] == '#'
}

// A position is a corner if it forces changing direction
func (g *grid) IsCorner(pos Pos, dir int) (bool, int) {
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
