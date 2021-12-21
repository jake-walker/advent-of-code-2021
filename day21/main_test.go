package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

var input = "Player 1 starting position: 4\nPlayer 2 starting position: 8"

func TestNewGame(t *testing.T) {
	actual := NewGame(input)
	expected := Game{
		Players: []Player{
			{
				Position: 4,
				Score:    0,
			},
			{
				Position: 8,
				Score:    0,
			},
		},
		BoardLength:  10,
		WinningScore: 1000,
		DieValue:     1,
		DieRolls:     0,
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("NewGame() = \n%v", diff)
	}
}

func TestPlayUntilWinner(t *testing.T) {
	actual := PlayUntilWinner(Game{
		Players: []Player{
			{
				Position: 4,
				Score:    0,
			},
			{
				Position: 8,
				Score:    0,
			},
		},
		BoardLength:  10,
		WinningScore: 1000,
		DieValue:     1,
		DieRolls:     0,
	})
	expected := Game{
		Players: []Player{
			{
				Position: 10,
				Score:    1000,
			},
			{
				Position: 3,
				Score:    745,
			},
		},
		BoardLength:  10,
		WinningScore: 1000,
		DieValue:     94,
		DieRolls:     993,
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("PlayUntilWinner() =\n%v", diff)
	}
}
