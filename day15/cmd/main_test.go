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
	gridInput, commsInput := utils.ConvertInput(in)
	fmt.Printf("This is the grid\n%v\n", gridInput)
	fmt.Printf("This is the commands\n%v\n", commsInput)

	grid := g.NewGrid(gridInput, commsInput)
	sum := grid.ProcessCommands()
	fmt.Printf("The sum of gps coordinates is %d\n", sum)
}
