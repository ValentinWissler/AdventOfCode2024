package main

import (
	"aoc9/utils"
	"fmt"
)

type Pos struct {
	x, y int
}

const (
	TEST_FILE  = "../inputs/mini.txt"
	TEST2_FILE = "../inputs/minii.txt"
	INPUT_FILE = "../inputs/input.txt"
	PINK       = "\033[35m"
	DEFAULT    = "\033[0m"
)

func main() {
	in, err := utils.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	map_ := utils.ConvertInput(in)
	heads := findTrailHeads(map_)
	// showTrailHeads(heads, map_)

	summits := 0
	for _, head := range heads {
		summits += findSummits(head, map[Pos]bool{}, map_, true)
	}
	fmt.Printf("All trails score is %d\n", summits)

	summitsV2 := 0
	for _, head := range heads {
		summitsV2 += findSummits(head, map[Pos]bool{}, map_, false)
	}
	fmt.Printf("All distinct trail score is %d\n", summitsV2)
}

func findTrailHeads(map_ [][]int) []Pos {
	trailHeads := []Pos{}
	for y, row := range map_ {
		for x, n := range row {
			if n == 0 {
				trailHeads = append(trailHeads, Pos{x: x, y: y})
			}
		}
	}
	return trailHeads
}

func showTrailHeads(pos []Pos, map_ [][]int) {
	posI := 0
	for y, row := range map_ {
		for x, n := range row {
			if pos[posI].x == x && pos[posI].y == y {
				fmt.Printf("%s%d%s", PINK, n, DEFAULT)
				posI++
				if posI >= len(pos) {
					posI = 0
				}
			} else {
				fmt.Print(n)
			}
		}
		fmt.Println()
	}
}

var DIRS = []Pos{{y: -1, x: 0}, {y: 0, x: 1}, {y: 1, x: 0}, {y: 0, x: -1}}

func findSummits(cp Pos, prevs map[Pos]bool, map_ [][]int, v1 bool) (count int) {
	// To prevent recursion from taking multiple time the same path
	if v1 {
		if _, beenHere := prevs[cp]; beenHere {
			return 0
		}
	}

	// We've been here now
	prevs[cp] = true

	// The current elevation
	curr := map_[cp.y][cp.x]

	// If we've reached the summit
	if curr == 9 {
		return 1
	}

	// If not, explore each direction of the trail
	// where the elevation is +1
	for _, dir := range DIRS {
		nPos := Pos{y: cp.y + dir.y, x: cp.x + dir.x}
		if isPosInBound(nPos, map_) {
			n := map_[nPos.y][nPos.x]
			if n-curr == 1 {
				if v1 {
					count += findSummits(nPos, prevs, map_, true)
				} else {
					count += findSummits(nPos, prevs, map_, false)
				}

			}
		}
	}
	return count
}

func isPosInBound(pos Pos, grid [][]int) bool {
	if (pos.y >= 0 && pos.y < len(grid)) && (pos.x >= 0 && pos.x < len(grid[0])) {
		return true
	}
	return false
}
