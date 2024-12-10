package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	INPUT_FILE = "./day5/input.txt" // Puzzle input
	TEST_FILE  = "./mini.txt"       // To test if the code is working, the mini contains 18 XMAS
)

func main() {
	input, err := ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	before, _, pages := GetRulesAndPages(input)
	vsum := 0
	esum := 0
	for _, page := range pages {
		isValid := true
		for i := 0; i < len(page)-1; i++ {
			curr := page[i]
			// Is curr meant to be before any of the next pages?
			for j := i + 1; j < len(page); j++ {
				next := page[j]
				_, found := before[next][curr]
				if found {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}
		if isValid {
			vsum += page[len(page)/2]
		} else {
			// Reorder the pag and calc its mul, can use bubble sort
			for i := 0; i < len(page)-1; i++ {
				curr, next := page[i], page[i+1]
				_, swapThem := before[next][curr]
				if swapThem {
					page[i], page[i+1] = next, curr
					i = -1 // reset loop
				}
			}
			esum += page[len(page)/2]

		}
	}
	// fmt.Printf("The sum of mul of valid page is %d\n", vsum)
	// fmt.Printf("The sum of mul of reorder page is %d\n", esum)

}

func GetRulesAndPages(input string) (map[int]map[int]bool, map[int]map[int]bool, [][]int) {
	before, after := make(map[int]map[int]bool), make(map[int]map[int]bool)
	pages := make([][]int, 0)
	split := strings.Split(input, "\n")
	for x, line := range split {
		split[x] = strings.TrimSpace(line)
		if !unicode.IsDigit(rune(line[0])) {
			continue
		} else if line[2] == '|' {
			b, errB := strconv.Atoi(line[:2])
			if errB != nil {
				fmt.Printf("Error converting %s to int\n", line[:2])
			}
			a, errA := strconv.Atoi(line[3:5])
			if errA != nil {
				fmt.Printf("Error converting %s to int\n", line[3:])
			}
			if before[b] == nil {
				before[b] = make(map[int]bool)
			}
			if after[a] == nil {
				after[a] = make(map[int]bool)
			}
			before[b][a] = true
			after[a][b] = true
		} else if line[2] == ',' {
			parts := strings.Split(strings.TrimSpace(line), ",")
			page := make([]int, 0)
			for _, p := range parts {
				n, err := strconv.Atoi(p)
				if err != nil {
					fmt.Println("Error converting ", p)
				}
				page = append(page, n)
			}
			pages = append(pages, page)
			// fmt.Printf("Extracted page: %v\n", page)
		}
	}
	return before, after, pages
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
