package main

import (
	"aoc11/utils"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const (
	TEST_FILE  = "../inputs/mini.txt"
	INPUT_FILE = "../inputs/input.txt"
	PINK       = "\033[35m"
	DEFAULT    = "\033[0m"
)

type T struct {
	stone     int64
	numBlinks int
}

var Cache = make(map[T]big.Int)

func main() {
	max := 25
	if len(os.Args) >= 2 {
		n, _ := strconv.Atoi(os.Args[1])
		max = n
	}
	in, err := utils.ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	stones := utils.ConvertInput(in)
	sum := big.Int{}
	for _, stone := range stones {
		numStones := hellRecurseCached(int64(stone), 0, max)
		sum = *sum.Add(&sum, &numStones)
	}
	fmt.Printf("For the starting sequence: %v, after %d blinks there will be %d stones\n", stones, max, sum.Add(&sum, big.NewInt(int64(len(stones)))))
}

// Returns the number of stones starting from curr after max turn
func hellRecurseCached(currStone int64, currT, maxT int) big.Int {
	add := big.Int{}
	if currT == maxT {
		return big.Int{}
	}
	state := []int64{}
	t := T{stone: int64(currStone), numBlinks: maxT - currT}
	v, found := Cache[t]
	if found {
		add = *v.Add(&add, &v)
	} else {
		split := strconv.Itoa(int(currStone))
		if currStone == 0 {
			state = append(state, 1)
		} else if len(split)%2 == 0 {
			a, _ := strconv.Atoi(split[:(len(split) / 2)])
			b, _ := strconv.Atoi(split[(len(split) / 2):])
			state = append(state, int64(a), int64(b))
			add = *add.Add(&add, big.NewInt(1))
		} else {
			n := currStone * 2024
			state = append(state, int64(n))
		}
	}

	// We recurse from this state
	for _, stone := range state {
		res := hellRecurseCached(stone, currT+1, maxT)
		add = *v.Add(&res, &add)
	}
	Cache[t] = add
	return add
}
