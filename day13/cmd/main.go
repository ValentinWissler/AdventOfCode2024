package main

import (
	"aoc13/utils"
	"fmt"
)

type Pos struct {
	y, x int
}

const (
	TEST_FILE  = "../inputs/mini.txt"
	INPUT_FILE = "../inputs/input.txt"
	PINK       = "\033[35m"
	DEFAULT    = "\033[0m"
)

type ClawMachine struct {
	ax, ay, bx, by, cx, cy int
}

type Exp struct {
	ax, ay, bx, by, i, j int
}

func main() {
	in, err := utils.ReadFile(TEST_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	inp := utils.IngestInput(in)
	for _, v := range inp {
		fmt.Println(v)
	}
	visited := make(map[Exp][2]int)
	sum := 0
	for _, v := range inp {
		cm := ClawMachine{ax: v[0] * 1_000_000, ay: v[1] * 1_000_000, bx: v[2] * 1_000_000, by: v[3] * 1_000_000, cx: v[4] + 10_000_000_000_000, cy: v[5] + 10_000_000_000_000}
		fewestTokens := 10_000_000_000_000
		// Test all A combos
		for i := 1_000_000; i >= 0; i-- {
			// Test all B combos
			for j := 0; j <= 1_000_000; j++ {
				nx, ny := 0, 0
				v, ok := visited[Exp{cm.ax, cm.ay, cm.bx, cm.by, i, j}]
				if ok {
					nx, ny = v[0], v[1]
				} else {
					nx, ny = (cm.ax*i)+(cm.bx*j), (cm.ay*i)+(cm.by*j)
					visited[Exp{cm.ax, cm.ay, cm.bx, cm.by, i, j}] = [2]int{nx, ny}
				}
				if nx == cm.cx && ny == cm.cy {
					a, b := i*3, j
					if a+b < fewestTokens {
						fewestTokens = a + b
					}
				}
			}
		}
		if fewestTokens != 10_000_000_000_000 {
			fmt.Printf("Fewest token to reach the target: %s%d%s\n", PINK, fewestTokens, DEFAULT)
			sum += fewestTokens
		}
	}
	fmt.Printf("Sum of all fewest tokens: %s%d%s\n", PINK, sum, DEFAULT)
}
