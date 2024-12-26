package utils

import (
	"io"
	"os"
	"strconv"
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

func ConvertInput(input string) []int {
	out := make([]int, 0)
	inputSplit := strings.Fields(input)
	for _, n := range inputSplit {
		ns, _ := strconv.Atoi(n)
		out = append(out, ns)
	}
	return out
}
