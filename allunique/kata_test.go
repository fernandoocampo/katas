package allunique_test

import (
	"testing"

	"github.com/fernandoocampo/katas/allunique"
)

type data struct {
	p string
	r bool
}

var tables = []data{
	{"  nAa", false},
	{"abcde", true},
	{"++-", false},
	{"AaBbC", true},
}

func TestHasUniqueChar(t *testing.T) {
	for _, item := range tables {
		result := allunique.HasUniqueChar(item.p)
		if result != item.r {
			t.Errorf("Given parameter %s, it expects %t, but it got %t", item.p, item.r, result)
		}
	}
}

func TestHasUniqueCharTwo(t *testing.T) {
	for _, item := range tables {
		result := allunique.HasUniqueCharTwo(item.p)
		if result != item.r {
			t.Errorf("Given parameter %s, it expects %t, but it got %t", item.p, item.r, result)
		}
	}
}
