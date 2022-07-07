package battlesnake

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

func (c *Coord) manhatanDistanceBetween(to Coord) int {
	return abs(c.X-to.X) + abs(c.Y-to.Y)
}

func (c *Coord) isNextTo(to Coord) bool {
	return c.manhatanDistanceBetween(to) == 1
}

func (c *Coord) findClosest(cs []Coord) Coord {
	var closestCoord Coord
	distance := 9999
	for _, c2 := range cs {
		dist := c.manhatanDistanceBetween(c2)
		if dist < distance {
			closestCoord = c2
			distance = dist
		}
	}
	return closestCoord
}
