package grid

import (
	"fmt"
)

type Pos struct {
	x, y int
}

const (
	U = iota
	R
	D
	L
)

// UP RIGHT DOWN LEFT
var DIRS = []Pos{{y: -1, x: 0}, {y: 0, x: 1}, {y: 1, x: 0}, {y: 0, x: -1}}

const (
	WALL  = '#'
	FREE  = '.'
	BOX   = 'O'
	BOXL  = '['
	BOXR  = ']'
	BOT   = '@'
	UP    = '^'
	RIGHT = '>'
	DOWN  = 'v'
	LEFT  = '<'
)

// A grid is bi-dimensional map containing a robot and some boxes
type Grid struct {
	grid     [][]rune
	commands []rune
}

func NewGrid(grid [][]rune, comms []rune) *Grid {
	return &Grid{grid: grid, commands: comms}
}

func (g *Grid) String() string {
	out := ""
	for x, row := range g.grid {
		out += string(row)
		if x+1 != len(row) {
			out += "\n"
		}
	}
	return out
}

func printgrid(grid [][]rune) string {
	out := ""
	for x, row := range grid {
		out += string(row)
		if x+1 != len(row) {
			out += "\n"
		}
	}
	return out
}

func (g *Grid) SumBoxGPS() int {
	sum := 0
	for y, row := range g.grid {
		for x, item := range row {
			if item == BOX || item == BOXL {
				sum += (100 * y) + x
			}
		}
	}
	return sum
}

func (g *Grid) SumBoxGPSV2() int {
	sum := 0
	for y, row := range g.grid {
		for x, item := range row {

			if item == BOXL {
				// left side distance from left edge
				// right side distance from right edge
				// top distance from box
				// bot distance from box
				// deltaL := x
				// deltaR := (len(row) - 1) - (x)
				// deltaU := y
				// deltaD := (len(g.grid) - 1) - y
				// order := []int{deltaD, deltaL, deltaU, deltaR}
				// sort.Slice(order, func(i, j int) bool {
				// 	return order[i] < order[j]
				// })
				if y <= x {
					sum += (100 * y) + x
				} else {
					sum += (100 * x) + y
				}
				// sum += (100 * order[0]) + order[1]
			} else if item == BOXR {
				// left side distance from left edge
				// right side distance from right edge
				// top distance from box
				// bot distance from box
				// deltaL := x
				// deltaR := (len(row) - 1) - x
				// deltaU := y
				// deltaD := (len(g.grid) - 1) - y
				// order := []int{deltaD, deltaL, deltaU, deltaR}
				// sort.Slice(order, func(i, j int) bool {
				// 	return order[i] < order[j]
				// })
				if y <= x {
					sum += (100 * y) + x
				} else {
					sum += (100 * x) + y
				}
				// sum += (100 * order[0]) + order[1]
			}
		}
	}
	return sum
}

// Returns the sum of all box GPS coordinates
func (g *Grid) ProcessCommands() int {
	for _, command := range g.commands {
		// fmt.Printf("Turn %d\n%s\nCommand %s\n", turn, g.String(), string(command))
		bot := g.FindRobot()
		// Identify what's on the next tile
		next, nPos := g.NextTile(bot, command)
		switch next {
		// The bot hits the wall and stays on the same position
		case WALL:
			continue
		// The bot goes on the free tile
		case FREE:
			g.grid[nPos.y][nPos.x] = BOT
			g.grid[bot.y][bot.x] = FREE
		// The bot hits the box and the box will move in the same direction as the bot
		// pushing any other boxes adjacent
		case BOX:
			last, walled := g.findLastBoxPosition(command, nPos)
			// We can't push the box so we continue to next command
			if walled {
				continue
			}
			g.shiftBox(command, last, bot)
		}
	}
	return g.SumBoxGPS()
}

