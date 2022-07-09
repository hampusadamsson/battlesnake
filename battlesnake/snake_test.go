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
			Food:   []Coord{{8, 8, 0}},
		},
		You: me,
	}
	s := SnakeNew(state)
	action := me.findMostLimitingMove(s.State.Board.Snakes, 3, &state.Board)
	assert.True(t, action.get("down") > action.get("right"))
	assert.Equal(t, "down", s.GetAction())
	s.GetAction()
}

func TestForcedMove(t *testing.T) {
	foe := Battlesnake{
		Head:   Coord{X: 0, Y: 0},
		Health: 50,
		Body: []Coord{
			{X: 1, Y: 0},
			{X: 2, Y: 0},
			{X: 3, Y: 0}},
	}
	me := Battlesnake{
		Head:   Coord{X: 1, Y: 1},
		Health: 50,
		Body: []Coord{
			{X: 2, Y: 1},
			{X: 3, Y: 1}},
	}
	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me, foe},
			Height: 4,
			Width:  4,
			Food:   []Coord{{0, 1, 0}},
		},
		You: me,
	}
	s := SnakeNew(state)
	s.GetAction()
	//assert.Equal(t, "down", s.GetAction())
}

func TestState(t *testing.T) {
	me := Battlesnake{
		Head:   Coord{X: 3, Y: 2},
		Health: 50,
		Body: []Coord{
			{X: 2, Y: 2}},
	}
	foe := Battlesnake{
		Head:   Coord{X: 1, Y: 1},
		Health: 50,
		Body: []Coord{
			{X: 0, Y: 1}},
	}
	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me, foe},
			Height: 4,
			Width:  4,
			Food:   []Coord{{0, 0, 0}},
		},
		You: me,
	}
	s := SnakeNew(state)

	fmt.Println(s)
	s2 := SnakeNew(state)
	s2.State.Board.Hazards = append(s2.State.Board.Hazards, Coord{9, 9, 0})
	fmt.Println(s)
	fmt.Println(s2.State)
}

func TestOpenSPace(t *testing.T) {
	foe := Battlesnake{
		Head:   Coord{X: 0, Y: 1},
		Health: 50,
		Body: []Coord{
			{X: 0, Y: 0}},
	}
	me := Battlesnake{
		Head:   Coord{X: 2, Y: 2},
		Health: 100,
		Body: []Coord{
			{X: 2, Y: 1},
			{X: 2, Y: 0},
		},
	}

	state := GameState{
		Board: Board{
			Snakes: []Battlesnake{me, foe},
			Height: 3,
			Width:  5,
		},
		You: me,
	}
	s := SnakeNew(state)
	m := me.findOpenSpace(s.State.You.Head, 2, &s.State.Board)
	assert.Equal(t, m.get("left"), 4)
	assert.Equal(t, m.get("right"), 6)
	assert.Equal(t, m.get("down"), 0)
	assert.Equal(t, m.get("up"), 0)
}
