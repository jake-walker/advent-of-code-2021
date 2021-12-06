package main

import (
	"fmt"
	"github.com/jake-walker/advent-of-code-2021/helpers"
	"strings"
)

func CalculateFish(input []int) []int {
	newFish := []int{}

	for i := 0; i < len(input); i++ {
		if input[i] <= 0 {
			input[i] = 6
			newFish = append(newFish, 8)
			continue
		}

		input[i] = input[i] - 1
	}

	return append(input, newFish...)
}

func CalculateFishLoop(fish []int, days int) []int {
	for i := 0; i < days; i++ {
		fish = CalculateFish(fish)
	}
	return fish
}

func main() {
	input := helpers.StringSliceToIntSlice(strings.Split(helpers.GetInput("day6/input.txt"), ","))

	part1 := CalculateFishLoop(input, 80)
	fmt.Printf("part 1: %v\n", len(part1))

	part2 := CalculateFishLoop(input, 256)
	fmt.Printf("part 2: %v\n", len(part2))
}
