package battlesnake

import (
	"fmt"
	"testing"
)

func TestPenalty(t *testing.T) {
	me := Battlesnake{
		Head:   Coord{X: 1, Y: 0},
		Health: 50,
		Body: []Coord{
			{X: 1, Y: 1},
			{X: 1, Y: 2},
		}}

	mm := make(map[Coord]int)
	mm[Coord{X: 0, Y: 0}] = 100
	b := Board{
		Snakes:    []Battlesnake{me},
		Height:    5,
		Width:     5,
		Food:      []Coord{{0, 0, 0}},
		Penalties: mm,
	}
	p := b.getPenalties(me.Head)
	fmt.Println(p)
}

func TestFindFood(t *testing.T) {

	me := Battlesnake{
		Head:   Coord{X: 1, Y: 0},
		Health: 50,
		Body: []Coord{
			{X: 1, Y: 1},
			{X: 1, Y: 2},
		}}

	foe := Battlesnake{
		Head:   Coord{X: 3, Y: 3},
		Health: 50,
		Body: []Coord{
			{X: 3, Y: 2},
			{X: 3, Y: 1}},
	}

	b := Board{
		Snakes: []Battlesnake{me, foe},
		Height: 5,
		Width:  5,
		Food:   []Coord{{0, 0, 0}},
	}
	path := b.findWayToFood(&Coord{4, 4, 0})
	fmt.Println(path.getNextCoord())
	fmt.Println(path.getNextDirection())
}
