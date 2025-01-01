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

func FormatInput(input string) [][]rune {
	inputSplit := strings.Split(input, "\n")
	out := make([][]rune, 0)
	for _, line := range inputSplit {
		out = append(out, []rune(line))
	}
	return out
}
