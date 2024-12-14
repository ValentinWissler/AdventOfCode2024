package main

import (
	g "aoc6/grid"
	"aoc6/utils"
	"fmt"
	"testing"
)

// Create a test function that will test a maze

func TestMaze(t *testing.T) {
	content, err := utils.ReadFile(TEST_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	maze := utils.ConvertInput(content)
	// niceGrid := g.NewGrid(maze)
	// evilGrid := g.NewGrid(maze)

	grids := utils.FindAllGrids(maze)
	for _, grid := range grids {
		ng := g.NewGrid(grid)
		fmt.Println(ng.PrintMaze())
	}
}
