package battlesnake

import "fmt"

func SnakeNew(state GameState) Snake {
	s := Snake{State: state, Up: true, Down: true, Left: true, Right: true}
	return s
}

type Snake struct {
	State        GameState
	PreferedMove string
	Up           bool
	Down         bool
	Left         bool
	Right        bool
}

// func (s *Snake) avoidCollisionWithSelf() {
// 	myNeck := s.State.You.Body[1] // Coordinates of body piece directly behind your head (your "neck")
// 	myHead := s.State.You.Body[0] // Coordinates of your head
// 	if myNeck.X < myHead.X {
// 		s.Left = false
// 	} else if myNeck.X > myHead.X {
// 		s.Right = false
// 	} else if myNeck.Y < myHead.Y {
// 		s.Down = false
// 	} else if myNeck.Y > myHead.Y {
// 		s.Up = false
// 	}
// }

// func (s *Snake) avoidWalls() {
// 	boardWidth := s.State.Board.Width
// 	boardHeight := s.State.Board.Height
// 	myHead := s.State.You.Body[0] // Coordinates of your head
// 	if myHead.X == 0 {
// 		s.Left = false
// 	} else if myHead.X == boardWidth-1 {
// 		s.Right = false
// 	}

// 	if myHead.Y == 0 {
// 		s.Down = false
// 	} else if myHead.Y == boardHeight-1 {
// 		s.Up = false
// 	}
// }

// func (s *Snake) avoidBody(mybody []Coord) {
// 	myHead := s.State.You.Body[0]
// 	for i := range mybody {
// 		if mybody[i].X == myHead.X+1 && mybody[i].Y == myHead.Y {
// 			s.Right = false
// 		}
// 		if mybody[i].X == myHead.X-1 && mybody[i].Y == myHead.Y {
// 			s.Left = false
// 		}
// 		if mybody[i].X == myHead.X && mybody[i].Y == myHead.Y+1 {
// 			s.Up = false
// 		}
// 		if mybody[i].X == myHead.X && mybody[i].Y == myHead.Y-1 {
// 			s.Down = false
// 		}
// 	}
// }

// func (s *Snake) avoidSnakes() {
// 	for _, sn := range s.State.Board.Snakes {
// 		snake := sn.Body
// 		s.avoidBody(snake)
// 	}
// }

func (s *Snake) findFood() {
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
			s.PreferedMove = "right"
		} else if myHead.X > nearestFood.X && s.Left == true {
			s.PreferedMove = "left"
		} else if myHead.Y < nearestFood.Y && s.Up == true {
			s.PreferedMove = "up"
		} else if myHead.Y > nearestFood.Y && s.Down == true {
			s.PreferedMove = "down"
		}
	}
}

func (s *Snake) findOpenSpace(c Coord) {
	m := make(map[Coord]int)
	m[c] = 1
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

	// // Step 0: Don't let your Battlesnake move back in on it's own neck
	// s.avoidCollisionWithSelf()

	// // TODO: Step 1 - Don't hit walls.
	// // Use information in GameState to prevent your Battlesnake from moving beyond the boundaries of the board.
	// s.avoidWalls()

	// // TODO: Step 2 - Don't hit yourself.
	// // Use information in GameState to prevent your Battlesnake from colliding with itself.
	// s.avoidBody(s.State.You.Body)

	// // TODO: Step 3 - Don't collide with others.
	// // Use information in GameState to prevent your Battlesnake from colliding with others.
	// s.avoidSnakes()

	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.
	s.findFood()

	s.findOpenSpace(Coord{0, 0})

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.

	fmt.Println("PREF", s.PreferedMove)
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
		fmt.Println(s.Left, s.Right, s.Down, s.Up, s.PreferedMove)
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
