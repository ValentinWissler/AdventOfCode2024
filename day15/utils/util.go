package utils

import (
	"io"
	"log"
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
	// fmt.Println(split)
	return GetGrid(split[0]), GetCommands(split[1])
}

func EnlargeGrid(input string) [][]rune {
	split := SplitInput(input)
	grid := make([][]rune, 0)
	for _, r := range strings.Split(split[0], "\n") {
		r = strings.ReplaceAll(r, "O", "[]")
		r = strings.ReplaceAll(r, ".", "..")
		r = strings.ReplaceAll(r, "#", "##")
		r = strings.ReplaceAll(r, "@", "@.")
		grid = append(grid, []rune(r))
	}
	return grid
}

func LogToHistory(grid string, command rune, turn int) {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %s", err)
	}
	defer file.Close()

	// Set the output of the standard logger to the file
	log.SetOutput(file)

	log.Printf("Turn %d\n%s\nCommand %s\n", turn, grid, string(command))
}
