package battlesnake

import (
	"errors"
	"fmt"
)

type Path struct {
	valid    bool
	dest     Coord
	origin   Coord
	path     map[Coord]Coord
	distance int
}

func (p *Path) getNextDirection() (string, error) {
	if p.valid {
		v, _ := p.getNextCoord()
		fmt.Println(v, p.origin)
		if v.X > p.origin.X {
			return "right", nil
		} else if v.X < p.origin.X {
			return "left", nil
		} else if v.Y < p.origin.Y {
			return "down", nil
		} else if v.Y > p.origin.Y {
			return "up", nil
		}
	}
	return "", errors.New("no path")
}

func (p *Path) getNextCoord() (*Coord, error) {
	var current Coord
	current = p.dest
	if p.valid {
		for {
			if p.origin == p.path[current] {
				return &current, nil
			} else {
				current = p.path[current]
			}
		}
	} else {
		return nil, errors.New("no path")
	}
}
