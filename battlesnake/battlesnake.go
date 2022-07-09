package battlesnake

func (b *Battlesnake) isOckupiedBySnake(c Coord) bool {
	mybody := b.Body
	if b.Head.X == c.X && b.Head.Y == c.Y {
		return true
	}
	for i := range b.Body {
		if mybody[i].X == c.X && mybody[i].Y == c.Y && i != len(b.Body)-1 { // Skip tail ?
			return true
		}
	}
	return false
}

func (b *Battlesnake) expectedSnakeNextMove() Coord {
	myNeck := b.Body[1]
	myHead := b.Body[0]
	if myNeck.X < myHead.X {
		return myHead.righ()
	} else if myNeck.X > myHead.X {
		return myHead.left()
	} else if myNeck.Y < myHead.Y {
		return myHead.up()
	} else if myNeck.Y > myHead.Y {
		return myHead.down()
	} else {
		return myHead
	}
}

type Battlesnake struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Health  int     `json:"health"`
	Body    []Coord `json:"body"`
	Head    Coord   `json:"head"`
	Length  int32   `json:"length"`
	Latency string  `json:"latency"`
	// IsYou   bool

	Shout string `json:"shout"`
	Squad string `json:"squad"`
}

func (b *Battlesnake) possibleFutureMoves(c Coord, depth int, blockedCoord *Coord, hist map[Coord]int, board *Board) int {
	if board.isOckupied(c) || (blockedCoord != nil && c.X == blockedCoord.X && c.Y == blockedCoord.Y) {
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
		cc += b.possibleFutureMoves(adjacent[i], depth-1, blockedCoord, hist, board)
	}
	return cc + 1
}

func (b *Battlesnake) countOpenPaths(targetHead Coord, depthLookahead int, blockedCoord *Coord, board *Board) int {
	freeSpace := 0
	for _, adjacent := range targetHead.adjacent() {
		freeSpace += b.possibleFutureMoves(adjacent, depthLookahead, blockedCoord, make(map[Coord]int), board)
	}
	return freeSpace
}

func (b *Battlesnake) helper(snek *Battlesnake, lookahead int, blockedCoord Coord, dir string, m *decision, board *Board) {
	if !board.isOckupied(blockedCoord) {
		be := b.countOpenPaths(snek.Head, lookahead, &blockedCoord, board)
		newLimitingFactor := b.countOpenPaths(snek.Head, lookahead, nil, board) - be
		if m.get(dir) < newLimitingFactor {
			m.set(dir, newLimitingFactor)
		}
	}
}

func (b *Battlesnake) findMostLimitingMove(snakes []Battlesnake, lookahead int, board *Board) *decision {
	limitingFactorMap := makeDecision()
	for i := range snakes {
		snek := snakes[i]
		if b.Head != snek.Head {
			b.helper(&snek, lookahead, b.Head.left(), "left", limitingFactorMap, board)
			b.helper(&snek, lookahead, b.Head.righ(), "right", limitingFactorMap, board)
			b.helper(&snek, lookahead, b.Head.up(), "up", limitingFactorMap, board)
			b.helper(&snek, lookahead, b.Head.down(), "down", limitingFactorMap, board)
		}
	}
	return limitingFactorMap
}

func (b *Battlesnake) hasMostHealth(snakes []Battlesnake) bool {
	for i := range snakes {
		if snakes[i].Health > b.Health {
			return false
		}
	}
	return true
}

func (b *Battlesnake) findOpenSpace(c Coord, depth int, board *Board) *decision {
	openSpace := makeDecision()
	openSpace.set("left", b.possibleFutureMoves(c.left(), depth, nil, make(map[Coord]int), board))
	openSpace.set("right", b.possibleFutureMoves(c.righ(), depth, nil, make(map[Coord]int), board))
	openSpace.set("up", b.possibleFutureMoves(c.up(), depth, nil, make(map[Coord]int), board))
	openSpace.set("down", b.possibleFutureMoves(c.down(), depth, nil, make(map[Coord]int), board))
	return openSpace
}
