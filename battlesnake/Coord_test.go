package battlesnake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNear(t *testing.T) {
	c1 := Coord{4, 4}
	c2 := Coord{5, 4}
	assert.True(t, c1.isNextTo(c2))
}

func TestNearFalse(t *testing.T) {
	c1 := Coord{4, 3}
	c2 := Coord{5, 4}
	assert.False(t, c1.isNextTo(c2))
}
