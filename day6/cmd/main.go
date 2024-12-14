package main

import (
	g "aoc6/grid"
	"aoc6/utils"
	"fmt"
)

const (
	INPUT_FILE = "../inputs/input.txt" // Puzzle input
	TEST_FILE  = "../inputs/mini.txt"  // To test if the code is working, the mini contains 18 XMAS#
)

func main() {
	content, err := utils.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	maze := utils.ConvertInput(content)
	// niceGrid := g.NewGrid(maze)

	count := 0
	evilGrids := utils.FindAllGrids(maze)
	for _, grid := range evilGrids {
		ng := g.NewGrid(grid)
		isInfinite := ng.StartPatrol(7000)
		if isInfinite {
			count++
		}
	}
	// fmt.Println(count)

	// niceGrid.StartPatrol(-1)
	// fmt.Printf("The guard visited %d unique tiles on the nice grid\n", niceGrid.CountVisitedTiles())
}
