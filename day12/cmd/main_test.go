package main

import (
	"aoc12/utils"
	"fmt"
	"testing"
)

// // Create a test function that will test a maze

func TestMaze(t *testing.T) {
	in, err := utils.ReadFile(TEST2_FILE)
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
		fmt.Printf("Found region of type %s its pos are %v\nPriceV1 %d\nPriceV2 %d\nNum of sides %d\n", string(region.type_), region.pos, priceV1, priceV2, sides)
		total1 += priceV1
		total2 += priceV2
	}
	fmt.Printf("The total1 price is %d\nThe total2 price is %d\n", total1, total2)
}
