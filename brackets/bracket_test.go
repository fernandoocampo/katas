package brackets_test

import (
	"testing"

	"github.com/fernandoocampo/katas/brackets"
)

type data struct {
	p1 string
	p2 int
	r  int
}

var tables = []data{
	{"[ABC[23]][89]", 0, 8},
	{"[ABC[23]][89]", 4, 7},
	{"[ABC[23]][89]", 9, 12},
}

func TestWhichIs(t *testing.T) {
	for _, item := range tables {
		result := brackets.WhichIs(item.p1, item.p2)

		if result != item.r {
			t.Errorf("with params %s, %d, it expects %d, but it got %d", item.p1, item.p2, item.r, result)
		}
	}
}

func TestWhichIsTwo(t *testing.T) {
	for _, item := range tables {
		result := brackets.WhichIsTwo(item.p1, item.p2)

		if result != item.r {
			t.Errorf("with params %s, %d, it expects %d, but it got %d", item.p1, item.p2, item.r, result)
		}
	}
}

func TestWhichIsThree(t *testing.T) {
	for _, item := range tables {
		result := brackets.WhichIsThree(item.p1, item.p2)

		if result != item.r {
			t.Errorf("with params %s, %d, it expects %d, but it got %d", item.p1, item.p2, item.r, result)
		}
	}
}