// Returns the last box in a line of adjacent boxes, if walled is true we can't shift the box as they rest against a wall, for part1 only
func (g *Grid) findLastBoxPosition(direction rune, nPos Pos) (last int, walled bool) {
	switch direction {
	case UP:
		last = nPos.y
		breakit := false
		for ny := nPos.y - 1; ny >= 0; ny-- {
			switch g.grid[ny][nPos.x] {
			case WALL:
				walled = true
				breakit = true
			case FREE:
				last = ny + 1
				breakit = true
			}
			if breakit {
				break
			}
		}
	case DOWN:
		last = nPos.y
		breakit := false
		for ny := nPos.y + 1; ny < len(g.grid); ny++ {
			switch g.grid[ny][nPos.x] {
			case WALL:
				walled = true
				breakit = true
			case FREE:
				last = ny - 1
				breakit = true
			}
			if breakit {
				break
			}
		}
	case RIGHT:
		last = nPos.x
		breakit := false
		for nx := nPos.x + 1; nx < len(g.grid); nx++ {
			switch g.grid[nPos.y][nx] {
			case WALL:
				walled = true
				breakit = true
			case FREE:
				last = nx - 1
				breakit = true
			}
			if breakit {
				break
			}
		}
	case LEFT:
		last = nPos.x
		breakit := false
		for nx := nPos.x - 1; nx >= 0; nx-- {
			switch g.grid[nPos.y][nx] {
			case WALL:
				walled = true
				breakit = true
			case FREE:
				last = nx + 1
				breakit = true
			}
			if breakit {
				break
			}
		}
	}
	return
}

// Move the box & bot in the direction, this function is called when we know there is space to shift the bot & box
func (g *Grid) shiftBox(direction rune, lastBoxPos int, bot Pos) {
	switch direction {
	case UP:
		for y := lastBoxPos - 1; y < bot.y; y++ {
			toshift := g.grid[y+1][bot.x]
			g.grid[y][bot.x] = toshift
			g.grid[y+1][bot.x] = FREE
		}
	case DOWN:
		for y := lastBoxPos + 1; y > bot.y; y-- {
			toshift := g.grid[y-1][bot.x]
			g.grid[y][bot.x] = toshift
			g.grid[y-1][bot.x] = FREE
		}
	case RIGHT:
		for x := lastBoxPos + 1; x > bot.x; x-- {
			toshift := g.grid[bot.y][x-1]
			g.grid[bot.y][x] = toshift
			g.grid[bot.y][x-1] = FREE
		}
	case LEFT:
		for x := lastBoxPos - 1; x < bot.x; x++ {
			toshift := g.grid[bot.y][x+1]
			g.grid[bot.y][x] = toshift
			g.grid[bot.y][x+1] = FREE
		}
	}
}

type BoxTree struct {
	root *Box
}

func (bx *BoxTree) Shift(direction rune, grid [][]rune) {
	// To avoid complex tree structure freeing a position we already shifted to
	// we will record the new pos we shifted to, and make sure to not free them on the map
	seen := make(map[Pos]bool)
	bx.root.Shift(direction, grid, seen)
}

func newBoxTree(box *Box) *BoxTree {
	return &BoxTree{root: box}
}

