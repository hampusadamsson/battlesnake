package battlesnake

type GameState struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

type Game struct {
	ID      string  `json:"id"`
	Ruleset Ruleset `json:"ruleset"`
	Timeout int32   `json:"timeout"`
}

type Ruleset struct {
	Name     string   `json:"name"`
	Version  string   `json:"version"`
	Settings Settings `json:"settings"`
}

type Settings struct {
	FoodSpawnChance     int32  `json:"foodSpawnChance"`
	MinimumFood         int32  `json:"minimumFood"`
	HazardDamagePerTurn int32  `json:"hazardDamagePerTurn"`
	Royale              Royale `json:"royale"`
	Squad               Squad  `json:"squad"`
}

type Royale struct {
	ShrinkEveryNTurns int32 `json:"shrinkEveryNTurns"`
}

type Squad struct {
	AllowBodyCollisions bool `json:"allowBodyCollisions"`
	SharedElimination   bool `json:"sharedElimination"`
	SharedHealth        bool `json:"sharedHealth"`
	SharedLength        bool `json:"sharedLength"`
}

type Board struct {
	Height int           `json:"height"`
	Width  int           `json:"width"`
	Food   []Coord       `json:"food"`
	Snakes []Battlesnake `json:"snakes"`
	// Used in non-standard game modes
	Hazards []Coord `json:"hazards"`
}

func (b *Board) isOckupied(c Coord) bool {
	if c.Y < 0 {
		return true
	}
	if c.Y == b.Height {
		return true
	}
	if c.X < 0 {
		return true
	}
	if c.X == b.Width {
		return true
	}
	for i := range b.Hazards {
		if b.Hazards[i] == c {
			return true
		}
	}
	for i := range b.Snakes {
		if b.Snakes[i].isOckupiedBySnake(c) {
			return true
		}
		if !b.Snakes[i].IsYou {
			expectedCoord := b.Snakes[i].expectedSnakeNextMove()
			if expectedCoord.X == c.X && expectedCoord.Y == c.Y {
				return true
			}
		}
	}
	return false
}

func (s *Battlesnake) isOckupiedBySnake(c Coord) bool {
	mybody := s.Body
	if s.Head.X == c.X && s.Head.Y == c.Y {
		return true
	}
	for i := range s.Body {
		if mybody[i].X == c.X && mybody[i].Y == c.Y {
			return true
		}
	}
	return false
}

func (s *Battlesnake) expectedSnakeNextMove() Coord {
	myNeck := s.Body[1]
	myHead := s.Body[0]
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
		//panic("Error")
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
	IsYou   bool
	// Used in non-standard game modes
	Shout string `json:"shout"`
	Squad string `json:"squad"`
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c *Coord) left() Coord {
	return Coord{c.X - 1, c.Y}
}

func (c *Coord) righ() Coord {
	return Coord{c.X + 1, c.Y}
}

func (c *Coord) up() Coord {
	return Coord{c.X, c.Y + 1}
}

func (c *Coord) down() Coord {
	return Coord{c.X, c.Y - 1}
}

func (c *Coord) adjacent() []Coord {
	return []Coord{
		c.down(),
		c.up(),
		c.left(),
		c.righ(),
	}
}
