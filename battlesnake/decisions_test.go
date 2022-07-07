package battlesnake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecsion(t *testing.T) {
	d := makeDecision()
	d.set("right", 2)
	d.set("up", 5)
	d.set("down", -1)
	assert.Equal(t, "up", d.max())
}

func TestDecsionGetDefault(t *testing.T) {
	d := makeDecision()
	d.set("right", 2)
	assert.Equal(t, 2, d.get("right"))
	assert.Equal(t, 0, d.get("up"))
}
