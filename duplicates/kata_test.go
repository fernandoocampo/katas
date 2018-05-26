package duplicates_test

import (
	"testing"

	"github.com/fernandoocampo/katas/duplicates"
)

type data struct {
	p string
	r int
}

var tables = []data{
	{"abcde", 0},
	{"abcdea", 1},
	{"abcdeaB11", 3},
	{"indivisibility", 1},
}

func TestDuplicateCount(t *testing.T) {
	for _, item := range tables {
		result := duplicates.Duplicate_count(item.p)
		if result != item.r {
			t.Errorf("Given parameter %s, it expects %d, but it got %d", item.p, item.r, result)
		}
	}
}

func TestDuplicateCountTwo(t *testing.T) {
	for _, item := range tables {
		result := duplicates.Duplicate_count_two(item.p)
		if result != item.r {
			t.Errorf("Given parameter %s, it expects %d, but it got %d", item.p, item.r, result)
		}
	}
}
