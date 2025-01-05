package main

import (
	cmp "aoc17/computer"
	"aoc17/utils"
	"fmt"
	"testing"
)

// Create a test function that will test a maze

func TestMaze(t *testing.T) {

	in, err := utils.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	a, b, c, prg := utils.ConvertInput(in)
	computer := cmp.NewComputer(a, b, c, prg...)
	fmt.Printf("Computer output: %s\n", computer.ProcessCmds())
}