func (b *Box) Shift(direction rune, grid [][]rune, seen map[Pos]bool) {
	// Apply Depth First Search
	if b.next != nil {
		b.next.Shift(direction, grid, seen)
	} else {
		if b.nextL != nil {
			b.nextL.Shift(direction, grid, seen)
		}
		if b.nextR != nil {
			b.nextR.Shift(direction, grid, seen)
		}
	}
	// No next, we arrived at the last leaf, or are recursing back up the tree and the current is our last leaf in the recursion
	switch direction {
	case UP:
		// shift current box up
		_, sawya := seen[Pos{y: b.L.y, x: b.L.x}]
		if !sawya {
			grid[b.L.y][b.L.x] = FREE
		}
		_, sawya2 := seen[Pos{y: b.R.y, x: b.R.x}]
		if !sawya2 {
			grid[b.R.y][b.R.x] = FREE
		}

		grid[b.L.y-1][b.L.x] = BOXL
		grid[b.R.y-1][b.R.x] = BOXR

		seen[Pos{y: b.R.y - 1, x: b.R.x}] = true
		seen[Pos{y: b.L.y - 1, x: b.L.x}] = true

	case DOWN:
		// shift current box down
		_, sawya := seen[Pos{y: b.L.y, x: b.L.x}]
		if !sawya {
			grid[b.L.y][b.L.x] = FREE
		}
		_, sawya2 := seen[Pos{y: b.R.y, x: b.R.x}]
		if !sawya2 {
			grid[b.R.y][b.R.x] = FREE
		}

		grid[b.L.y+1][b.L.x] = BOXL
		grid[b.R.y+1][b.R.x] = BOXR

		seen[Pos{y: b.L.y + 1, x: b.L.x}] = true
		seen[Pos{y: b.R.y + 1, x: b.R.x}] = true
	case LEFT:
		// shift current box left
		grid[b.L.y][b.L.x] = FREE
		grid[b.R.y][b.R.x] = FREE
		grid[b.L.y][b.L.x-1] = BOXL
		grid[b.R.y][b.R.x-1] = BOXR
	case RIGHT:
		// shift current box right
		grid[b.L.y][b.L.x] = FREE
		grid[b.R.y][b.R.x] = FREE
		grid[b.L.y][b.L.x+1] = BOXL
		grid[b.R.y][b.R.x+1] = BOXR
	}
}

/*
NL NR
[][]
L[]R
[][]
NL NR
*/
type Box struct {
	L, R               Pos // The left and right part of the box coordinates
	nextL, next, nextR *Box
}

func newBox(left, right Pos) *Box {
	return &Box{L: left, R: right}
}

func (b *Box) SetNextLeft(next *Box) {
	b.nextL = next
}

func (b *Box) SetNextRight(next *Box) {
	b.nextR = next
}

// Each box builds its own tree
func (b *Box) FindNext(grid [][]rune, direction rune) (continue_ bool) {
	continue_ = true
	switch direction {
	case UP:
		// Look above L & R pos
		upL, upR := grid[b.L.y-1][b.L.x], grid[b.R.y-1][b.R.x]
		skipR := false
		switch upL {
		case BOXR:
			// There is a box top left of this one
			next := newBox(Pos{y: b.L.y - 1, x: b.L.x - 1}, Pos{y: b.L.y - 1, x: b.L.x})
			b.nextL = next
			continue_ = b.nextL.FindNext(grid, direction)
			if !continue_ {
				return false
			}
		case BOXL:
			// There is a box right above it
			next := newBox(Pos{y: b.L.y - 1, x: b.L.x}, Pos{y: b.R.y - 1, x: b.R.x})
			b.next = next
			return b.next.FindNext(grid, direction)
		case WALL:
			return false
		}
		if !skipR {
			switch upR {
			case BOXR:
				panic("Should have been skipped, map or logic must be wrong, look last map logged to figure this out")
			case BOXL:
				// There is a box top right of this one
				next := newBox(Pos{y: b.R.y - 1, x: b.R.x}, Pos{y: b.R.y - 1, x: b.R.x + 1})
				b.nextR = next
				return b.nextR.FindNext(grid, direction)
			case WALL:
				return false
			}
		}
	case DOWN:
		// Look down L & R pos
		if b.L.y+1 >= len(grid) {
			fmt.Printf("%s\n", printgrid(grid))
		}
		dL, dR := grid[b.L.y+1][b.L.x], grid[b.R.y+1][b.R.x]
		skipR := false
		switch dL {
		case BOXR:
			// There is a box down left of this one
			next := newBox(Pos{y: b.L.y + 1, x: b.L.x - 1}, Pos{y: b.L.y + 1, x: b.L.x})
			b.nextL = next
			continue_ = b.nextL.FindNext(grid, direction)
			if !continue_ {
				return false
			}
		case BOXL:
			// There is a box right under it
			next := newBox(Pos{y: b.L.y + 1, x: b.L.x}, Pos{y: b.R.y + 1, x: b.R.x})
			b.next = next
			return b.next.FindNext(grid, direction)
		case WALL:
			return false
		}
		if !skipR {
			switch dR {
			case BOXR:
				panic(fmt.Sprintf("Should have been skipped, map must be wrong %v", grid))
			case BOXL:
				// There is a box down right of this one
				next := newBox(Pos{y: b.R.y + 1, x: b.R.x}, Pos{y: b.R.y + 1, x: b.R.x + 1})
				b.nextR = next
				return b.nextR.FindNext(grid, direction)
			case WALL:
				return false
			}
		}
	case RIGHT:
		// Look right of current box
		switch grid[b.R.y][b.R.x+1] {
		case BOXL:
			next := newBox(Pos{y: b.R.y, x: b.R.x + 1}, Pos{y: b.R.y, x: b.R.x + 2})
			b.next = next
			return b.next.FindNext(grid, direction)
		case BOXR:
			panic("shouldn't have a box right next to a box right")
		case WALL:
			return false
		}
	case LEFT:
		switch grid[b.L.y][b.L.x-1] {
		case BOXL:
			panic("shouldn't have a box left next to a box left")
		case BOXR:
			next := newBox(Pos{y: b.L.y, x: b.L.x - 2}, Pos{y: b.L.y, x: b.L.x - 1})
			b.next = next
			return b.next.FindNext(grid, direction)
		case WALL:
			return false
		}
	}
	return
}

