package utils

import (
	"fmt"
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

func GetGrid(in string) [][]rune {
	grid := make([][]rune, 0)
	for _, r := range strings.Split(in, "\n") {
		grid = append(grid, []rune(r))
	}
	return grid
}

func GetCommands(in string) []rune {
	return []rune(strings.Join(strings.Split(in, "\r\n"), ""))
}

func SplitInput(input string) []string {
	end := make([]string, 0)
	for _, line := range strings.Split(input, "\r\n\r\n") {
		end = append(end, strings.TrimSpace(line))
	}
	return end
}

func ConvertInput(input string) (grid [][]rune, commands []rune) {
	split := SplitInput(input)
	fmt.Println(split)
	return GetGrid(split[0]), GetCommands(split[1])
}

func EnlargeGrid(input string) [][]rune {
	grid := make([][]rune, 0)
	for _, r := range strings.Split(input, "\n") {
		r = strings.ReplaceAll(r, "O", "[]")
		r = strings.ReplaceAll(r, ".", "..")
		r = strings.ReplaceAll(r, "#", "##")
		r = strings.ReplaceAll(r, "@", "@.")
		grid = append(grid, []rune(r))
	}
	return grid
}
