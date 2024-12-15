package grid

import (
	g "aoc6/guard"
	"aoc6/utils"
	"fmt"
)

const (
	FREE     = '.'
	SEEN_H   = '-' // Tile visited while walking horizontally
	SEEN_V   = '|' // Tile visited while walking vertically
	SEEN_C   = '+' // Tile is a crossing
	OBSTACLE = '#'
	VORTEX   = '0'
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

// USE the const UP DOWN LEFT RIGHT to access DIR value
var DIR = []rune{'^', 'v', '<', '>'}

// A grid is bi-dimensional map containing various obstacles and a moving guard
type Grid struct {
	// Current state of the grid
	grid [][]rune
	// The guard in the grid
	guard *g.Guard
}

func NewGrid(input [][]rune) *Grid {
	grid := &Grid{grid: input}
	pos, dir := findGuard(input)
	if dir == -1 {
		panic(fmt.Sprintf("couldn't find guard on the grid\n%s", grid.PrintMaze()))
	}
	grid.guard = g.NewGuard(dir, pos)
	return grid
}

// Print the grid state, for debugging
func (g *Grid) PrintMaze() string {
	str := ""
	for x, row := range g.grid {
		if x+1 == len(g.grid) {
			str += fmt.Sprint(string(row))
			continue
		}
		str += fmt.Sprintln(string(row))
	}
	return str
}

// Count the number of tiles visited by the guard
func (g *Grid) CountVisitedTiles() int {
	c := 0
	for _, row := range g.grid {
		for _, char := range row {
			if char == '|' || char == '-' || char == '+' || char == '0' {
				c++
			}
		}
	}
	return c
}

// Finds the guard position and direction in the grid
func findGuard(grid [][]rune) ([2]int, int) {
	for ri, row := range grid {
		for ci, item := range row {
			switch item {
			case '^':
				return [2]int{ri, ci}, UP
			case '<':
				return [2]int{ri, ci}, LEFT
			case '>':
				return [2]int{ri, ci}, RIGHT
			case 'v':
				return [2]int{ri, ci}, DOWN
			default:
				continue
			}
		}
	}
	return [2]int{}, -1
}

type Pos struct {
	Row int
	Col int
	Dir int
}

// Moves the guard on the grid until it leaves the grid bounds
func (g *Grid) StartPatrol(maxSteps int) (bool, map[Pos]bool) {
	currentSteps := 0
	prevs := make(map[Pos]bool, 0)
	// While the guard is in bounds:
	for g.isPosInBound(g.guard.Pos()) {
		// Record the last dir and pos
		prevGuardPos := g.guard.Pos()
		prevGuardDir := g.guard.Dir()
		if currentSteps != 0 {
			prevs[Pos{Row: prevGuardPos[0], Col: prevGuardPos[1], Dir: prevGuardDir}] = true
		}
		supposedNextPos := g.guard.NextPos(prevGuardDir)
		if !g.isPosInBound(supposedNextPos) {
			g.grid[prevGuardPos[0]][prevGuardPos[1]] = SeenGlyph(prevGuardDir, g.guard.Dir())
			break
		}
		nextTile := g.grid[supposedNextPos[0]][supposedNextPos[1]]
		for nextTile == OBSTACLE || nextTile == VORTEX {
			// Change the guard direction
			g.guard.ChangeDir()
			supposedNextPos = g.guard.NextPos(g.guard.Dir())
			nextTile = g.grid[supposedNextPos[0]][supposedNextPos[1]]
		}
		// Mark the current tile with the right SEEN glyph
		g.grid[prevGuardPos[0]][prevGuardPos[1]] = SeenGlyph(prevGuardDir, g.guard.Dir())

		// Move the guard position to the next tile
		g.guard.Move()
		nextPos := g.guard.Pos()

		// Show the guard on the grid
		g.grid[nextPos[0]][nextPos[1]] = DIR[g.guard.Dir()]

		currentSteps++
		if currentSteps >= maxSteps && maxSteps != -1 {
			return true, prevs
		}
	}
	return false, prevs
}

// Moves the guard on the grid until it leaves the map, but there's a catch
// This function attempts to find all the grid positions where we can spawn
// a new single box as to get the guard stuck in an infinite loop
// The strategy I am going to apply is to see if spawning a box in front of the guard
// would cause them to divert to a SEEN glyph in the right direction, that way it means
// we are diverting the guard back on its previous tracks, and it should in theory results
// in the guard being stuck in an infinite loop, it cannot leave the grid anymore
// We will want to record the number of different box positions that gets the guard stuck in a loop.
func (g *Grid) StartEvilPatrol(seen map[Pos]bool) int {
	count := 0
	prevVortex := make(map[[2]int]bool, 0)
	for saw, _ := range seen {
		_, prevV := prevVortex[[2]int{saw.Row, saw.Col}]
		if prevV {
			continue
		}
		prevVortex[[2]int{saw.Row, saw.Col}] = true
		gc := utils.DeepCopyMatrix(g.grid)
		ng := NewGrid(gc)
		ng.grid[saw.Row][saw.Col] = VORTEX

		infinite, _ := ng.StartPatrol(25000)
		if infinite {
			count++
		}
	}
	return count
}

func (g *Grid) isPosInBound(pos [2]int) bool {
	if (pos[0] >= 0 && pos[0] < len(g.grid)) && (pos[1] >= 0 && pos[1] < len(g.grid[0])) {
		return true
	}
	return false
}

// Define the seen glyph to put on the tile grid after the guard walked on it
func SeenGlyph(currentDir, previousDir int) rune {
	switch currentDir {
	case UP, DOWN:
		if previousDir == currentDir {
			return '|'
		}
		return '+'
	case LEFT, RIGHT:
		if previousDir == currentDir {
			return '-'
		}
		return '+'
	default:
		panic(fmt.Sprintf("unexpected guardDirection: %s", string(currentDir)))
	}
}
