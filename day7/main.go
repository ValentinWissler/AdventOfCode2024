package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	MAIN_FILE = "./input.txt"
	TEST_FILE = "./mini.txt"
)

func main() {
	content, err := ReadFile(MAIN_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	equations := formatInputToEquations(content)
	// PrintEquations(equations)
	sum := 0
	for _, eq := range equations {
		for _, signs := range eq.signList {
			opsCopy := make([]int, len(eq.ops))
			copy(opsCopy, eq.ops)
			res := 0
			for x, symbol := range signs {
				a, b := opsCopy[0], opsCopy[1]
				c := 0
				if symbol == '+' {
					c = a + b
				} else if symbol == '*' {
					c = a * b
				} else if symbol == '|' {
					cs := strconv.Itoa(a) + strconv.Itoa(b)
					c, _ = strconv.Atoi(cs)
				}
				res = c
				if x+1 < len(signs) {
					opsCopy = append([]int{c}, opsCopy[2:]...)
				}
			}
			if res == eq.res {
				sum += res
				break
			}
		}
	}
	fmt.Printf("The sum of valid equation %d", sum)
}

func formatInputToEquations(input string) []equation {
	// Format the input
	splitInput := strings.Split(input, "\n")
	equations := []equation{}
	for _, line := range splitInput {
		parts := strings.Fields(strings.Join(strings.Split(line, ":"), " "))
		numbers := convertStrSliceToIntSlice(parts)
		equations = append(equations, NewEquation(numbers[0], numbers[1:]...))
	}
	return equations
}

func convertStrSliceToIntSlice(str []string) []int {
	row := []int{}
	for _, op := range str {
		v, err := strconv.Atoi(op)
		if err != nil {
			panic(fmt.Sprintf("error converting %s to int", op))
		}
		row = append(row, v)
	}
	return row
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

const (
	PLUS = iota
	MUL
	PIPE
)

var Sign = []rune{'+', '*', '|'}

type equation struct {
	res      int
	ops      []int
	signList [][]rune // Possible values are: '+' or '*'
}

func NewEquation(res int, ops ...int) equation {
	e := equation{res: res, ops: ops}
	e.generateSignList()
	return e
}

func (e equation) String() string {
	return fmt.Sprintf("Res %d, operands %v\nSigns\n%s", e.res, e.ops, e.printSigns())
}

func (e equation) PrintEquation() {
	fmt.Println(e.String())
}

func (e equation) printSigns() string {
	txt := ""
	for _, row := range e.signList {
		txt += fmt.Sprintln(string(row))
	}
	return txt
}

func PrintEquations(eq []equation) {
	for _, e := range eq {
		e.PrintEquation()
	}
}

func (e *equation) generateSignList() {
	numOfSigns := len(e.ops) - 1
	e.signList = generateCombinations(numOfSigns)
}

// generateCombinations generates all unique combinations of + and * arrays of length N
func generateCombinations(N int) [][]rune {
	if N <= 0 {
		return [][]rune{}
	}

	// Initialize the result array
	result := [][]rune{}

	// Recursive function to build the combinations
	var backtrack func(current []rune)
	backtrack = func(current []rune) {
		// Base case: if the current combination length equals N, add it to the result
		if len(current) == N {
			// Make a copy of current to avoid reference issues
			combination := make([]rune, N)
			copy(combination, current)
			result = append(result, combination)
			return
		}

		// Add 0 and recurse
		current = append(current, Sign[PLUS])
		backtrack(current)
		current = current[:len(current)-1] // Backtrack

		// Add 1 and recurse
		current = append(current, Sign[MUL])
		backtrack(current)
		current = current[:len(current)-1] // Backtrack

		// Add 2 and recurse
		current = append(current, Sign[PIPE])
		backtrack(current)
		current = current[:len(current)-1] // Backtrack
	}

	// Start the recursion with an empty combination
	backtrack([]rune{})

	return result
}
