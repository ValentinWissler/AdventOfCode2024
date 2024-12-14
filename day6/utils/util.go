package utils

import (
	"io"
	"os"
	"strings"
)

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

// Convert the raw string in a slice of rune
func ConvertInput(input string) [][]rune {
	out := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		out = append(out, []rune(strings.TrimSpace(line)))
	}
	return out
}

// Return a deep copy of provided input slice
func DeepCopyMatrix(input [][]rune) [][]rune {
	copy_ := make([][]rune, len(input))
	for i, row := range input {
		copy_[i] = make([]rune, len(row))
		copy(copy_[i], row)
	}
	return copy_
}

// Return all the different combinations of grid where we add a single obstacle in the grid
func FindAllGrids(input [][]rune) [][][]rune {
	grids := make([][][]rune, 0)
	for ri, row := range input {
		for ci, char := range row {
			if char != '#' && char != '^' {
				newGrid := DeepCopyMatrix(input)
				newGrid[ri][ci] = '0'
				grids = append(grids, newGrid)
			}
		}
	}
	return grids
}
