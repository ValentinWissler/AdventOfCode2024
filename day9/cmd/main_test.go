package main

import (
	"aoc9/utils"
	"fmt"
	"testing"
)

// // Create a test function that will test a maze

func TestMaze(t *testing.T) {
	in, err := utils.ReadFile(TEST_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	input := utils.ConvertInput(in)
	copy_ := make([]int, len(input))
	copy(copy_, input)
	// fmt.Printf("Input:\n%v\n", input)
	output := CompactFS(copy_)
	// fmt.Printf("Output:\n%v\n", output)
	fmt.Printf("Checksum: %d\n", Checksum(output))

	out := CompactFSV2(input)
	fmt.Printf("%v\n", out)
}
