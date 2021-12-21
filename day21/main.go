package main

import (
	"fmt"
	"github.com/jake-walker/advent-of-code-2021/helpers"
	"log"
	"strconv"
	"strings"
)

type Player struct {
	Position int
	Score    int
}

type Game struct {
	Players      []Player
	BoardLength  int
	WinningScore int
	DieValue     int
	DieRolls     int
}

func NewGame(input string) Game {
	lines := helpers.GetLines(input)
	players := make([]Player, len(lines))

	for i, line := range lines {
		pos, err := strconv.Atoi(strings.Split(line, ": ")[1])
		if err != nil {
			log.Fatalf("failed to convert position: %v", err)
		}

		players[i] = Player{
			Position: pos,
			Score:    0,
		}
	}

	return Game{
		Players:      players,
		BoardLength:  10,
		WinningScore: 1000,
		DieValue:     1,
	}
}

func PlayUntilWinner(game Game) Game {
	player := -1

	for {
		player = (player + 1) % len(game.Players)

		roll := game.DieValue + (game.DieValue + 1) + (game.DieValue + 2)
		game.DieRolls += 3
		game.DieValue = (game.DieValue + 3) % 100

		game.Players[player].Position = ((game.Players[player].Position + roll - 1) % game.BoardLength) + 1
		game.Players[player].Score += game.Players[player].Position

		log.Printf("Player %v rolls %v and moves to space %v for a total score of %v", player+1, roll, game.Players[player].Position, game.Players[player].Score)

		if game.Players[player].Score >= game.WinningScore {
			break
		}
	}

	return game
}

func main() {
	input := NewGame(helpers.GetInput("day21/input.txt"))

	part1 := PlayUntilWinner(input)
	losingScore := -1
	for _, player := range part1.Players {
		if player.Score < part1.WinningScore {
			losingScore = player.Score
		}
	}
	fmt.Printf("part 1: %v\n", losingScore*part1.DieRolls)
}
