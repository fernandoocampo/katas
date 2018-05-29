package howmuch_test

import (
	"testing"

	"github.com/fernandoocampo/katas/howmuch"
)

type data struct {
	p1 int
	p2 int
	r  [][3]string
}

var tables = []data{
	{1, 100, [][3]string{{"M: 37", "B: 5", "C: 4"}, {"M: 100", "B: 14", "C: 11"}}},
	{0, 200, [][3]string{{"M: 37", "B: 5", "C: 4"}, {"M: 100", "B: 14", "C: 11"}, {"M: 163", "B: 23", "C: 18"}}},
	{1500, 1600, [][3]string{{"M: 1549", "B: 221", "C: 172"}}},
	{2950, 2950, nil},
	{20000, 20100, [][3]string{{"M: 20008", "B: 2858", "C: 2223"}, {"M: 20071", "B: 2867", "C: 2230"}}},
}

func TestHowMuch(t *testing.T) {
	for _, item := range tables {
		result := howmuch.HowMuch(item.p1, item.p2)
		if result != item.r {
			t.Errorf("Given parameters %d and %d, it expects %x, but it got %x", item.p1, item.p2, item.r, result)
		}
	}
}
