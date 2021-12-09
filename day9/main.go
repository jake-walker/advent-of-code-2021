package main

import (
	"fmt"
	"github.com/jake-walker/advent-of-code-2021/helpers"
	"log"
	"strconv"
)

func LoadInput(content string) [][]int {
	lines := helpers.SplitLines(helpers.GetLines(content))
	out := make([][]int, len(lines))

	for i, line := range lines {
		out[i] = make([]int, len(line))
		for j, char := range line {
			n, err := strconv.Atoi(char)
			if err != nil {
				log.Fatalf("failed to convert %v to int: %v", char, err)
			}
			out[i][j] = n
		}
	}

	return out
}

func MinMax(slice []int) (min, max int) {
	min, max = slice[0], slice[0]
	for _, i := range slice {
		if max < i {
			max = i
		}
		if min > i {
			min = i
		}
	}
	return
}

func GetLowPoints(heightmap [][]int) []int {
	out := []int{}

	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			adjacent := []int{}
			if i-1 >= 0 {
				adjacent = append(adjacent, heightmap[i-1][j])
			}
			if i+1 < len(heightmap) {
				adjacent = append(adjacent, heightmap[i+1][j])
			}
			if j-1 >= 0 {
				adjacent = append(adjacent, heightmap[i][j-1])
			}
			if j+1 < len(heightmap[i]) {
				adjacent = append(adjacent, heightmap[i][j+1])
			}

			min, _ := MinMax(adjacent)

			if heightmap[i][j] < min {
				out = append(out, heightmap[i][j])
			}
		}
	}

	return out
}

func GetTotalRiskLevel(points []int) int {
	total := 0
	for _, point := range points {
		total += point + 1
	}
	return total
}

func main() {
	input := helpers.GetInput("day9/input.txt")
	heightmap := LoadInput(input)

	lowPoints := GetLowPoints(heightmap)
	part1 := GetTotalRiskLevel(lowPoints)
	fmt.Printf("part 1: %v\n", part1)
}
