package utils

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	buttonFinder = regexp.MustCompile(`X\+(\d+),\sY\+(\d+)`)
	goalFinder   = regexp.MustCompile(`X=(\d+),\sY=(\d+)`)
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

func IngestInput(input string) [][6]int {
	out := make([][]string, 0)
	inputSplit := strings.Split(input, "\n")
	inputSplit2 := make([]string, 0)
	for _, v := range inputSplit {
		if v != "" {
			inputSplit2 = append(inputSplit2, v)
		}
	}
	for _, v := range inputSplit2 {
		fmt.Println(v)
	}
	for x := 0; x < len(inputSplit2); x += 3 {
		out = append(out, []string{inputSplit2[x], inputSplit2[x+1], inputSplit2[x+2]})
	}

	exps := make([][6]int, 0)
	for _, exp := range out {
		formated := [6]int{}
		for x, v := range exp {
			if x < 2 {
				matches := buttonFinder.FindStringSubmatch(v)
				formated[x*2], _ = strconv.Atoi(matches[1])
				formated[x*2+1], _ = strconv.Atoi(matches[2])
			} else {
				matches := goalFinder.FindStringSubmatch(v)
				formated[x*2], _ = strconv.Atoi(matches[1])
				formated[x*2+1], _ = strconv.Atoi(matches[2])
			}
		}
		exps = append(exps, formated)
	}

	return exps
}
