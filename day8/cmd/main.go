package main

import (
	"aoc8/grid"
	"aoc8/utils"
	"fmt"
)

const (
	TEST_FILE          = "../inputs/mini.txt"
	TEST_RESOLVED_FILE = "../inputs/miniresolved.txt"
	INPUT_FILE         = "../inputs/input.txt"
)

func main() {
	in, err := utils.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	input := utils.ConvertInput(in)

	grid := grid.NewGrid(input)
	grid.FindAntenas()

	grid.AntenasEmit()
	fmt.Printf("There are %d unique antinode positions\n", grid.CountAntinodes())

	grid.AntenasEmitV2()
	fmt.Printf("There are %d unique echoed antinode positions\n", grid.CountAntinodes())
}
