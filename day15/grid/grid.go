package grid

import "fmt"

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

func (g *Grid) SumBoxGPS() int {
	sum := 0
	for y, row := range g.grid {
		for x, item := range row {
			if item == BOX {
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
			if item == BOX {
				if y <= x {
					sum += (100 * y) + x
				} else {
					sum += (100 * x) + y
				}
			}
		}
	}
	return sum
}

// Returns the sum of all box GPS coordinates
func (g *Grid) ProcessCommands() int {
	for turn, command := range g.commands {
		fmt.Printf("Turn %d\n%s\nCommand %s\n", turn, g.String(), string(command))
		bot := g.FindRobot()
		// Identify what's on the next tile
		nPos := Pos{}
		switch command {
		case UP:
			nPos = Pos{x: bot.x + DIRS[U].x, y: bot.y + DIRS[U].y}
		case DOWN:
			nPos = Pos{x: bot.x + DIRS[D].x, y: bot.y + DIRS[D].y}
		case RIGHT:
			nPos = Pos{x: bot.x + DIRS[R].x, y: bot.y + DIRS[R].y}
		case LEFT:
			nPos = Pos{x: bot.x + DIRS[L].x, y: bot.y + DIRS[L].y}
		}
		switch g.grid[nPos.y][nPos.x] {
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
			// Check how many boxes are going to get pushed
			// if we're walled we can't push the boxes as the last box is touching a wall
			// and record the last box Y position should we have to shift all the boxes
			walled := false
			switch command {
			case UP:
				lastY := nPos.y
				breakit := false
				for ny := nPos.y - 1; ny >= 0; ny-- {
					switch g.grid[ny][nPos.x] {
					case WALL:
						walled = true
						breakit = true
					case FREE:
						lastY = ny + 1
						breakit = true
					}
					if breakit {
						break
					}
				}
				// We can't push the box so we continue to next command
				if walled {
					continue
				}
				// Shift the bot and the boxes in the direction
				for y := lastY - 1; y < bot.y; y++ {
					toshift := g.grid[y+1][bot.x]
					g.grid[y][bot.x] = toshift
					g.grid[y+1][bot.x] = FREE
				}
			case DOWN:
				lastY := nPos.y
				breakit := false
				for ny := nPos.y + 1; ny < len(g.grid); ny++ {
					switch g.grid[ny][nPos.x] {
					case WALL:
						walled = true
						breakit = true
					case FREE:
						lastY = ny - 1
						breakit = true
					}
					if breakit {
						break
					}
				}
				// We can't push the box so we continue to next command
				if walled {
					continue
				}
				// Shift the bot and the boxes in the direction
				for y := lastY + 1; y > bot.y; y-- {
					toshift := g.grid[y-1][bot.x]
					g.grid[y][bot.x] = toshift
					g.grid[y-1][bot.x] = FREE
				}
			case RIGHT:
				lastX := nPos.x
				breakit := false
				for nx := nPos.x + 1; nx < len(g.grid); nx++ {
					switch g.grid[nPos.y][nx] {
					case WALL:
						walled = true
						breakit = true
					case FREE:
						lastX = nx - 1
						breakit = true
					}
					if breakit {
						break
					}
				}
				// We can't push the box so we continue to next command
				if walled {
					continue
				}
				// Shift the bot and the boxes in the direction
				for x := lastX + 1; x > bot.x; x-- {
					toshift := g.grid[bot.y][x-1]
					g.grid[bot.y][x] = toshift
					g.grid[bot.y][x-1] = FREE
				}
			case LEFT:
				lastX := nPos.x
				breakit := false
				for nx := nPos.x - 1; nx >= 0; nx-- {
					switch g.grid[nPos.y][nx] {
					case WALL:
						walled = true
						breakit = true
					case FREE:
						lastX = nx + 1
						breakit = true
					}
					if breakit {
						break
					}
				}
				// We can't push the box so we continue to next command
				if walled {
					continue
				}
				// Shift the bot and the boxes in the direction
				for x := lastX - 1; x < bot.x; x++ {
					toshift := g.grid[bot.y][x+1]
					g.grid[bot.y][x] = toshift
					g.grid[bot.y][x+1] = FREE
				}
			}
		}
	}
	return g.SumBoxGPS()
}

// Returns the sum of all box GPS coordinates for part 2
func (g *Grid) ProcessCommandsV2() int {
	for turn, command := range g.commands {
		fmt.Printf("Turn %d\n%s\nCommand %s\n", turn, g.String(), string(command))
		bot := g.FindRobot()
		// Identify what's on the next tile
		nPos := Pos{}
		switch command {
		case UP:
			nPos = Pos{x: bot.x + DIRS[U].x, y: bot.y + DIRS[U].y}
		case DOWN:
			nPos = Pos{x: bot.x + DIRS[D].x, y: bot.y + DIRS[D].y}
		case RIGHT:
			nPos = Pos{x: bot.x + DIRS[R].x, y: bot.y + DIRS[R].y}
		case LEFT:
			nPos = Pos{x: bot.x + DIRS[L].x, y: bot.y + DIRS[L].y}
		}
		next := g.grid[nPos.y][nPos.x]
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
		case BOXL, BOXR:
			// Check how many boxes are going to get pushed
			// if we're walled we can't push the boxes as the last box is touching a wall
			// and record the last box Y position should we have to shift all the boxes
			walled := false
			switch command {
			case UP:
				lastY := nPos.y
				breakit := false
				if next == BOXL {
					// Look up and up right
					for ny := nPos.y - 1; ny >= 0; ny-- {
						// Look up
						switch g.grid[ny][nPos.x] {
						case WALL:
							walled = true
							breakit = true
						case FREE:
							lastY = ny + 1
							breakit = true
						}
						if breakit {
							break
						}
						// Look up right
						switch g.grid[ny][nPos.x+1] {
						case WALL:
							walled = true
							breakit = true
						case FREE:
							lastY = ny + 1
							breakit = true
						}
						if breakit {
							break
						}
					}
					// We can't push the box so we continue to next command
					if walled {
						continue
					}
					// Shift the bot and the boxes in the direction
					for y := lastY - 1; y < bot.y; y++ {
						toshiftUp := g.grid[y+1][bot.x]
						toshiftUpRight := g.grid[y+1][bot.x+1]
						g.grid[y][bot.x] = toshiftUp
						g.grid[y][bot.x+1] = toshiftUpRight
						g.grid[y+1][bot.x] = FREE
						g.grid[y+1][bot.x+1] = FREE
					}
				} else if next == BOXR {
					// Look up and up left
					for ny := nPos.y - 1; ny >= 0; ny-- {
						// Look up
						switch g.grid[ny][nPos.x] {
						case WALL:
							walled = true
							breakit = true
						case FREE:
							lastY = ny + 1
							breakit = true
						}
						if breakit {
							break
						}
						// Look up left
						switch g.grid[ny][nPos.x-1] {
						case WALL:
							walled = true
							breakit = true
						case FREE:
							lastY = ny + 1
							breakit = true
						}
						if breakit {
							break
						}
					}
					// We can't push the box so we continue to next command
					if walled {
						continue
					}
					for y := lastY - 1; y < bot.y; y++ {
						toshiftUp := g.grid[y+1][bot.x]
						toshiftUpLeft := g.grid[y+1][bot.x-1]
						g.grid[y][bot.x] = toshiftUp
						g.grid[y][bot.x-1] = toshiftUpLeft
						g.grid[y+1][bot.x] = FREE
						g.grid[y+1][bot.x-1] = FREE
					}
				}
			case DOWN:
				lastY := nPos.y
				breakit := false
				if next == BOXL {
					// Look down and down right
					for ny := nPos.y + 1; ny < len(g.grid); ny++ {
						// Look down
						switch g.grid[ny][nPos.x] {
						case WALL:
							walled = true
							breakit = true
						case FREE:
							lastY = ny - 1
							breakit = true
						}
						if breakit {
							break
						}
						// Look down right
						switch g.grid[ny][nPos.x+1] {
						case WALL:
							walled = true
							breakit = true
						case FREE:
							lastY = ny - 1
							breakit = true
						}
						if breakit {
							break
						}
					}
					// We can't push the box so we continue to next command
					if walled {
						continue
					}
					// Shift the bot and the boxes in the direction
					for y := lastY + 1; y > bot.y; y-- {
						toshiftDown := g.grid[y-1][bot.x]
						toshiftDownRight := g.grid[y-1][bot.x+1]
						g.grid[y][bot.x] = toshiftDown
						g.grid[y][bot.x+1] = toshiftDownRight
						g.grid[y-1][bot.x] = FREE
						g.grid[y-1][bot.x+1] = FREE
					}
				} else if next == BOXR {
					// Look down and down left
					for ny := nPos.y + 1; ny < len(g.grid); ny++ {
						// Look up
						switch g.grid[ny][nPos.x] {
						case WALL:
							walled = true
							breakit = true
						case FREE:
							lastY = ny - 1
							breakit = true
						}
						if breakit {
							break
						}
						// Look up left
						switch g.grid[ny][nPos.x-1] {
						case WALL:
							walled = true
							breakit = true
						case FREE:
							lastY = ny - 1
							breakit = true
						}
						if breakit {
							break
						}
					}
					// We can't push the box so we continue to next command
					if walled {
						continue
					}
					// Shift the bot and the boxes in the direction
					for y := lastY + 1; y > bot.y; y-- {
						toshiftDown := g.grid[y-1][bot.x]
						toshiftDownLeft := g.grid[y-1][bot.x-1]
						g.grid[y][bot.x] = toshiftDown
						g.grid[y][bot.x-1] = toshiftDownLeft
						g.grid[y-1][bot.x] = FREE
						g.grid[y-1][bot.x-1] = FREE
					}
				}
			case RIGHT:
				lastX := nPos.x
				breakit := false
				for nx := nPos.x + 1; nx < len(g.grid); nx++ {
					switch g.grid[nPos.y][nx] {
					case WALL:
						walled = true
						breakit = true
					case FREE:
						lastX = nx - 1
						breakit = true
					}
					if breakit {
						break
					}
				}
				// We can't push the box so we continue to next command
				if walled {
					continue
				}
				// Shift the bot and the boxes in the direction
				for x := lastX + 1; x > bot.x; x-- {
					toshift := g.grid[bot.y][x-1]
					g.grid[bot.y][x] = toshift
					g.grid[bot.y][x-1] = FREE
				}
			case LEFT:
				lastX := nPos.x
				breakit := false
				for nx := nPos.x - 1; nx >= 0; nx-- {
					switch g.grid[nPos.y][nx] {
					case WALL:
						walled = true
						breakit = true
					case FREE:
						lastX = nx + 1
						breakit = true
					}
					if breakit {
						break
					}
				}
				// We can't push the box so we continue to next command
				if walled {
					continue
				}
				// Shift the bot and the boxes in the direction
				for x := lastX - 1; x < bot.x; x++ {
					toshift := g.grid[bot.y][x+1]
					g.grid[bot.y][x] = toshift
					g.grid[bot.y][x+1] = FREE
				}
			}
		}
	}
	return g.SumBoxGPSV2()
}

func (g *Grid) FindRobot() Pos {
	for y, row := range g.grid {
		for x, char := range row {
			if char == '@' {
				return Pos{x: x, y: y}
			}
		}
	}
	return Pos{}
}
