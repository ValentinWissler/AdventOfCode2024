package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	INPUT_FILE = "./list.txt"
	UP         = iota
	DOWN
)

func main() {
	fileContent, err := ReadFile(INPUT_FILE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	list := FormatList(fileContent)
	safeReports := FindSafeReports(list, false)
	fmt.Printf("Num of safe reports: %d\n", safeReports)
	safeEditedReports := FindSafeReports(list, true)
	fmt.Printf("Num of safe reports after edition: %d\n", safeEditedReports)
}

// returns the number of safe reports
func FindSafeReports(reports [][]int, recurse bool) int {
	count := 0
	for _, report := range reports {
		if recurse {
			if IsReportSafe(report, true) {
				count++
			}
		} else {
			if IsReportSafe(report, false) {
				count++
			}
		}
	}
	return count
}

/*
7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
*/
func IsReportSafe(report []int, retry bool) bool {
	// Determine an initial direction, whichever pair we use, we keep this as a baseline direction
	direction := -1
	// Loop through the report, make sure the direction is constant and that the delta is between 1 and 3 included
	for i := 0; i < len(report)-1; i++ {
		delta := report[i] - report[i+1]

		// Set the inital direction
		if i == 0 {
			if delta > 0 { // We are decreasing and thus going down
				direction = DOWN
			} else if delta < 0 { // We are increasing and thus going up
				direction = UP
			} else { // The level are equal and thus unsafe
				if retry {
					newReps := OtherReports(report)
					for _, rep := range newReps {
						if IsReportSafe(rep, false) {
							return true
						}
					}
				}
				return false
			}
		}

		if (delta > 0 && delta < 4 && direction == DOWN) || (delta < 0 && delta > -4 && direction == UP) {
			continue
		} else { // The levels are unsafe
			if retry { // retry by editing the report and return all possible version of the report
				newReps := OtherReports(report)
				for _, rep := range newReps {
					if IsReportSafe(rep, false) {
						return true
					}
				}
			}
			return false
		}
	}
	return true
}

// Bruteforce the report mod by creating every mod of a report
func OtherReports(report []int) [][]int {
	newReports := make([][]int, 0)
	for x := 0; x < len(report); x++ {
		newRep := make([]int, 0)
		for i := 0; i < len(report); i++ {
			if i != x {
				newRep = append(newRep, report[i])
			}
		}
		newReports = append(newReports, newRep)
	}
	return newReports
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

// Format the list in a slice of slice of ints
func FormatList(list string) [][]int {
	rows := strings.Split(list, "\n")
	fList := make([][]int, 0)
	for _, r := range rows {
		row := StringSliceToIntSlice(strings.Fields(r))
		fList = append(fList, row)
		// fmt.Printf("Extracted row: %v\n", row)
	}
	return fList
}

func StringSliceToIntSlice(l []string) []int {
	slice := make([]int, 0)
	for _, n := range l {
		nInt, err := strconv.Atoi(n)
		if err != nil {
			panic(fmt.Sprintf("error converting %s to int", n))
		}
		slice = append(slice, nInt)
	}
	return slice
}
