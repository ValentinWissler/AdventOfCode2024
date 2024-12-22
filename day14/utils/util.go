package utils

import (
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	posFinder = regexp.MustCompile(`p=(-?\d+),(-?\d+)\sv=(-?\d+),(-?\d+)`)
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

func IngestInput(input string) [][4]int {
	inputSplit := strings.Split(input, "\n")
	// for _, v := range inputSplit {
	// 	fmt.Println(v)
	// }
	exps := make([][4]int, 0)
	for _, exp := range inputSplit {
		formated := [4]int{}
		matches := posFinder.FindStringSubmatch(exp)
		// fmt.Println(matches)
		if len(matches) != 5 {
			continue
		}
		for i := 1; i < 5; i++ {
			val, err := strconv.Atoi(matches[i])
			if err != nil {
				continue
			}
			formated[i-1] = val
		}
		exps = append(exps, formated)
	}
	return exps
}
