package main

// import (
// 	"fmt"
// 	"strings"
// 	"testing"
// )

// // Create a test function that will test a maze

// func TestMaze(t *testing.T) {
// 	content, err := ReadFile(TEST_FILE)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	// Format the input
// 	str := strings.Split(content, "\n")
// 	maze := make([][]rune, 0)
// 	for _, m := range str {
// 		maze = append(maze, []rune(strings.TrimSpace(m)))
// 	}

// 	d := CountStepsV2(maze)
// 	fmt.Println("Seen v2: ", d)
// }
