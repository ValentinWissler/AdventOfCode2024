package main

import (
	g "aoc14/guard"
	"aoc14/utils"
	"fmt"
)

const (
	TEST_FILE  = "../inputs/mini.txt"
	INPUT_FILE = "../inputs/input.txt"
	MAXX       = 103
	MAXV       = 101
)

func main() {
	in, err := utils.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	input := utils.IngestInput(in)
	// for _, v := range input {
	// 	fmt.Println(v)
	// }
	sum := 0
	midX, midY := (MAXX+1)/2, (MAXV+1)/2
	for _, grd := range input {
		guard := g.Newguard(g.Pos{X: grd[0], Y: grd[1]}, g.Pos{X: grd[2], Y: grd[3]}, MAXX, MAXV)
		for x := 0; x < 100; x++ {
			guard.Move()
		}
		// if the guard is not in the middle y or middle x, count it
		if guard.Pos().X != midX && guard.Pos().Y != midY {
			sum++
		}
	}
	fmt.Println(sum)
	// grid := grid.NewGrid(input)
	// grid.FindAntenas()

}
