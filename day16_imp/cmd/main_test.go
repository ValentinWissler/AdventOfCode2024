package main

import (
	g "aoc16/grid"
	"aoc16/utils"
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
	input := utils.FormatInput(in)
	grid := g.NewGrid(input)
	nodes := grid.FindNodes()
	fmt.Println(grid.PrintNodes(nodes))
	grid.ConnectNodes(nodes)
	// fmt.Print(grid.Print())
}
