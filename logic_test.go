package main

import (
	"fmt"
	"testing"

	"github.com/BattlesnakeOfficial/starter-snake-go/battlesnake"
)

func TestEmpty(t *testing.T) {
	var tt string
	fmt.Println(tt)
	fmt.Println(tt == "")
}

func TestNeckAvoidance(t *testing.T) {
	// Arrange
	me := battlesnake.Battlesnake{
		// Length 3, facing right
		Head: battlesnake.Coord{X: 2, Y: 0},
		Body: []battlesnake.Coord{{X: 2, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 0}},
	}
	state := battlesnake.GameState{
		Board: battlesnake.Board{
			Snakes: []battlesnake.Battlesnake{me},
		},
		You: me,
	}

	// Act 1,000x (this isn't a great way to test, but it's okay for starting out)
	for i := 0; i < 1000; i++ {
		nextMove := move(state)
		// Assert never move left
		if nextMove.Move == "left" {
			t.Errorf("snake moved onto its own neck, %s", nextMove.Move)
		}
	}
}

// TODO: More GameState test cases!
