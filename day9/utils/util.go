package utils

import (
	"fmt"
	"io"
	"os"
	"strconv"
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
	end := make([]int, 0)
	id := 0
	for x, char := range input {
		// This is a file block
		if x%2 == 0 {
			numBlock, err := strconv.Atoi(string(char))
			if err != nil {
				panic(fmt.Sprintf("err converting %s to int", string(char)))
			}
			for i := 0; i < numBlock; i++ {
				end = append(end, id)
			}
			id++
		} else {
			numSpace, err := strconv.Atoi(string(char))
			if err != nil {
				panic(fmt.Sprintf("err converting %s to int", string(char)))
			}
			for i := 0; i < numSpace; i++ {
				end = append(end, -1)
			}
		}
	}
	return end
}
