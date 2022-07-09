package battlesnake

import (
	"fmt"
)

func SnakeNew(state GameState) Engine {
	s := Engine{State: state}
	return s
}

type Engine struct {
	State        GameState
	PreferedMove string
}

func (s *Engine) findNearestFood() Coord {
	c := s.State.You.Head
	return c.findClosest(s.State.Board.Food)
}

// GetAction retrieves the action for the snake
func (s *Engine) GetAction() string {
	s.State.Board.Penalties = make(map[Coord]int)
	myHead := s.State.You.Head
	var youHaveMostHealth bool
	var myID int
	for i := range s.State.Board.Snakes {
		if s.State.Board.Snakes[i].Head == myHead {
			myID = i
			youHaveMostHealth = s.State.Board.Snakes[i].hasMostHealth(s.State.Board.Snakes)
		}
	}

	for i := range s.State.Board.Snakes {
		if myID != i {
			fmt.Println("SNAKE: ", s.State.Board.Snakes[i].Head)
			expectedAction, forExpectedPath, probabilityOfPaths := s.getActionForSnake(&s.State.Board.Snakes[i])
			fmt.Println("EXPECTED ACTION: ", expectedAction, forExpectedPath)

			// Introduce this into my plan
			if len(s.State.Board.Snakes[myID].Body) > len(s.State.Board.Snakes[i].Body) && s.State.Board.Snakes[myID].Head.isNextTo(forExpectedPath) {
			} else {
				fmt.Println("Most probabel moves:", probabilityOfPaths.highest(2))
				fmt.Println("Head", s.State.Board.Snakes[i].Head)
				for _, dir := range probabilityOfPaths.highest(2) {
					fmt.Println(">>>Head", s.State.Board.Snakes[i].Head, dir, s.State.Board.Snakes[i].Head.offset(dir))
					c := s.State.Board.Snakes[i].Head.offset(dir)
					//s.State.Board.Hazards = append(s.State.Board.Hazards, c)
					if youHaveMostHealth {
						s.State.Board.Penalties[c] = 50
					} else {
						s.State.Board.Penalties[c] = 100
					}
					fmt.Println("Is now danger:", c)
				}

				// s.State.Board.Hazards = append(s.State.Board.Hazards, s.State.Board.Snakes[i].Head)
				// s.State.Board.Snakes[i].Head = forExpectedPath
			}
		}
	}

	fmt.Println("---START---")
	action, kk, _ := s.getActionForSnake(&s.State.Board.Snakes[myID])
	fmt.Println(kk)
	fmt.Println(&s.State.Board.Hazards)
	fmt.Println("---END---")
	return action
}

func (e *Engine) getActionForSnake(snake *Battlesnake) (string, Coord, *decision) {
	// These moves can't be done
	impossibleMoves := e.State.Board.createImpossibleMoves(snake.Head)

	// Find best direction
	longFutureSpace := snake.findOpenSpace(snake.Head, 12, &e.State.Board)

	// Find best destruction move
	longFutureLimit := snake.findMostLimitingMove(e.State.Board.Snakes, 10, &e.State.Board)

	// Add penalties
	penalties := e.State.Board.getPenalties(snake.Head).invert()

	// Add hungry move
	hungryMove := makeDecision()
	p := e.State.Board.findWayToFood(&snake.Head)
	if p.valid {
		hungryDirection, _ := p.getNextDirection()
		hungryMove.set(hungryDirection, 30)
	}

	// Prioritize and aggregate movement
	composite := makeDecision()
	composite.addAll(
		hungryMove,
		impossibleMoves,
		longFutureSpace,
		longFutureLimit,
		penalties,
	)

	// Log
	fmt.Println("-------------")
	fmt.Println(snake.Head)
	fmt.Println("REC:", composite.max())
	fmt.Println("composite", composite)
	fmt.Println("dirFreeSpace", longFutureSpace)
	fmt.Println("maxLimitPerDirection", longFutureLimit)
	fmt.Println("hungryMove", hungryMove)
	fmt.Println("Penalties", penalties, e.State.Board.Penalties)

	e.PreferedMove = composite.max()

	switch e.PreferedMove {
	case "left":
		return e.PreferedMove, snake.Head.left(), composite
	case "right":
		return e.PreferedMove, snake.Head.righ(), composite
	case "up":
		return e.PreferedMove, snake.Head.up(), composite
	case "down":
		return e.PreferedMove, snake.Head.down(), composite
	}
	panic("Can't get here")
}
