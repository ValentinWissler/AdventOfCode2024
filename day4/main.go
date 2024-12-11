package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	INPUT_FILE = "./crossword.txt" // Puzzle input
	TEST_FILE  = "./mini.txt"      // To test if the code is working, the mini contains 18 XMAS
)

var Directions = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

func main() {
	puzzle, err := ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c := SearchPuzzle(puzzle, "XMAS")
	fmt.Println("The count of XMAS is ", c)
	c2 := SearchPuzzleV2(puzzle)
	fmt.Println("The count of X-MAS is ", c2)
}

// Return how many times word appears in input, see Directions for possible moves
func SearchPuzzle(input, word string) int {
	puzzle := strings.Split(input, "\n")
	for x, p := range puzzle {
		puzzle[x] = strings.TrimSpace(p)
	}
	count := 0
	for rowIndex := 0; rowIndex < len(puzzle); rowIndex++ {
		for colIndex := 0; colIndex < len(puzzle[0]); colIndex++ {
			count += FindWordAt(puzzle, word, rowIndex, colIndex)
		}
	}
	return count
}

func SearchPuzzleV2(input string) int {
	puzzle := strings.Split(input, "\n")
	for x, p := range puzzle {
		puzzle[x] = strings.TrimSpace(p)
	}
	count := 0
	for rowIndex := 1; rowIndex < len(puzzle)-1; rowIndex++ {
		for colIndex := 1; colIndex < len(puzzle[0])-1; colIndex++ {
			// Search diagonaly for MAS or SAM
			if puzzle[rowIndex][colIndex] == 'A' {
				// Now that the moves are in bound, see if we can find the word
				if ((puzzle[rowIndex-1][colIndex-1] == 'M' && puzzle[rowIndex+1][colIndex+1] == 'S') || (puzzle[rowIndex-1][colIndex-1] == 'S' && puzzle[rowIndex+1][colIndex+1] == 'M')) && ((puzzle[rowIndex-1][colIndex+1] == 'S' && puzzle[rowIndex+1][colIndex-1] == 'M') || (puzzle[rowIndex-1][colIndex+1] == 'M' && puzzle[rowIndex+1][colIndex-1] == 'S')) {
					count++
				}
			}
		}
	}
	return count
}

func FindWordAt(puzzle []string, word string, begCi, begRi int) int {
	count := 0
	// Check if moving in a direction is within bound
	for _, dir := range Directions {
		found := ""
		c, r := begCi, begRi
		for i := 0; i < len(word); i++ {
			// If the move is within bounds
			if (c >= 0 && c < len(puzzle)) && (r >= 0 && r < len(puzzle[0])) {
				found += string(puzzle[c][r])
			} else {
				break
			}
			// Move cursor
			c, r = c+dir[0], r+dir[1]
		}
		if found == word {
			count++
		}
	}
	return count
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
