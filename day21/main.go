package main

import (
	"fmt"
	"github.com/jake-walker/advent-of-code-2021/helpers"
	"log"
	"strconv"
	"strings"
)

func NewGame(input string) []int {
	lines := helpers.GetLines(input)
	out := make([]int, len(lines))

	for i, line := range lines {
		pos, err := strconv.Atoi(strings.Split(line, ": ")[1])
		if err != nil {
			log.Fatalf("failed to convert position: %v", err)
		}

		out[i] = pos
	}

	return out
}

func PlayDeterministicDie(startPositions []int) int {
	turn := 0
	positions := startPositions
	scores := make([]int, len(startPositions))
	dieValue := 1

	for {
		turn += 1
		player := (turn + 1) % len(positions)

		roll := dieValue + (dieValue + 1) + (dieValue + 2)
		dieValue = (dieValue + 3) % 100

		positions[player] = ((positions[player] + roll - 1) % 10) + 1
		scores[player] += positions[player]

		//log.Printf("Player %v rolls %v and moves to space %v for a total score of %v", player+1, roll, positions[player], scores[player])

		if scores[player] >= 1000 {
			break
		}
	}

	// 3 rolls per turn * losing score
	return (turn * 3) * (scores[(turn)%len(positions)])
}

func PlayQuantumDie(positions []int, scores []int, wins []int, turn int) {
	turn += 1
	player := (turn + 1) % len(positions)

	// min: 1+1+1=3, max: 3+3+3=9
	for roll := 3; roll <= 9; roll++ {
		newPositions := make([]int, len(positions))
		newScores := make([]int, len(scores))
		copy(newPositions, positions)
		copy(newScores, scores)

		newPositions[player] = ((newPositions[player] + roll - 1) % 10) + 1
		newScores[player] += newPositions[player]

		//log.Printf("%vplayer %v: roll %v, pos %v, score %v", strings.Repeat(" ", turn*2), player+1, roll, newPositions[player], newScores[player])

		if newScores[player] >= 21 {
			//log.Printf("%v-- win --", strings.Repeat(" ", turn*2))
			wins[player] += 1
			continue
		}

		// the original wins is passed so that it can get updated globally
		PlayQuantumDie(newPositions, newScores, wins, turn)
	}
}

func main() {
	startPositions := NewGame(helpers.GetInput("day21/input.txt"))

	part1 := PlayDeterministicDie(startPositions)
	fmt.Printf("part 1: %v\n", part1)
}
