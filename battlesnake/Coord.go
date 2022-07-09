package battlesnake

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`

	Cost int
}

func (c *Coord) left() Coord {
	return Coord{X: c.X - 1, Y: c.Y}
}

func (c *Coord) righ() Coord {
	return Coord{X: c.X + 1, Y: c.Y}
}

func (c *Coord) up() Coord {
	return Coord{X: c.X, Y: c.Y + 1}
}

func (c *Coord) down() Coord {
	return Coord{X: c.X, Y: c.Y - 1}
}

func (c *Coord) adjacent() []Coord {
	return []Coord{
		c.down(),
		c.up(),
		c.left(),
		c.righ(),
	}
}

func (c *Coord) offset(dir string) Coord {
	if dir == "down" {
		return Coord{X: c.X, Y: c.Y - 1}
	} else if dir == "up" {
		return Coord{X: c.X, Y: c.Y + 1}
	} else if dir == "left" {
		return Coord{X: c.X - 1, Y: c.Y}
	} else if dir == "right" {
		return Coord{X: c.X + 1, Y: c.Y}
	}
	panic("Cant get heres")
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
