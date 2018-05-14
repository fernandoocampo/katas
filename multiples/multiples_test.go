package multiples_test

import (
	"testing"

	"github.com/fernandoocampo/katas/multiples"
)

type data struct {
	x int
	y int
}

var tables = []data{
	{10, 23},
}

func TestMultiple3And5(t *testing.T) {
	for _, item := range tables {
		result := multiples.Multiple3And5(item.x)

		if result != item.y {
			t.Errorf("parameter %d, it expected result %d, but it got %d", item.x, item.y, result)
		}
	}
}

func TestMultiple3And5Two(t *testing.T) {
	for _, item := range tables {
		result := multiples.Multiple3And5Two(item.x)

		if result != item.y {
			t.Errorf("parameter %d, it expected result %d, but it got %d", item.x, item.y, result)
		}
	}
}
