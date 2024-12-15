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
	niceGrid := g.NewGrid(utils.ConvertInput(content))
	_, seen := niceGrid.StartPatrol(-1)
	evilGrid := g.NewGrid(utils.ConvertInput(content))
	evil := evilGrid.StartEvilPatrol(seen)

	fmt.Printf("The guard visited %d unique tiles on the nice grid\n", niceGrid.CountVisitedTiles())
	fmt.Printf("The guard can be trapped with a new obstacle on %d unique tiles on the evil grid\n", evil)
}
