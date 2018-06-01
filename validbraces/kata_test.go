package validbraces_test

import (
	"testing"

	"github.com/fernandoocampo/katas/validbraces"
)

type data struct {
	p string
	r bool
}

var tables = []data{
	{"(){}[]", true},
	{"([{}])", true},
	{"(}", false},
	{"[(])", false},
	{"[({)](]", false},
	{"{()}[{}]}{}[]", false},
}

func TestValidBraces(t *testing.T) {
	for _, item := range tables {
		result := validbraces.ValidBraces(item.p)
		if result != item.r {
			t.Errorf("Given parameter %s expect %t, but got %t", item.p, item.r, result)
		}
	}
}
