package antena

type Antena struct {
	// Antenas have a type, they only emit towards antena of the same type
	type_ rune
	// Antenas have a position on the grid
	pos [2]int
	// Antenas emit to sister antenas and need to know their position
	sistersPos [][2]int
}

func NewAntena(type_ rune, pos [2]int, sisterPos [][2]int) *Antena {
	return &Antena{type_: type_, pos: pos, sistersPos: sisterPos}
}

// Emits a signal to sisters antena and returns the reflected node position
// A reflected node position is defined as the position twice longer than the path from tower a to b
// that spans from a to node pos
/*
.0.
...
...
.A.
...
...
.A.
...
...
.0.

*/
// The antinode positions are returned for this antena
func (a *Antena) Emit() [][2]int {
	currR, currC := a.pos[0], a.pos[1]
	antinodes := make([][2]int, 0)
	for _, sisPos := range a.sistersPos {
		// Calculate how much horizontal and vertical movement are required
		movement := [2]int{(currR - sisPos[0]) * -1, (currC - sisPos[1]) * -1}
		// Apply this movement to the sister pos and record the position
		antinode := [2]int{sisPos[0] + movement[0], sisPos[1] + movement[1]}
		antinodes = append(antinodes, antinode)
	}
	return antinodes
}

// The antinode positions are returned for this antena
func (a *Antena) EmitV2(maxV, maxH int) [][2]int {
	currR, currC := a.pos[0], a.pos[1]
	antinodes := make([][2]int, 0)
	for _, sisPos := range a.sistersPos {
		// Calculate how much horizontal and vertical movement are required
		movement := [2]int{(currR - sisPos[0]) * -1, (currC - sisPos[1]) * -1}
		// Apply this movement to the sister pos and record the position
		antinode := [2]int{sisPos[0] + movement[0], sisPos[1] + movement[1]}
		for isWithinBound(maxV, maxH, antinode) {
			antinodes = append(antinodes, antinode)
			antinode = [2]int{antinode[0] + movement[0], antinode[1] + movement[1]}
		}

	}
	if len(a.sistersPos) > 0 {
		antinodes = append(antinodes, a.pos)
	}
	return antinodes
}

func isWithinBound(maxV, maxH int, pos [2]int) bool {
	if (pos[0] >= 0 && pos[0] < maxV) && (pos[1] >= 0 && pos[1] < maxH) {
		return true
	}
	return false
}
