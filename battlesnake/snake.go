package battlesnake

type Snake struct {
	State        GameState
	PreferedMove string
	Up           bool
	Down         bool
	Left         bool
	Right        bool
}

func (s *Snake) Calc() {
	// Step 0: Don't let your Battlesnake move back in on it's own neck
	myHead := s.State.You.Body[0] // Coordinates of your head
	myNeck := s.State.You.Body[1] // Coordinates of body piece directly behind your head (your "neck")
	if myNeck.X < myHead.X {
		s.Left = false
	} else if myNeck.X > myHead.X {
		s.Right = false
	} else if myNeck.Y < myHead.Y {
		s.Down = false
	} else if myNeck.Y > myHead.Y {
		s.Up = false
	}

	// TODO: Step 1 - Don't hit walls.
	// Use information in GameState to prevent your Battlesnake from moving beyond the boundaries of the board.
	boardWidth := s.State.Board.Width
	boardHeight := s.State.Board.Height

	if myHead.X == 0 {
		s.Left = false
	} else if myHead.X == boardWidth-1 {
		s.Right = false
	}

	if myHead.Y == 0 {
		s.Down = false
	} else if myHead.Y == boardHeight-1 {
		s.Up = false
	}

	// TODO: Step 2 - Don't hit yourself.
	// Use information in GameState to prevent your Battlesnake from colliding with itself.
	mybody := s.State.You.Body
	for i := range mybody {
		if mybody[i].X == myHead.X+1 && mybody[i].Y == myHead.Y {
			s.Right = false
		}
		if mybody[i].X == myHead.X-1 && mybody[i].Y == myHead.Y {
			s.Left = false
		}
		if mybody[i].X == myHead.X && mybody[i].Y == myHead.Y+1 {
			s.Up = false
		}
		if mybody[i].X == myHead.X && mybody[i].Y == myHead.Y-1 {
			s.Down = false
		}
	}

	// TODO: Step 3 - Don't collide with others.
	// Use information in GameState to prevent your Battlesnake from colliding with others.
	for _, sn := range s.State.Board.Snakes {
		snake := sn.Body
		// myHead = s.Head
		for i := range snake {
			if snake[i].X == myHead.X+1 && snake[i].Y == myHead.Y {
				s.Right = false
			}
			if snake[i].X == myHead.X-1 && snake[i].Y == myHead.Y {
				s.Left = false
			}
			if snake[i].X == myHead.X && snake[i].Y == myHead.Y+1 {
				s.Up = false
			}
			if snake[i].X == myHead.X && snake[i].Y == myHead.Y-1 {
				s.Down = false
			}
		}
	}
	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.
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

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.

	if s.PreferedMove == "" {
		if s.Down {
			s.PreferedMove = "down"
		} else if s.Up {
			s.PreferedMove = "up"
		} else if s.Left {
			s.PreferedMove = "left"
		} else if s.Right {
			s.PreferedMove = "righ"
		}
	}
}

// func getWaysOut(from Coord, )

func manhatanDistanceBetween(from Coord, to Coord) int {
	return abs(from.X-to.X) + abs(from.Y-to.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
