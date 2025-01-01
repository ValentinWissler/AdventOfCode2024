package main

import (
	g "aoc16/grid"
	"aoc16/utils"
	"fmt"
	"sort"
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
	input := utils.FormatInput(in)
	grid := g.NewGrid(input)
	// fmt.Print(grid.Print())
	scores := grid.FindPath(g.NewNode(grid.GetStart()), make(map[g.Pos]bool), g.RIGHT, 0)
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})
	fmt.Printf("The smallest score returned is\n%d\n", scores[0])
}
