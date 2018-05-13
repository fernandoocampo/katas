package cycles_test

import (
	"testing"

	"github.com/fernandoocampo/katas/cycles"
)

func TestCycle(t *testing.T) {
	tables := []struct {
		x int
		y int
	}{
		{33, 2},
		{18118, -1},
		{69, 22},
		{197, 98},
		{65, -1},
		{94, -1},
	}

	for _, item := range tables {
		result := cycles.Cycle(item.x)

		if result != item.y {
			t.Errorf("for value %d is expected %d, but it got %d", item.x, item.y, result)
		}
	}
}
