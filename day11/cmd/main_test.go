package main

import (
	"aoc11/utils"
	"fmt"
	"math/big"
	"testing"
)

// // Create a test function that will test a maze

func TestMaze(t *testing.T) {
	in, err := utils.ReadFile(TEST_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	stones := utils.ConvertInput(in)
	fmt.Println(stones)
	sum := big.Int{}
	for _, stone := range stones {
		numStones := hellRecurseCached(int64(stone), 0, 3)
		sum = *sum.Add(&sum, &numStones)
	}
	fmt.Println(sum.Add(&sum, big.NewInt(int64(len(stones)))))
}
