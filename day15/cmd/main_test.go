package main

import (
	g "aoc15/grid"
	"aoc15/utils"
	"fmt"
	"testing"
)

// Create a test function that will test a maze

func TestMaze(t *testing.T) {
	in, err := utils.ReadFile(TEST_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Part 1
	gridInput, commsInput := utils.ConvertInput(in)
	grid := g.NewGrid(gridInput, commsInput)
	sum := grid.ProcessCommands()
	fmt.Printf("The sum of gps coordinates is %d\n", sum)

	// Part 2
	gridInput2 := utils.EnlargeGrid(in)
	grid2 := g.NewGrid(gridInput2, commsInput)
	// fmt.Println(grid2.String())
	sum2 := grid2.ProcessCommandsV2()
	fmt.Printf("The sum of gps coordinates is %d\n", sum2)
}
