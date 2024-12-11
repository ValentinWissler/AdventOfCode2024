package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	INPUT_FILE = "./input.txt" // Puzzle input
	TEST_FILE  = "./mini.txt"  // To test if the code is working, the mini contains 18 XMAS#
	UP         = '^'
	DOWN       = 'v'
	RIGHT      = '>'
	LEFT       = '<'
	FREE       = '.'
	SEEN       = '0'
	SEEN1      = '-'
	SEEN2      = '|'
	SEEN3      = '+'
	START      = 'E'
	OBSTACLE   = '#'
	VORTEX     = '0'
)

// Directions the guard can take
var (
	DIR_UP    = []int{-1, 0}
	DIR_DOWN  = []int{1, 0}
	DIR_LEFT  = []int{0, -1}
	DIR_RIGHT = []int{0, 1}
)

func main() {
	content, err := ReadFile(TEST_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Format the input
	str := strings.Split(content, "\n")
	maze := make([][]rune, 0)
	for _, m := range str {
		maze = append(maze, []rune(strings.TrimSpace(m)))
	}
	mazeCopy := deepCopyMatrix(maze)

	c := CountStepsV1(mazeCopy)
	fmt.Println("Seen v1: ", c)

	mazeCopy = deepCopyMatrix(maze)

	d := CountStepsV2(mazeCopy)
	fmt.Println("Seen v2: ", d)
}

func CountStepsV1(maze [][]rune) int {
	_, steps := WalkThePlantV2(maze, -1, false)
	return steps
}

func CountStepsV2(maze [][]rune) int {
	_, steps := WalkThePlantV2(maze, -1, true)
	return steps
}

func MoveGuard(maze [][]rune, currPos, newPos [2]int, guardDirection rune) ([][]rune, [2]int, bool, rune, rune) {
	// Maze limits
	// fmt.Printf("New Guard pos: %d|%d", newPos[0], newPos[1])
	var maxV, maxH = len(maze), len(maze[0])
	var next rune
	// Check position is within bounds
	newDir := '^'
	if (newPos[0] >= 0 && newPos[0] < maxV) && (newPos[1] >= 0 && newPos[1] < maxH) {
		next = maze[newPos[0]][newPos[1]]
		switch guardDirection {
		case UP:
			switch next {
			case FREE, SEEN1, SEEN2, SEEN3, START:
				maze[newPos[0]][newPos[1]] = UP
				newDir = guardDirection
			case OBSTACLE, VORTEX:
				next = SEEN2 // '|'
				newPos[0], newPos[1] = currPos[0]+DIR_RIGHT[0], currPos[1]+DIR_RIGHT[1]
				maze[newPos[0]][newPos[1]] = RIGHT
				newDir = RIGHT
			default:
				panic(fmt.Sprintf("unknown next: %s", string(next)))
			}
			// Mark the previous glyph as seen
			maze[currPos[0]][currPos[1]] = SeenGlyph(newDir, next)
		case DOWN:
			switch next {
			case FREE, SEEN1, SEEN2, SEEN3, START:
				maze[newPos[0]][newPos[1]] = DOWN
				newDir = DOWN
			case OBSTACLE, VORTEX:
				next = SEEN2 // '|'
				newPos[0], newPos[1] = currPos[0]+DIR_LEFT[0], currPos[1]+DIR_LEFT[1]
				maze[newPos[0]][newPos[1]] = LEFT
				newDir = LEFT
			default:
				PrintMaze(maze)
				panic(fmt.Sprintf("unknown next: %s", string(next)))
			}
			maze[currPos[0]][currPos[1]] = SeenGlyph(newDir, next)
		case LEFT:
			switch next {
			case FREE, SEEN1, SEEN2, SEEN3, START:
				maze[newPos[0]][newPos[1]] = LEFT
				newDir = LEFT
			case OBSTACLE, VORTEX:
				next = SEEN1
				newPos[0], newPos[1] = currPos[0]+DIR_UP[0], currPos[1]+DIR_UP[1]
				maze[newPos[0]][newPos[1]] = UP
				newDir = UP
			default:
				panic(fmt.Sprintf("unknown next: %s", string(next)))
			}
			maze[currPos[0]][currPos[1]] = SeenGlyph(newDir, next)
		case RIGHT:
			switch next {
			case FREE, SEEN1, SEEN2, SEEN3, START:
				maze[newPos[0]][newPos[1]] = RIGHT
				newDir = RIGHT
			case OBSTACLE, VORTEX:
				next = SEEN1
				newPos[0], newPos[1] = currPos[0]+DIR_DOWN[0], currPos[1]+DIR_DOWN[1]
				maze[newPos[0]][newPos[1]] = DOWN
				newDir = DOWN
			default:
				panic(fmt.Sprintf("unknown next: %s", string(next)))
			}
			maze[currPos[0]][currPos[1]] = SeenGlyph(newDir, next)
		default:
			panic(fmt.Sprintf("unknown guardDirection: %s", string(guardDirection)))
		}
	} else { // Position is OOO, we reached the end
		return maze, currPos, true, next, newDir
	}
	currPos = newPos
	// fmt.Printf("In Guard pos: %d|%d", currPos[0], currPos[1])
	return maze, currPos, false, next, newDir
}

func WalkThePlantV2(maze [][]rune, maxSteps int, probeInfiniteLoop bool) (infiniteLoop bool, steps int) {
	// Get initial guard position
	guardPosition := GetGuardPosition(maze)
	stop := false
	count := 0
	// next := 'E'
	guardDirection := maze[guardPosition[0]][guardPosition[1]]
	for !stop {
		switch guardDirection {
		case UP:
			// Define next position
			newPos := [2]int{guardPosition[0] + DIR_UP[0], guardPosition[1] + DIR_UP[1]}
			// fmt.Printf("Old Guard pos: %d|%d", guardPosition[0], guardPosition[1])
			maze, guardPosition, stop, _, guardDirection = MoveGuard(maze, guardPosition, newPos, guardDirection)
			// fmt.Printf("New Guard pos: %d|%d", guardPosition[0], guardPosition[1])
			if stop {
				stop = true
				break
			}
			if probeInfiniteLoop {
				// If current is '+' and we were to spawn an object in front of us
				// Would the next pos be a SEEN pos? If yes, the guard will likely end in a loop
				newPos = [2]int{guardPosition[0] + DIR_RIGHT[0], guardPosition[1] + DIR_RIGHT[1]}
				newSquare := maze[newPos[0]][newPos[1]]
				switch newSquare {
				case SEEN1, SEEN3:
					mazeCopy := deepCopyMatrix(maze)
					// Put the new obstacle
					mazeCopy[guardPosition[0]+DIR_UP[0]][guardPosition[1]+DIR_UP[1]] = VORTEX
					// Move guard
					mazeCopy[newPos[0]][newPos[1]] = RIGHT
					infiniteLoop, _ := WalkThePlantV2(mazeCopy, 1000, false)
					if infiniteLoop {
						fmt.Println("Infinit loop detected with maze")
						PrintMaze(mazeCopy)
					}
				}
			}
			// PrintMaze(maze)
		case DOWN:
			// Define next position
			newPos := [2]int{guardPosition[0] + DIR_DOWN[0], guardPosition[1] + DIR_DOWN[1]}
			maze, guardPosition, stop, _, guardDirection = MoveGuard(maze, guardPosition, newPos, guardDirection)
			if stop {
				stop = true
				break
			}
			// Mark the previous glyph as seen
			if probeInfiniteLoop {
				// If current is '+' and we were to spawn an object in front of us
				// Would the next pos be a SEEN pos? If yes, the guard will likely end in a loop
				newPos = [2]int{guardPosition[0] + DIR_LEFT[0], guardPosition[1] + DIR_LEFT[1]}
				newSquare := maze[newPos[0]][newPos[1]]
				switch newSquare {
				case SEEN1, SEEN3:
					mazeCopy := deepCopyMatrix(maze)
					// Put the new obstacle
					mazeCopy[guardPosition[0]+DIR_DOWN[0]][guardPosition[1]+DIR_DOWN[1]] = VORTEX
					// Move guard
					mazeCopy[newPos[0]][newPos[1]] = LEFT
					infiniteLoop, _ := WalkThePlantV2(mazeCopy, 1000, false)
					if infiniteLoop {
						fmt.Println("Infinit loop detected with maze")
						PrintMaze(mazeCopy)
					}
				}
			}

			// PrintMaze(maze)
		case RIGHT:
			// Define next position
			newPos := [2]int{guardPosition[0] + DIR_RIGHT[0], guardPosition[1] + DIR_RIGHT[1]}
			// Check position is within bounds
			maze, guardPosition, stop, _, guardDirection = MoveGuard(maze, guardPosition, newPos, guardDirection)
			if stop {
				stop = true
				break
			}

			if probeInfiniteLoop {
				// If current is '+' and we were to spawn an object in front of us
				// Would the next pos be a SEEN pos? If yes, the guard will likely end in a loop
				newPos = [2]int{guardPosition[0] + DIR_DOWN[0], guardPosition[1] + DIR_DOWN[1]}
				newSquare := maze[newPos[0]][newPos[1]]
				switch newSquare {
				case SEEN2, SEEN3:
					mazeCopy := deepCopyMatrix(maze)
					// Put the new obstacle
					mazeCopy[guardPosition[0]+DIR_RIGHT[0]][guardPosition[1]+DIR_RIGHT[1]] = VORTEX
					// Move guard
					mazeCopy[newPos[0]][newPos[1]] = DOWN
					infiniteLoop, _ := WalkThePlantV2(mazeCopy, 1000, false)
					if infiniteLoop {
						fmt.Println("Infinit loop detected with maze")
						PrintMaze(mazeCopy)
					}
				}
			}

			// PrintMaze(maze)
		case LEFT:
			// Define next position
			newPos := [2]int{guardPosition[0] + DIR_LEFT[0], guardPosition[1] + DIR_LEFT[1]}
			// Check position is within bounds
			maze, guardPosition, stop, _, guardDirection = MoveGuard(maze, guardPosition, newPos, guardDirection)
			if stop {
				stop = true
				break
			}
			if probeInfiniteLoop {
				// If current is '+' and we were to spawn an object in front of us
				// Would the next pos be a SEEN pos? If yes, the guard will likely end in a loop
				newPos = [2]int{guardPosition[0] + DIR_UP[0], guardPosition[1] + DIR_UP[1]}
				newSquare := maze[newPos[0]][newPos[1]]
				switch newSquare {
				case SEEN2, SEEN3:
					mazeCopy := deepCopyMatrix(maze)
					// Put the new obstacle
					mazeCopy[guardPosition[0]+DIR_LEFT[0]][guardPosition[1]+DIR_LEFT[1]] = VORTEX
					// Move guard
					mazeCopy[newPos[0]][newPos[1]] = UP
					infiniteLoop, _ := WalkThePlantV2(mazeCopy, 1000, false)
					if infiniteLoop {
						fmt.Println("Infinit loop detected with maze")
						PrintMaze(mazeCopy)
					}
				}
			}

			// PrintMaze(maze)
		}
		count++
		if maxSteps != -1 && count >= maxSteps {
			return true, -1
		}
	}
	PrintMaze(maze)
	return false, CountSeenV2(maze)
}

func deepCopyMatrix(matrix [][]rune) [][]rune {
	// Create a new slice for the outer slice
	copied := make([][]rune, len(matrix))

	for i, row := range matrix {
		// Create a new slice for each inner slice
		copied[i] = make([]rune, len(row))
		copy(copied[i], row) // Use copy to copy the contents
	}
	return copied
}

func SeenGlyph(guardDirection, tile rune) rune {
	switch guardDirection {
	case UP, DOWN:
		if tile == '-' {
			return '+'
		}
		return '|'
	case LEFT, RIGHT:
		if tile == '|' {
			return '+'
		}
		return '-'
	default:
		panic(fmt.Sprintf("unexpected guardDirection: %s", string(guardDirection)))
	}
}

func GetGuardPosition(maze [][]rune) [2]int {
	guardPosition := [2]int{}
	for rowI, row := range maze {
		for colI, char := range row {
			if char == UP || char == DOWN || char == LEFT || char == RIGHT {
				guardPosition[0], guardPosition[1] = rowI, colI
				break
			}
		}
	}
	return guardPosition
}

func PrintMaze(m [][]rune) {
	for _, row := range m {
		fmt.Println(string(row))
	}
	fmt.Println("")
}

func CountSeen(maze [][]rune) int {
	c := 0
	for _, row := range maze {
		for _, char := range row {
			if char == '0' {
				c++
			}
		}
	}
	return c
}

func CountSeenV2(maze [][]rune) int {
	c := 0
	for _, row := range maze {
		for _, char := range row {
			if char == '|' || char == '-' || char == '+' {
				c++
			}
		}
	}
	return c
}

// Return the content of a file as a string
func ReadFile(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
