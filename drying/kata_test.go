package drying_test

import (
	"testing"

	"github.com/fernandoocampo/katas/drying"
)

type data struct {
	p0 int
	w0 int
	p1 int
	r  int
}

var tables = []data{
	{99, 100, 98, 50},
	{82, 127, 80, 114},
}

func TestPotatoes(t *testing.T) {
	for _, item := range tables {
		result := drying.Potatoes(item.p0, item.w0, item.p1)

		if result != item.r {
			t.Errorf("Given params %d, %d, %d, it expects %d but it got %d", item.p0, item.w0, item.p1, item.r, result)
		}
	}
}
