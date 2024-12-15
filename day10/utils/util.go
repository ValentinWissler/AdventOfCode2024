package utils

import (
	"fmt"
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

func ConvertInput(input string) [][]int {
	out := make([][]int, 0)
	inputSplit := strings.Split(input, "\n")
	for _, row := range inputSplit {
		numbers := []int{}
		for _, char := range row {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				panic(fmt.Sprintf("error converting %s to int", string(char)))
			}
			numbers = append(numbers, n)
		}
		out = append(out, numbers)
	}
	return out
}
