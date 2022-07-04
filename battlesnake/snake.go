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

func (s *Snake) aggDistanceFromOthersToCoord(c Coord) int {
	penalty := 0
	for i := range s.State.Board.Snakes {
		snek := s.State.Board.Snakes[i]
		if !snek.IsYou {
			penalty += manhatanDistanceBetween(snek.Head, c)
		}
	}
	return penalty
}

func (s *Snake) distanceToNearestEnemy(c Coord) int {
	currentNearest := 9999999
	for i := range s.State.Board.Snakes {
		snek := s.State.Board.Snakes[i]
		if !snek.IsYou {
			dist := manhatanDistanceBetween(snek.Head, c)
			if dist < currentNearest {
				currentNearest = dist
			}
		}
	}
	return currentNearest
}

func (s *Snake) findFood() string {
	myHead := s.State.You.Body[0]
	var nearestFood Coord
	distance := 999999
	for _, c := range s.State.Board.Food {
		dist := manhatanDistanceBetween(myHead, c)
		dist -= s.aggDistanceFromOthersToCoord(c)
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
func (s *Snake) possibleFutureMoves(c Coord, depth int, blockedCoord *Coord, hist map[Coord]int) int {
	if s.State.Board.isOckupied(c) || (blockedCoord != nil && c.X == blockedCoord.X && c.Y == blockedCoord.Y) {
		return 0
	}
	if _, ok := hist[c]; ok {
		return 0
	} else {
		hist[c] = 1
	}
	if depth < 0 {
		return 1
	}
	cc := 0
	adjacent := c.adjacent()
	for i := range adjacent {
		if !s.State.Board.isOckupied(adjacent[i]) {
			cc += s.possibleFutureMoves(adjacent[i], depth-1, blockedCoord, hist)
		}
	}
	return cc + 1
}

func (s *Snake) helper(snek *Battlesnake, lookahead int, blockedCoord Coord, dir string, m map[string]int, possiblePathCount map[string]int) {
	be := s.countOpenPaths(snek.Head, lookahead, &blockedCoord)
	newLimitingFactor := s.countOpenPaths(snek.Head, lookahead, nil) - be
	if m[dir] < newLimitingFactor {
		m[dir] = newLimitingFactor
		possiblePathCount[dir] = be
	}
}

func (s *Snake) findMostLimitingMove(head Coord, lookahead int) (map[string]int, map[string]int, []string) {
	limitingFactorMap := make(map[string]int)
	possiblePathCount := make(map[string]int)
	limitingFactorMap["left"] = 0
	limitingFactorMap["right"] = 0
	limitingFactorMap["up"] = 0
	limitingFactorMap["down"] = 0
	possiblePathCount["left"] = 9999
	possiblePathCount["right"] = 9999
	possiblePathCount["up"] = 9999
	possiblePathCount["down"] = 9999
	for i := range s.State.Board.Snakes {
		snek := s.State.Board.Snakes[i]
		if !snek.IsYou {
			//left
			if !s.State.Board.isOckupied(head.left()) {
				s.helper(&snek, lookahead, head.left(), "left", limitingFactorMap, possiblePathCount)
			}
			//right
			if !s.State.Board.isOckupied(head.righ()) {
				s.helper(&snek, lookahead, head.righ(), "right", limitingFactorMap, possiblePathCount)
			}
			//up
			if !s.State.Board.isOckupied(head.up()) {
				s.helper(&snek, lookahead, head.up(), "up", limitingFactorMap, possiblePathCount)
			}
			//down
			if !s.State.Board.isOckupied(head.down()) {
				s.helper(&snek, lookahead, head.down(), "down", limitingFactorMap, possiblePathCount)
			}
		}
	}
	keys := []string{"left", "right", "up", "down"}
	sort.SliceStable(keys, func(i, j int) bool {
		return limitingFactorMap[keys[i]] > limitingFactorMap[keys[j]]
	})
	return limitingFactorMap, possiblePathCount, keys
}

func (s *Snake) countOpenPaths(targetHead Coord, depthLookahead int, blockedCoord *Coord) int {
	l := s.possibleFutureMoves(targetHead.left(), depthLookahead, blockedCoord, make(map[Coord]int))
	r := s.possibleFutureMoves(targetHead.righ(), depthLookahead, blockedCoord, make(map[Coord]int))
	u := s.possibleFutureMoves(targetHead.up(), depthLookahead, blockedCoord, make(map[Coord]int))
	d := s.possibleFutureMoves(targetHead.down(), depthLookahead, blockedCoord, make(map[Coord]int))
	return l + r + d + u
}

func (s *Snake) findOpenSpace(c Coord, depth int) (string, map[string]int) {
	m := make(map[string]int)
	m["left"] = s.possibleFutureMoves(c.left(), depth, nil, make(map[Coord]int))
	m["right"] = s.possibleFutureMoves(c.righ(), depth, nil, make(map[Coord]int))
	m["up"] = s.possibleFutureMoves(c.up(), depth, nil, make(map[Coord]int))
	m["down"] = s.possibleFutureMoves(c.down(), depth, nil, make(map[Coord]int))

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

func (s *Snake) getMostCenteringMove() map[string]int {
	m := make(map[string]int)
	m["left"] = abs(s.State.You.Head.X - 1 - s.State.Board.Width)
	m["right"] = abs(s.State.You.Head.X + 1 - s.State.Board.Width)
	m["up"] = abs(s.State.You.Head.Y - 1 - s.State.Board.Height)
	m["down"] = abs(s.State.You.Head.Y - 1 - s.State.Board.Height)
	return m
}

func (s *Snake) youHaveMostLife() bool {
	for i := range s.State.Board.Snakes {
		if !s.State.Board.Snakes[i].IsYou && len(s.State.Board.Snakes[i].Body) >= len(s.State.You.Body) {
			return false
		}
	}
	return true
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
	safeMove, dirFreeSpace := s.findOpenSpace(s.State.You.Head, 7)

	// strive for center
	mostCenterMoves := s.getMostCenteringMove()

	// Find best destruction move
	maxLimitPerDirection, totPathForEnemyInThatDir, paths := s.findMostLimitingMove(s.State.You.Head, 7)
	// fmt.Println("DESTRUCTION:", limitingMove, limitingFactor)

	// Prioritize movement
	keys := []string{"left", "right", "up", "down"}
	composite := make(map[string]int)
	for _, v := range keys {
		composite[v] = dirFreeSpace[v] + maxLimitPerDirection[v] - mostCenterMoves[v]*3
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return composite[keys[i]] > composite[keys[j]]
	})
	fmt.Println("REC:", keys[0], composite)

	// 1) survival
	// 1.1) health
	// 1.2) free space
	// 2) attacking

	fmt.Println("-------------")
	fmt.Println("REC:", keys[0])
	fmt.Println("composite", composite)
	fmt.Println("dirFreeSpace", dirFreeSpace)
	fmt.Println("maxLimitPerDirection", maxLimitPerDirection)
	fmt.Println("mostCenterMoves", mostCenterMoves)
	fmt.Println("totPathForEnemyInThatDir", totPathForEnemyInThatDir)

	if eatMove != "" {
		if s.State.You.Health < 10 {
			fmt.Println("--------HEAL")
			s.PreferedMove = eatMove
		} else if maxLimitPerDirection[paths[0]] < 3 && dirFreeSpace[paths[0]] > totPathForEnemyInThatDir[paths[0]] && s.distanceToNearestEnemy(s.State.You.Head) > 2 {
			fmt.Println("--------KILLer000")
			s.PreferedMove = paths[0]
		} else if !s.youHaveMostLife() && dirFreeSpace[eatMove] > 5 && s.distanceToNearestEnemy(s.State.You.Head) > 2 {
			fmt.Println("--------HEAL")
			s.PreferedMove = eatMove
		} else if s.State.You.Health < 75 && dirFreeSpace[eatMove] > 20 {
			fmt.Println("--------HEAL")
			s.PreferedMove = eatMove
		} else if s.State.You.Health < 50 && dirFreeSpace[eatMove] > 15 {
			fmt.Println("--------HEAL")
			s.PreferedMove = eatMove
		} else if s.State.You.Health < 30 && dirFreeSpace[eatMove] > 10 {
			fmt.Println("--------HEAL")
			s.PreferedMove = eatMove
		} else if s.State.You.Health < 20 && dirFreeSpace[eatMove] > 8 {
			fmt.Println("--------HEAL")
			s.PreferedMove = eatMove
		} else if s.State.You.Health < 10 && dirFreeSpace[eatMove] > 5 {
			fmt.Println("--------HEAL")
			s.PreferedMove = eatMove
			// } else if limitingMove != "" && dirFreeSpace[limitingMove] > 10 && limitingFactor > (dirFreeSpace[limitingMove]*2) {
		} else if dirFreeSpace[keys[0]] > 5 {
			fmt.Println("--------DEFAULT 1")
			s.PreferedMove = keys[0]
		} else {
			fmt.Println("--------DEFAULT 2")
			// if limitingMove != "" && dirFreeSpace[limitingMove] > 10 { // && (limitingFactor > dirFreeSpace[limitingMove]) {
			// 	s.PreferedMove = limitingMove
			// } else {
			s.PreferedMove = safeMove
			// }
			// s.PreferedMove = safeMove
		}
	} else if dirFreeSpace[keys[0]] > 10 {
		fmt.Println("--------DEFAULT 3")
		s.PreferedMove = keys[0]
	} else {
		fmt.Println("--------DEFAULT 4")
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
