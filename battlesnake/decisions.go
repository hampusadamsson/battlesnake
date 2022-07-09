package battlesnake

import (
	"sort"
)

type decision struct {
	path map[string]int
}

func makeDecision() *decision {
	return &decision{path: make(map[string]int)}
}

func (d *decision) set(key string, val int) {
	d.path[key] = val
}

func (d *decision) get(key string) int {
	if k, ok := d.path[key]; ok {
		return k
	}
	return 0
}

func (d *decision) addAll(c2 ...*decision) {
	for _, d2 := range c2 {
		d.add(d2)
	}
}

func (d *decision) add(c2 *decision) {
	keys := []string{"left", "right", "up", "down"}
	for _, v := range keys {
		val := d.get(v)
		d.set(v, val+c2.get(v))
	}
}

func (d *decision) sorted() []string {
	keys := []string{"left", "right", "up", "down"}
	sort.SliceStable(keys, func(i, j int) bool {
		return d.path[keys[i]] > d.path[keys[j]]
	})
	return keys
}

func (d *decision) max() string {
	return d.sorted()[0]
}

func (d *decision) highest(count int) []string {
	high := make([]string, 0)
	for _, v := range d.sorted()[:count] {
		high = append(high, v)
	}
	return high
}

func (d *decision) invert() *decision {
	return d.multiply(-1)
}

func (d *decision) multiply(mul int) *decision {
	keys := []string{"left", "right", "up", "down"}
	for _, v := range keys {
		val := d.get(v)
		d.set(v, val*mul)
	}
	return d
}
