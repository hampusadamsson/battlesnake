package battlesnake

import (
	"fmt"
	"sort"
)

func SnakeNew(state GameState) Snake {
	for i := range state.Board.Snakes {
		s := state.Board.Snakes[i].Head
		you := state.You.Head
		if s == you {
			state.Board.Snakes[i].IsYou = true
		}
	}
	s := Snake{State: state, Up: true, Down: true, Left: true, Right: true}
	return s
}

// type Direction string

// const (
// 	up    Direction = "up"
// 	Down  Direction = "down"
// 	Left  Direction = "left"
// 	Right Direction = "right"
// )

type Snake struct {
	State        GameState
	PreferedMove string
	Up           bool
	Down         bool
	Left         bool
	Right        bool
}

func (s *Snake) findFood() string {
	myHead := s.State.You.Body[0]
	var nearestFood Coord
	distance := 999999
	for _, c := range s.State.Board.Food {
		dist := manhatanDistanceBetween(myHead, c)
		if dist < distance {
			nearestFood = c
			distance = dist
		}
	}
	if &nearestFood != nil {
		if myHead.X < nearestFood.X && s.Right == true {
			return "right"
		} else if myHead.X > nearestFood.X && s.Left == true {
			return "left"
		} else if myHead.Y < nearestFood.Y && s.Up == true {
			return "up"
		} else if myHead.Y > nearestFood.Y && s.Down == true {
			return "down"
		}
	}
	return ""
}
func (s *Snake) possibleFutureMoves(c Coord, depth int, blockedCoord *Coord) int {
	if s.State.Board.isOckupied(c) || (blockedCoord != nil && c.X == blockedCoord.X && c.Y == blockedCoord.Y) {
		return 0
	}
	if depth < 0 {
		return 1
	}
	cc := 0
	adjacent := c.adjacent()
	for i := range adjacent {
		if !s.State.Board.isOckupied(adjacent[i]) {
			cc += s.possibleFutureMoves(adjacent[i], depth-1, blockedCoord)
		}
	}
	return cc
}

func (s *Snake) findMostLimitingMove(head Coord, lookahead int) (string, int, map[string]int) {
	bestDirection := ""
	limitingFactor := 0

	m := make(map[string]int)
	m["left"] = 0
	m["right"] = 0
	m["up"] = 0
	m["down"] = 0

	for i := range s.State.Board.Snakes {
		snek := s.State.Board.Snakes[i]
		if !snek.IsYou {
			//left
			if !s.State.Board.isOckupied(head.left()) {
				blockedCoord := head.left()
				newLimitingFactor := s.blockingEffect(snek.Head, lookahead, nil) - s.blockingEffect(snek.Head, lookahead, &blockedCoord)
				if limitingFactor < newLimitingFactor {
					bestDirection = "left"
					limitingFactor = newLimitingFactor
					m["left"] = limitingFactor
				}
			}
			//right
			if !s.State.Board.isOckupied(head.righ()) {
				blockedCoord := head.righ()
				newLimitingFactor := s.blockingEffect(snek.Head, lookahead, nil) - s.blockingEffect(snek.Head, lookahead, &blockedCoord)
				if limitingFactor < newLimitingFactor {
					bestDirection = "right"
					limitingFactor = newLimitingFactor
					m["right"] = limitingFactor
				}
			}
			//up
			if !s.State.Board.isOckupied(head.up()) {
				blockedCoord := head.up()
				newLimitingFactor := s.blockingEffect(snek.Head, lookahead, nil) - s.blockingEffect(snek.Head, lookahead, &blockedCoord)
				if limitingFactor < newLimitingFactor {
					bestDirection = "up"
					limitingFactor = newLimitingFactor
					m["up"] = limitingFactor
				}
			}
			//down
			if !s.State.Board.isOckupied(head.down()) {
				blockedCoord := head.down()
				newLimitingFactor := s.blockingEffect(snek.Head, lookahead, nil) - s.blockingEffect(snek.Head, lookahead, &blockedCoord)
				if limitingFactor < newLimitingFactor {
					bestDirection = "down"
					limitingFactor = newLimitingFactor
					m["down"] = limitingFactor
				}
			}

		}
	}
	return bestDirection, limitingFactor, m
}

