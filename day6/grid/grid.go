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
	DOWN
	LEFT
	RIGHT
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

// Moves the guard on the grid until it leaves the grid bounds
func (g *Grid) StartPatrol(maxSteps int) (isInfinite bool) {
	currentSteps := 0
	// prevs := make(map[[2]int][]int, 0)
	// initDir := g.guard.Dir()
	// initPos := g.guard.Pos()
	// prevs[initPos] = append(prevs[initPos], initDir)
	// While the guard is in bounds:
	for g.isPosInBound(g.guard.Pos()) {
		// Record the last dir and pos
		prevGuardPos := g.guard.Pos()
		prevGuardDir := g.guard.Dir()
		supposedNextPos := g.guard.NextPos(prevGuardDir)
		if !g.isPosInBound(supposedNextPos) {
			g.grid[prevGuardPos[0]][prevGuardPos[1]] = SeenGlyph(prevGuardDir, g.guard.Dir())
			break
		}
		nextTile := g.grid[supposedNextPos[0]][supposedNextPos[1]]
		switch nextTile {
		case OBSTACLE, VORTEX:
			// Change the guard direction
			g.guard.ChangeDir()
		}
		// Mark the current tile with the right SEEN glyph
		g.grid[prevGuardPos[0]][prevGuardPos[1]] = SeenGlyph(prevGuardDir, g.guard.Dir())
		// Move the guard position to the next tile
		g.guard.Move()
		nextPos := g.guard.Pos()

		// Show the guard on the grid
		g.grid[nextPos[0]][nextPos[1]] = DIR[prevGuardDir]
		currentSteps++
		if currentSteps >= maxSteps && maxSteps != -1 {
			return true
		}
		// if maxSteps != -1 {
		// 	currDir := g.guard.Dir()
		// 	currPos := g.guard.Pos()
		// 	dirs, visited := prevs[currPos]
		// 	if visited {
		// 		for _, dir := range dirs {
		// 			if dir == currDir {
		// 				return true
		// 			}
		// 		}
		// 	}
		// 	prevs[currPos] = append(prevs[currPos], currDir)
		// }

	}
	return false
}

// Moves the guard on the grid until it leaves the map, but there's a catch
// This function attempts to find all the grid positions where we can spawn
// a new single box as to get the guard stuck in an infinite loop
// The strategy I am going to apply is to see if spawning a box in front of the guard
// would cause them to divert to a SEEN glyph in the right direction, that way it means
// we are diverting the guard back on its previous tracks, and it should in theory results
// in the guard being stuck in an infinite loop, it cannot leave the grid anymore
// We will want to record the number of different box positions that gets the guard stuck in a loop.
func (g *Grid) StartEvilPatrol() int {
	// startPos := g.guard.Pos()
	knownVortexPos := make(map[[2]int]bool, 0)
	evilVortex := 0
	// While the guard is in bounds:
	for g.isPosInBound(g.guard.Pos()) {
		// Record the last dir and pos
		prevGuardPos := g.guard.Pos()
		prevGuardDir := g.guard.Dir()
		supposedNextPos := g.guard.NextPos(prevGuardDir)
		if !g.isPosInBound(supposedNextPos) {
			g.grid[prevGuardPos[0]][prevGuardPos[1]] = SeenGlyph(prevGuardDir, g.guard.Dir())
			break
		}
		nextTile := g.grid[supposedNextPos[0]][supposedNextPos[1]]
		switch nextTile {
		case OBSTACLE, VORTEX:
			g.guard.ChangeDir()
		}

		// Mark the current tile with the right SEEN glyph
		g.grid[prevGuardPos[0]][prevGuardPos[1]] = SeenGlyph(prevGuardDir, g.guard.Dir())
		// Move the guard position to the next tile
		g.guard.Move()

		nextPos := g.guard.Pos()
		// Show the guard on the grid
		g.grid[nextPos[0]][nextPos[1]] = DIR[g.guard.Dir()]

		// Try to get the guard stuck
		// Check if newTile with newDir is SEEN and if it's the right direction
		newDir := g.guard.NextDir()
		newPos := g.guard.NextPos(newDir)
		newTile := g.grid[newPos[0]][newPos[1]]
		supposedNextPos = g.guard.NextPos(g.guard.Dir())
		if !g.isPosInBound(supposedNextPos) {
			continue
		}
		supposedNextTile := g.grid[supposedNextPos[0]][supposedNextPos[1]]
		beEvil := false
		if newTile != OBSTACLE && supposedNextTile != OBSTACLE {
			// if !reflect.DeepEqual(newPos, startPos) {
			// 	beEvil = true
			// }
			beEvil = true
		}
		if beEvil {
			// Spawn the vortex in front of the guard in a copy of the grid
			gridCopy := utils.DeepCopyMatrix(g.grid)
			vortexPos := g.guard.NextPos(g.guard.Dir())
			if !g.isPosInBound(vortexPos) {
				continue
			}
			// is the vortex pos is already known?
			_, alreadyKnown := knownVortexPos[vortexPos]
			if alreadyKnown {
				continue
			}
			gridCopy[vortexPos[0]][vortexPos[1]] = VORTEX
			evilGrid := NewGrid(gridCopy)
			isInfinite := evilGrid.StartPatrol(7000)
			if isInfinite {
				evilVortex++
				// fmt.Printf("Infinite loop detected with maze:\n%s\n", evilGrid.PrintMaze())
			}
		}
	}
	return evilVortex
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
