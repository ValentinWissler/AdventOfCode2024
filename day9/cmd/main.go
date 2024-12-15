package main

import (
	"aoc9/utils"
	"fmt"
	"slices"
)

const (
	TEST_FILE  = "../inputs/mini.txt"
	INPUT_FILE = "../inputs/input.txt"
)

func main() {
	in, err := utils.ReadFile(INPUT_FILE)
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
	// fmt.Printf("%v\n", out)
	fmt.Printf("Checksum V2 %d\n", Checksum(out))
}

func Checksum(input []int) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		n := input[i]
		if n == -1 {
			continue
		} else {
			sum += n * i
		}
	}
	return sum
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func CompactFSV2(input []int) []int {
	blocks := FindBlocks(input)
	slices.Reverse(blocks)

	for _, block := range blocks {
		spaces := FindSpaces(input, block.Index)
		for _, space := range spaces {
			if space[0] >= block.Index {
				break
			} else if block.Length <= space[1] {
				// Shift block to free space
				for shift, remove := space[0], block.Index; shift < space[0]+block.Length; shift, remove = shift+1, remove+1 {
					input[shift] = block.ID
					input[remove] = -1
				}
				break
			}
		}
	}
	return input
}

func FindNextSpace(input []int, beg, end int) (int, int) {
	for input[beg] != -1 {
		beg++
		if beg > len(input)-1 {
			break
		}
	}
	// Find the length of this free space
	space := 0
	for i := beg; i < end; i++ {
		if input[i] == -1 {
			space++
		} else {
			break
		}
	}
	return beg, space
}

func FindNextBlock(input []int, beg, end int) (int, int, int) {
	for input[end] == -1 {
		end--
		if end < 0 {
			break
		}
	}
	// Find data block len
	id := input[end]
	blockLen := 0
	for i := end; i > beg; i-- {
		if input[i] == id {
			blockLen++
		} else {
			break
		}
	}
	return end, blockLen, id
}

type Block struct {
	Index  int
	Length int
	ID     int
}

func FindBlocks(input []int) []Block {
	blocks := []Block{}
	prev := input[0]
	startIndex := 0
	length := 1
	for i := 1; i < len(input); i++ {
		curr := input[i]
		if curr == -1 {
			if prev != -1 {
				blocks = append(blocks, Block{Index: startIndex, Length: length, ID: prev})
				length = 0
				prev = input[i]
			} else {
				continue
			}
		} else if curr != prev {
			if prev != -1 {
				blocks = append(blocks, Block{Index: startIndex, Length: length, ID: prev})
			}
			startIndex = i
			prev = curr
			length = 1
		} else if curr == prev {
			length++
		}
		if i+1 == len(input) {
			blocks = append(blocks, Block{Index: startIndex, Length: length, ID: prev})
		}
	}
	return blocks
}

func FindSpaces(input []int, limit int) [][2]int {
	spaces := [][2]int{}
	prev := input[0]
	startIndex := 0
	length := 0
	for i := 1; i < limit; i++ {
		curr := input[i]
		if curr == -1 {
			if prev != -1 {
				startIndex = i
			}
			length++
			prev = curr
		} else if curr != -1 {
			if prev != -1 {
				continue
			} else {
				spaces = append(spaces, [2]int{startIndex, length})
				prev = curr
				length = 0
			}
		}

	}
	return spaces
}

func CompactFS(input []int) []int {
	// Create two cursors
	beg, end := 0, len(input)-1
	for beg < end {
		// Find the next free space with beg
		for input[beg] != -1 {
			beg++
			if beg > len(input)-1 {
				break
			}
		}
		// Find the next block of data with end
		for input[end] == -1 {
			end--
			if end < 0 {
				break
			}
		}
		if beg >= end {
			break
		}
		// Swap beg and end values
		input[beg], input[end] = input[end], input[beg]

		// For debugging
		// fmt.Println(string(input))
	}
	return input
}