func (s *Snake) blockingEffect(targetHead Coord, depthLookahead int, blockedCoord *Coord) int {
	l := s.possibleFutureMoves(targetHead.left(), depthLookahead, blockedCoord)
	r := s.possibleFutureMoves(targetHead.righ(), depthLookahead, blockedCoord)
	u := s.possibleFutureMoves(targetHead.up(), depthLookahead, blockedCoord)
	d := s.possibleFutureMoves(targetHead.down(), depthLookahead, blockedCoord)
	return l + r + d + u
}

func (s *Snake) findOpenSpace(c Coord, depth int) (string, map[string]int) {
	m := make(map[string]int)
	m["left"] = s.possibleFutureMoves(c.left(), depth, nil)
	m["right"] = s.possibleFutureMoves(c.righ(), depth, nil)
	m["up"] = s.possibleFutureMoves(c.up(), depth, nil)
	m["down"] = s.possibleFutureMoves(c.down(), depth, nil)

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	fmt.Println(c, keys[0], keys, m)
	return keys[0], m
}

func (s *Snake) GetAction() string {
	you := s.State.You

	if s.State.Board.isOckupied(you.Head.left()) {
		s.Left = false
	}
	if s.State.Board.isOckupied(you.Head.righ()) {
		s.Right = false
	}
	if s.State.Board.isOckupied(you.Head.down()) {
		s.Down = false
	}
	if s.State.Board.isOckupied(you.Head.up()) {
		s.Up = false
	}

	// Find food.
	eatMove := s.findFood()

	// Find best direction
	safeMove, dirFreeSpace := s.findOpenSpace(s.State.You.Head, 5)

	// Find best destruction move
	limitingMove, limitingFactor, maxLimitPerDirection := s.findMostLimitingMove(s.State.You.Head, 7)
	fmt.Println("DESTRUCTION:", limitingMove, limitingFactor)

	// Prioritize movement
	keys := []string{"left", "right", "up", "down"}
	composite := make(map[string]int)
	for _, v := range keys {
		composite[v] = dirFreeSpace[v] + maxLimitPerDirection[v]*3
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return composite[keys[i]] > composite[keys[j]]
	})
	fmt.Println("REC:", keys[0], composite)

	// 1) survival
	// 1.1) health
	// 1.2) free space
	// 2) attacking

	if eatMove != "" {
		if s.State.You.Health < 10 {
			s.PreferedMove = eatMove
		} else if s.State.You.Health < 75 && dirFreeSpace[eatMove] > 75 {
			s.PreferedMove = eatMove
		} else if s.State.You.Health < 50 && dirFreeSpace[eatMove] > 50 {
			s.PreferedMove = eatMove
		} else if s.State.You.Health < 30 && dirFreeSpace[eatMove] > 30 {
			s.PreferedMove = eatMove
		} else if s.State.You.Health < 20 && dirFreeSpace[eatMove] > 20 {
			s.PreferedMove = eatMove
		} else if s.State.You.Health < 10 && dirFreeSpace[eatMove] > 10 {
			s.PreferedMove = eatMove
			// } else if limitingMove != "" && dirFreeSpace[limitingMove] > 10 && limitingFactor > (dirFreeSpace[limitingMove]*2) {
		} else if dirFreeSpace[keys[0]] > 10 {
			s.PreferedMove = keys[0]
		} else {
			// if limitingMove != "" && dirFreeSpace[limitingMove] > 10 { // && (limitingFactor > dirFreeSpace[limitingMove]) {
			// 	s.PreferedMove = limitingMove
			// } else {
			s.PreferedMove = safeMove
			// }
			// s.PreferedMove = safeMove
		}
	} else {
		s.PreferedMove = safeMove
	}

	// Act on movement
	if s.PreferedMove == "" {
		if s.Down {
			s.PreferedMove = "down"
		} else if s.Up {
			s.PreferedMove = "up"
		} else if s.Left {
			s.PreferedMove = "left"
		} else if s.Right {
			s.PreferedMove = "right"
		}
	}
	return s.PreferedMove
}

// ---------- HELPER FUNCTIONS -----------

func manhatanDistanceBetween(from Coord, to Coord) int {
	return abs(from.X-to.X) + abs(from.Y-to.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
