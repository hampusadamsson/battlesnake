package battlesnake

import "testing"

func TestNeckAvoidance(t *testing.T) {
	// Arrange
	me := Battlesnake{
		// Length 3, facing right
		Head: Coord{X: 3, Y: 1},
		Body: []Coord{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}},
	}
	state := GameState{
		Board: Board{
			Height: 4,
			Width:  4,
			Snakes: []Battlesnake{me},
		},
		You: me,
	}
	s := SnakeNew(state)
	s.findOpenSpace(Coord{2, 2})
}
