package main

import (
	"fmt"
	"github.com/jake-walker/advent-of-code-2021/helpers"
	"log"
	"strings"
)

func LoadInput(input []string) ([]string, map[string]string) {
	pairInsertions := make(map[string]string)

	for i := 1; i < len(input); i++ {
		if input[i] == "" {
			continue
		}

		fields := strings.Fields(input[i])
		pairInsertions[fields[0]] = fields[2]
	}

	return strings.Split(input[0], ""), pairInsertions
}

func InsertPairs(template []string, pairs map[string]string) []string {
	for i := 0; i < len(template)-1; i = i + 2 {
		pair := template[i : i+2]

		toInsert, found := pairs[strings.Join(pair, "")]
		if !found {
			log.Fatalf("could not find insertion for pair %v", pair)
		}

		template = append(template[:i+2], template[i+1:]...)
		template[i+1] = toInsert
	}

	return template
}

func InsertLoop(template []string, pairs map[string]string, n int) int {
	for i := 0; i < n; i++ {
		log.Printf("running loop %v...", i)
		template = InsertPairs(template, pairs)
	}

	counts := make(map[string]int)
	for _, item := range template {
		counts[item] += 1
	}
	values := make([]int, 0, len(counts))
	for _, value := range counts {
		values = append(values, value)
	}
	min, max := helpers.MinMax(values)

	return max - min
}

func main() {
	input := helpers.GetInputLines("day14/input.txt")
	template, pairs := LoadInput(input)

	part1 := InsertLoop(template, pairs, 10)
	fmt.Printf("part 1: %v\n", part1)

	part2 := InsertLoop(template, pairs, 40)
	fmt.Printf("part 2: %v\n", part2)
}
