package main

import (
	"aoc12/utils"
	"fmt"
	"sort"
)

type Pos struct {
	y, x int
}

const (
	TEST_FILE  = "../inputs/mini.txt"
	TEST2_FILE = "../inputs/minii.txt"
	INPUT_FILE = "../inputs/input.txt"
	PINK       = "\033[35m"
	DEFAULT    = "\033[0m"
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

// UP RIGHT DOWN LEFT
var DIRS = []Pos{{y: -1, x: 0}, {y: 0, x: 1}, {y: 1, x: 0}, {y: 0, x: -1}}

func main() {
	in, err := utils.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	garden := utils.ConvertInput(in)
	regions := findRegions(garden)
	total1, total2 := 0, 0
	for _, region := range regions {
		priceV1 := region.area() * region.perimeter(garden)
		sides := region.FindSides(garden)
		priceV2 := region.area() * sides
		// fmt.Printf("Found region of type %s its pos are %v\nPriceV1 %d\nPriceV2 %d\nNum of sides %d\n", string(region.type_), region.pos, priceV1, priceV2, sides)
		total1 += priceV1
		total2 += priceV2
	}
	fmt.Printf("The total1 price is %d\nThe total2 price is %d\n", total1, total2)

}

// Find the number of sides for the region
func (r Region) FindSides(garden [][]rune) int {
	cols := orderbyColumn(r.pos)
	rows := orderbyRow(r.pos)
	toSubtract := 0
	for _, positions := range cols {
		if len(positions) < 2 {
			continue
		}
		emptyL, emptyR := false, false
		prevY := 0
		for i := 0; i < len(positions); i++ {
			pos := positions[i]
			// Check left & right, is there a non friend tile there?
			nPosL, nPosR := Pos{x: pos.x + DIRS[LEFT].x, y: pos.y + DIRS[LEFT].y}, Pos{x: pos.x + DIRS[RIGHT].x, y: pos.y + DIRS[RIGHT].y}
			if isPosInBound(nPosL, len(garden), len(garden[0])) && garden[nPosL.y][nPosL.x] == r.type_ {
				emptyL = false
			} else {
				// We're currently empty, was the previous empty too and contiguous?
				if emptyL && prevY+1 == pos.y {
					toSubtract++
				}
				emptyL = true
			}
			if isPosInBound(nPosR, len(garden), len(garden[0])) && garden[nPosR.y][nPosR.x] == r.type_ {
				emptyR = false
			} else {
				// We're currently empty, was the previous empty too and contiguous?
				if emptyR && prevY+1 == pos.y {
					toSubtract++
				}
				emptyR = true
			}
			prevY = pos.y
		}
	}
	for _, positions := range rows {
		if len(positions) < 2 {
			continue
		}
		emptyU, emptyD := false, false
		prevX := 0
		for i := 0; i < len(positions); i++ {
			pos := positions[i]
			// Check up & down, is there a non friend tile there?
			nPosU, nPosD := Pos{x: pos.x + DIRS[UP].x, y: pos.y + DIRS[UP].y}, Pos{x: pos.x + DIRS[DOWN].x, y: pos.y + DIRS[DOWN].y}
			if isPosInBound(nPosU, len(garden), len(garden[0])) && garden[nPosU.y][nPosU.x] == r.type_ {
				emptyU = false
			} else {
				// We're currently empty, was the previous empty too and contiguous?
				if emptyU && prevX+1 == pos.x {
					toSubtract++
				}
				emptyU = true
			}
			if isPosInBound(nPosD, len(garden), len(garden[0])) && garden[nPosD.y][nPosD.x] == r.type_ {
				emptyD = false
			} else {
				// We're currently empty, was the previous empty too and contiguous?
				if emptyD && prevX+1 == pos.x {
					toSubtract++
				}
				emptyD = true
			}
			prevX = pos.x
		}
	}
	return r.perimeter(garden) - toSubtract
}

func orderbyColumn(positions []Pos) map[int][]Pos {
	cols := make(map[int][]Pos)
	for _, pos := range positions {
		cols[pos.x] = append(cols[pos.x], pos)
	}
	for x := range cols {
		ordered := cols[x]
		sort.Slice(ordered, func(i, j int) bool {
			return ordered[i].y < ordered[j].y
		})
		cols[x] = ordered
	}
	return cols
}

func orderbyRow(positions []Pos) map[int][]Pos {
	rows := make(map[int][]Pos)
	for _, pos := range positions {
		rows[pos.y] = append(rows[pos.y], pos)
	}
	for y := range rows {
		ordered := rows[y]
		sort.Slice(ordered, func(i, j int) bool {
			return ordered[i].x < ordered[j].x
		})
		rows[y] = ordered
	}
	return rows
}

type PlantFinder struct {
	garden [][]rune
}

type Region struct {
	type_ rune
	pos   []Pos
}

func NewRegion(t rune, p []Pos) Region {
	return Region{type_: t, pos: p}
}

// The area is the number of plants in the region
func (r Region) area() int {
	return len(r.pos)
}

func (r Region) perimeter(garden [][]rune) int {
	// Each plant looks around itself and counts the number of neighbouring friends
	// We then subtract the number of friends to 4 and sum that up for each plant
	fences := 0
	for _, pos := range r.pos {
		fence := 4
		for _, dir := range DIRS {
			nPos := Pos{x: pos.x + dir.x, y: pos.y + dir.y}
			if !isPosInBound(nPos, len(garden), len(garden[0])) {
				continue
			}
			if garden[nPos.y][nPos.x] == r.type_ {
				fence--
			}
		}
		fences += fence
	}
	return fences
}

// Find all the neighbouring plants of the same type forming a region
func findRegions(garden [][]rune) []Region {
	visited := make(map[Pos]bool, 0)
	regions := make([]Region, 0)
	pf := &PlantFinder{garden: garden}
	// Start at the top left corner of the garden and left to right, top to bottom
	// try to find a region from this starting pos, make sure the starting pos is not
	// already visited to avoid cataloguing one region multiple times
	for y, row := range garden {
		for x, plant := range row {

			pos := Pos{x: x, y: y}
			// Avoid visited tiles
			if _, found := visited[pos]; found {
				continue
			}

			// Else, find the neighbouring plants of this type
			friends := pf.findFriends(plant, pos, visited)
			regions = append(regions, NewRegion(plant, friends))
		}
	}
	return regions
}

func (pf *PlantFinder) findFriends(type_ rune, currPos Pos, visitedPos map[Pos]bool) (friends []Pos) {
	// This is now a visited position
	visitedPos[currPos] = true
	// This is now a pos part of the friends
	friends = append(friends, currPos)
	// Check all around the plant to see if there are friends
	// if yes recurse in that position, then return the found friends
	for _, dir := range DIRS {
		nPos := Pos{x: currPos.x + dir.x, y: currPos.y + dir.y}
		// Skip outbound positions and check next dir
		if !isPosInBound(nPos, len(pf.garden), len(pf.garden[0])) {
			continue
		}
		switch nPlant := pf.garden[nPos.y][nPos.x]; nPlant {
		case type_:
			// Don't backtrack and recurse, otherwise to infinity and beyond we go
			if _, seenBefore := visitedPos[nPos]; seenBefore {
				continue
			}
			// We can recurse from that tile & append the result to friends
			friends = append(friends, pf.findFriends(type_, nPos, visitedPos)...)
		default:
			continue
		}
	}
	return friends
}

func isPosInBound(pos Pos, maxV, maxH int) bool {
	if (pos.y >= 0 && pos.y < maxV) && (pos.x >= 0 && pos.x < maxH) {
		return true
	}
	return false
}
