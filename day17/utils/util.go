package utils

import (
	"fmt"
	"io"
	"log"
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

func GetRegisters(in string) (int, int, int) {
	regs := strings.Split(in, "\r\n")
	register := []string{"A", "B", "C"}
	registers := []int{}
	for x, g := range regs {
		st := strings.Replace(g, fmt.Sprintf("Register %s: ", register[x]), "", -1)
		n, err := strconv.Atoi(st)
		if err != nil {
			panic(fmt.Sprint(err.Error()))
		}
		registers = append(registers, n)
	}
	return registers[0], registers[1], registers[2]
}

func GetInstructions(in string) []string {
	cur := strings.Replace(in, "Program: ", "", -1)
	cur = strings.Replace(cur, ",", " ", -1)
	st := strings.Fields(cur)
	return st
}

func SplitInput(input string) []string {
	end := make([]string, 0)
	for _, line := range strings.Split(input, "\r\n\r\n") {
		end = append(end, strings.TrimSpace(line))
	}
	return end
}

func ConvertInput(input string) (int, int, int, []string) {
	split := SplitInput(input)
	a, b, c := GetRegisters(split[0])
	return a, b, c, GetInstructions(split[1])
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
