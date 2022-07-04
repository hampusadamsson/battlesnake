package battlesnake

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	me := Battlesnake{
		Head:   Coord{X: 5, Y: 1},
		Health: 50,
		Body: []Coord{
			{X: 0, Y: 1},
			{X: 1, Y: 1},
			{X: 2, Y: 1},
			{X: 3, Y: 1},
			{X: 4, Y: 1}},
	}
	foe := Battlesnake{
		Head:   Coord{X: 3, Y: 0},
		Health: 100,
		Body: []Coord{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 2, Y: 0}},
	}

	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me, foe},
			Height: 12,
			Width:  12,
			Food:   []Coord{{8, 8}},
		},
		You: me,
	}
	s := SnakeNew(state)
	action := s.GetAction()

	fmt.Println("ACTION", action)
	assert.Equal(t, "down", action)
}

func TestOpenSPace(t *testing.T) {
	me := Battlesnake{
		Head:   Coord{X: 0, Y: 1},
		Health: 50,
		Body: []Coord{
			{X: 0, Y: 0}},
	}
	foe := Battlesnake{
		Head:   Coord{X: 2, Y: 2},
		Health: 100,
		Body: []Coord{
			{X: 2, Y: 1}},
	}

	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me, foe},
			Height: 3,
			Width:  3,
		},
		You: me,
	}
	s := SnakeNew(state)
	action, m := s.findOpenSpace(s.State.You.Head, 7)
	fmt.Println(action, m)

}
