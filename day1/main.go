package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	INPUT_FILE = "./list.txt"
)

func main() {
	// Read the list of IDs
	list, err := ReadList()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Separate the list in half
	a, b := SplitList(list)
	fmt.Printf("The last row is: %d, %d\n", a[len(a)-1], b[len(b)-1])

	sort.Ints(a)
	sort.Ints(b)

	// Measure the total distance between As & Bs
	totalDistance := 0
	for i := 0; i < len(a); i++ {
		delta := int(math.Abs(float64(a[i]) - float64(b[i])))
		totalDistance += delta
	}
	fmt.Printf("The distance between As and Bs is: %d\n", totalDistance)

	// figure out exactly how often each number from the list A appears in the list B
	fmap := FindFrequency(b)
	score := FindSimilarity(a, fmap)
	fmt.Printf("The similarity score is: %d", score)

}

func ReadList() (string, error) {
	f, err := os.Open(INPUT_FILE)
	if err != nil {
		return "", err
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func SplitList(list string) (listA, listB []int) {
	l := strings.Split(list, "\n")
	for _, row := range l {
		parts := strings.Fields(row)
		if len(parts) == 2 {
			nA, errA := strconv.Atoi(parts[0])
			if errA != nil {
				fmt.Printf("Error converting: %s to int", parts[0])
				return nil, nil
			}
			nB, errB := strconv.Atoi(parts[1])
			if errB != nil {
				fmt.Printf("Error converting: %s to int", parts[1])
				return nil, nil
			}
			listA = append(listA, nA)
			listB = append(listB, nB)
		} else {
			fmt.Println("Error while splitting row")
		}
	}
	return
}

// Find frequency of an int in an ordered list
func FindFrequency(list []int) map[int]int {
	if len(list) < 1 {
		return nil
	}
	fmap := make(map[int]int)
	current := list[0]
	count := 1
	for i := 1; i < len(list); i++ {
		if list[i] != current {
			fmap[current] = count
			current = list[i]
			count = 1
		} else if i+1 == len(list) { // add the last one
			fmap[current] = count
		} else {
			count++
		}
	}
	return fmap
}

// Calculate a total similarity score by adding up each number in the left list
// after multiplying it by the number of times that number appears in the right list.
func FindSimilarity(a []int, fmap map[int]int) int {
	score := 0
	for _, n := range a {
		f, found := fmap[n]
		if !found {
			continue
		}
		score += n * f
	}
	return score
}
