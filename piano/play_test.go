package piano_test

import (
	"testing"

	"github.com/fernandoocampo/katas/piano"
)

type data struct {
	p int
	r string
}

var tables = []data{
	{1, "white"},
	{2, "black"},
	{5, "black"},
	{12, "black"},
	{42, "white"},
	{88, "white"},
	{89, "white"},
	{92, "white"},
	{100, "black"},
	{111, "white"},
	{200, "black"},
	{2017, "white"},
}

func TestBlackOrWhiteKey(t *testing.T) {
	for _, item := range tables {
		result := piano.BlackOrWhiteKey(item.p)

		if result != item.r {
			t.Errorf("Given parameter %d, it expects %s, but it got %s", item.p, item.r, result)
		}
	}
}

func TestBlackOrWhiteKeyTwo(t *testing.T) {
	for _, item := range tables {
		result := piano.BlackOrWhiteKeyTwo(item.p)

		if result != item.r {
			t.Errorf("Given parameter %d, it expects %s, but it got %s", item.p, item.r, result)
		}
	}
}
