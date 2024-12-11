package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

const (
	INPUT_FILE = "./list.txt"
	UP         = iota
	DOWN
)

func main() {
	code, err := ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("The sum of mul is %d\n", FindAllMul(code))
	fmt.Printf("The sum of mul with constraints is %d\n", FindAllMulV2(code))
}

func FindAllMul(code string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(code, -1)
	sum := 0
	for _, match := range matches {
		a, b := match[1], match[2]
		nA, _ := strconv.Atoi(a)
		nB, _ := strconv.Atoi(b)
		sum += nA * nB
	}
	return sum
}

func FindAllMulV2(code string) int {
	newCode := EditList(code)
	mul := FindAllMul(newCode)
	return mul
}

// Removes the instructions that are blocked by donts, we'll keep what's in between a do and its next dont
func EditList(code string) string {
	// Find the position of all DOS and DONTS
	dont := regexp.MustCompile(`don't\(\)`)
	do := regexp.MustCompile(`do\(\)`)
	dosIndicesRange := do.FindAllStringIndex(code, -1)
	dontIndicesRange := dont.FindAllStringIndex(code, -1)
	dos, donts := make([]int, 0), make([]int, 0)

	// There list of dos and dnts aren't equal, we need to equalize them later
	for _, pos := range dosIndicesRange {
		dos = append(dos, pos[1])
	}
	for _, pos := range dontIndicesRange {
		donts = append(donts, pos[1])
	}

	// We want to have a list with the same numbers of dos and donts as to remove the instructions that are blocked
	dosA, dntA := EqualiseIndices(dos, donts)

	// fmt.Printf("List of dos: %v\n", dosA)
	// fmt.Printf("List of donts: %v\n", dntA)

	// Cut the list to remove the parts between don'ts and dos index
	return RemoveInstructions(dntA, dosA, code)
}

func RemoveInstructions(dnts, dos []int, code string) string {
	finalStr := code[:dnts[0]]
	for x := 0; x < len(dos); x++ {
		part := ""
		if x+1 >= len(dos) {
			part = code[dos[x]:]
		} else {
			part = code[dos[x]:dnts[x+1]]
		}
		finalStr += part
	}
	return finalStr
}

func EqualiseIndices(dos, donts []int) (dosF, dontsF []int) {
	docount := 0
	for dntcount := 0; dntcount < len(donts) && docount < len(dos); dntcount++ {
		if donts[dntcount] < dos[docount] {
			if !(len(dosF) > 0) || !(donts[dntcount] < dosF[len(dosF)-1]) {
				dosF = append(dosF, dos[docount])
				dontsF = append(dontsF, donts[dntcount])
				docount++
			}
		} else {
			docount++  // Next do
			dntcount-- // Same dnt
		}
	}
	return
}

func ReadFile(file string) (string, error) {
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