// Returns the next tile type and position coordinates depending on the current position and direction
func (g *Grid) NextTile(curr Pos, direction rune) (rune, Pos) {
	nPos := Pos{}
	switch direction {
	case UP:
		nPos = Pos{x: curr.x + DIRS[U].x, y: curr.y + DIRS[U].y}
	case DOWN:
		nPos = Pos{x: curr.x + DIRS[D].x, y: curr.y + DIRS[D].y}
	case RIGHT:
		nPos = Pos{x: curr.x + DIRS[R].x, y: curr.y + DIRS[R].y}
	case LEFT:
		nPos = Pos{x: curr.x + DIRS[L].x, y: curr.y + DIRS[L].y}
	}
	return g.grid[nPos.y][nPos.x], nPos
}

// Returns the sum of all box GPS coordinates for part 2
func (g *Grid) ProcessCommandsV2() int {
	for _, command := range g.commands {
		// utils.LogToHistory(g.String(), command, turn)
		// fmt.Printf("Turn %d\n%s\nCommand %s\n", turn, g.String(), string(command))
		bot := g.FindRobot()
		// Identify what's on the next tile
		next, nPos := g.NextTile(bot, command)
		switch next {
		case WALL:
			continue
		case FREE:
			g.grid[nPos.y][nPos.x] = BOT
			g.grid[bot.y][bot.x] = FREE
		case BOXL, BOXR:
			box := Box{}
			if next == BOXL {
				box.L = Pos{y: nPos.y, x: nPos.x}
				box.R = Pos{y: nPos.y, x: nPos.x + 1}
			} else {
				box.R = Pos{y: nPos.y, x: nPos.x}
				box.L = Pos{y: nPos.y, x: nPos.x - 1}
			}
			continue_ := box.FindNext(g.grid, command)
			if continue_ {
				bt := newBoxTree(&box)
				bt.Shift(command, g.grid)
				g.grid[nPos.y][nPos.x] = BOT
				g.grid[bot.y][bot.x] = FREE
			}
		}
	}
	return g.SumBoxGPS()
}

func (g *Grid) FindRobot() Pos {
	for y, row := range g.grid {
		for x, char := range row {
			if char == BOT {
				return Pos{x: x, y: y}
			}
		}
	}
	return Pos{}
}
