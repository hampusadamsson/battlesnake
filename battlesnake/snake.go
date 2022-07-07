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

func (s *Engine) pathToNearbyFood() *decision {
	myHead := s.State.You.Head
	pathToFood := makeDecision()
	nearestFood := s.findNearestFood()
	if &nearestFood != nil {
		pathToFood.set("left", abs(myHead.left().X-nearestFood.X))
		pathToFood.set("right", abs(myHead.righ().X-nearestFood.X))
		pathToFood.set("up", abs(myHead.up().Y-nearestFood.Y))
		pathToFood.set("down", abs(myHead.down().Y-nearestFood.Y))
	}
	return pathToFood.invert() // invert to able maximize()
}

func (s *Engine) createImpossibleMoves(c Coord) *decision {
	impossibleMoves := makeDecision()
	if s.State.Board.isOckupied(c.left()) {
		impossibleMoves.set("left", -10000)
	}
	if s.State.Board.isOckupied(c.righ()) {
		impossibleMoves.set("right", -10000)
	}
	if s.State.Board.isOckupied(c.down()) {
		impossibleMoves.set("down", -10000)
	}
	if s.State.Board.isOckupied(c.up()) {
		impossibleMoves.set("up", -10000)
	}
	return impossibleMoves
}

// GetAction retrieves the action for the snake
func (s *Engine) GetAction() string {
	myHead := s.State.You.Head
	var myID int
	for i := range s.State.Board.Snakes {
		if s.State.Board.Snakes[i].Head == myHead {
			myID = i
		}
	}

	for i := range s.State.Board.Snakes {
		if myID != i {
			fmt.Println("SNAKE: ", s.State.Board.Snakes[i].Head)
			expectedAction, forExpectedPath := s.getActionForSnake(&s.State.Board.Snakes[i])
			fmt.Println("EXPECTED ACTION: ", expectedAction, forExpectedPath)

			// Introduce this into my plan
			if len(s.State.Board.Snakes[myID].Body) > len(s.State.Board.Snakes[i].Body) && s.State.Board.Snakes[myID].Head.isNextTo(forExpectedPath) {
			} else {
				// for _, c := range s.State.Board.Snakes[i].Head.adjacent() {
				// 	s.State.Board.Hazards = append(s.State.Board.Hazards, c)
				// }
				s.State.Board.Hazards = append(s.State.Board.Hazards, s.State.Board.Snakes[i].Head)
				s.State.Board.Snakes[i].Head = forExpectedPath
			}
		}
	}

	fmt.Println("---START---")
	action, kk := s.getActionForSnake(&s.State.Board.Snakes[myID])
	fmt.Println(kk)
	fmt.Println(&s.State.Board.Hazards)
	fmt.Println("---END---")
	return action
}

func (e *Engine) getActionForSnake(snake *Battlesnake) (string, Coord) {
	// These moves can't be done
	impossibleMoves := e.createImpossibleMoves(snake.Head)

	// Find food.
	eatMove := e.pathToNearbyFood()
	eatMove.multiply(7)

	// Find best direction
	longFutureSpace := snake.findOpenSpace(snake.Head, 10, &e.State.Board)

	// Find best destruction move
	longFutureLimit := snake.findMostLimitingMove(e.State.Board.Snakes, 10, &e.State.Board)

	// Prioritize and aggregate movement
	composite := makeDecision()
	composite.addAll(
		// eatMove,
		impossibleMoves,
		longFutureSpace,
		longFutureLimit,
	)

	// Log
	fmt.Println("-------------")
	fmt.Println(snake.Head)
	fmt.Println("REC:", composite.max())
	fmt.Println("composite", composite)
	fmt.Println("dirFreeSpace", longFutureSpace)
	fmt.Println("maxLimitPerDirection", longFutureLimit)
	fmt.Println("eatMove", eatMove)

	// Act
	// if len(e.State.Board.Food) != 0 && e.State.You.Health < 10 {
	// 	composite.add(eatMove.multiply(10))
	// 	e.PreferedMove = composite.max()
	// } else {
	// 	e.PreferedMove = composite.max()
	// }

	e.PreferedMove = composite.max()

	p := e.State.Board.findWayToFood(&snake.Head)
	if p.valid {
		hungryDirection, _ := p.getNextDirection()
		if longFutureSpace.get(hungryDirection) > 20 {
			e.PreferedMove = hungryDirection
		}
	}

	switch e.PreferedMove {
	case "left":
		return e.PreferedMove, snake.Head.left()
	case "right":
		return e.PreferedMove, snake.Head.righ()
	case "up":
		return e.PreferedMove, snake.Head.up()
	case "down":
		return e.PreferedMove, snake.Head.down()
	}
	panic("Can't get here")
}

// ---------- HELPER FUNCTIONS -----------

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
