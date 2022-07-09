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
	// custom
	Penalties map[Coord]int
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

		// // AVOID OTHERS
		// if !b.Snakes[i].IsYou {
		// 	expectedCoord := b.Snakes[i].expectedSnakeNextMove()
		// 	if expectedCoord.X == c.X && expectedCoord.Y == c.Y { // TODO - if you are bigger its ok
		// 		return true
		// 	}
		// }
	}
	return false
}

func (b *Board) getPenalties(c Coord) *decision {
	pen := makeDecision()

	adj := c.down()
	if p, ok := b.Penalties[adj]; ok {
		pen.set("down", p)
	}
	adj = c.left()
	if p, ok := b.Penalties[adj]; ok {
		pen.set("left", p)
	}
	adj = c.righ()
	if p, ok := b.Penalties[adj]; ok {
		pen.set("right", p)
	}
	adj = c.up()
	if p, ok := b.Penalties[adj]; ok {
		pen.set("up", p)
	}

	return pen
}

func (b *Board) createImpossibleMoves(c Coord) *decision {
	impossibleMoves := makeDecision()
	if b.isOckupied(c.left()) {
		impossibleMoves.set("left", -10000)
	}
	if b.isOckupied(c.righ()) {
		impossibleMoves.set("right", -10000)
	}
	if b.isOckupied(c.down()) {
		impossibleMoves.set("down", -10000)
	}
	if b.isOckupied(c.up()) {
		impossibleMoves.set("up", -10000)
	}
	return impossibleMoves
}

func (b *Board) findWayToFood(c *Coord) *Path {
	m := make(map[Coord]int)
	p := make(map[Coord]Coord)
	b.visit(c, c, 0, m, p)

	var closeFood Coord
	best := 999
	if len(b.Food) > 0 {
		for i := range b.Food {
			if v, ok := m[b.Food[i]]; ok {
				if v < best {
					closeFood = b.Food[i]
					best = v
				}
			}
		}
	} else {
		return &Path{valid: false}
	}
	if best == 999 {
		return &Path{valid: false}
	}
	return &Path{
		valid:    true,
		dest:     closeFood,
		origin:   *c,
		distance: best,
		path:     p,
	}
}

func (b *Board) visit(from *Coord, to *Coord, dist int, hist map[Coord]int, path map[Coord]Coord) {
	if b.isOckupied(*to) && from != to {
		return
	}
	if prev, ok := hist[*to]; ok {
		if prev > dist {
			hist[*to] = dist
			path[*to] = *from
		} else {
			return
		}
	} else {
		hist[*to] = dist
		path[*to] = *from
	}

	adjacent := to.adjacent()
	for i := range adjacent {
		b.visit(to, &adjacent[i], dist+1, hist, path)
	}
}
