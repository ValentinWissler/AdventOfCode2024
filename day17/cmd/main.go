package main

import (
	cmp "aoc17/computer"
	"aoc17/utils"
	"fmt"
)

const (
	TEST_FILE  = "../inputs/mini.txt"
	TEST2_FILE = "../inputs/minii.txt"
	INPUT_FILE = "../inputs/input.txt"
)

func main() {
	in, err := utils.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	a, b, c, prg := utils.ConvertInput(in)
	computer := cmp.NewComputer(a, b, c, prg...)
	fmt.Printf("Computer output: %s\n", computer.ProcessCmds())
}
