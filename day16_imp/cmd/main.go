package main

import (
	g "aoc16/grid"
	"aoc16/utils"
	"fmt"
)

const (
	TEST_FILE  = "../inputs/mini.txt"
	TEST2_FILE = "../inputs/minii.txt"
	INPUT_FILE = "../inputs/input.txt"
)

func main() {
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
