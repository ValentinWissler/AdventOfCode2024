package grid

import (
	a "aoc8/antena"
	"fmt"
	"unicode"
)

// A grid is bi-dimensional map containing antenas
type Grid struct {
	// Raw input containing the antenas positions
	input [][]rune
	// The antenas on the grid, the key represent the type of antena
	antenas      map[rune][]*a.Antena
	antinodesPos map[[2]int]bool
}

func NewGrid(grid [][]rune) *Grid {
	return &Grid{input: grid}
}

// Find the different type of antenas on the grid to fill the antenas property
func (g *Grid) FindAntenas() {
	antenaPos := make(map[rune][][2]int)
	for ri, row := range g.input {
		for ci, item := range row {
			if unicode.IsLetter(item) || unicode.IsNumber(item) {
				antenaPos[item] = append(antenaPos[item], [2]int{ri, ci})
			}
		}
	}
	g.antenas = make(map[rune][]*a.Antena)
	for type_, allPos := range antenaPos {
		antenas := make([]*a.Antena, 0)
		for curr, pos := range allPos {
			sisters := findSisters(allPos, curr)
			fmt.Printf("Antena %s in %v with sisters %v\n", string(type_), pos, sisters)
			antenas = append(antenas, a.NewAntena(type_, pos, sisters))
		}
		g.antenas[type_] = append(g.antenas[type_], antenas...)
		fmt.Printf("Added %d antenas of type %s\n", len(antenas), string(type_))
	}
}

func findSisters(positions [][2]int, curr int) [][2]int {
	switch curr {
	case 0: // First
		return positions[1:]
	case len(positions) - 1: // Last
		return positions[:len(positions)-1]
	default: // Middle
		b := positions[:curr]
		a := positions[curr+1:]
		return append(a, b...)
	}
}

// Send a signal to each antena to emit its signal
// in order to capture the reflection of this signal
// behind the receiving towers
func (g *Grid) AntenasEmit() {
	g.antinodesPos = make(map[[2]int]bool)
	for _, antenas := range g.antenas {
		for _, antena := range antenas {
			reflectedSignalPos := antena.Emit()
			// Record antinodes position in bounds
			for _, pos := range reflectedSignalPos {
				if (pos[0] >= 0 && pos[0] < len(g.input)) && (pos[1] >= 0 && pos[1] < len(g.input[0])) {
					g.antinodesPos[pos] = true
				}
			}
		}
	}
}

// Send a signal to each antena to emit its signal
// in order to capture the reflection of this signal
// behind the receiving towers
func (g *Grid) AntenasEmitV2() {
	g.antinodesPos = make(map[[2]int]bool)
	for _, antenas := range g.antenas {
		for _, antena := range antenas {
			reflectedSignalPos := antena.EmitV2(len(g.input), len(g.input[0]))
			// Record antinodes position in bounds
			for _, pos := range reflectedSignalPos {
				if (pos[0] >= 0 && pos[0] < len(g.input)) && (pos[1] >= 0 && pos[1] < len(g.input[0])) {
					g.antinodesPos[pos] = true
				}
			}
		}
	}
}

func (g *Grid) CountAntinodes() int {
	return len(g.antinodesPos)
}
