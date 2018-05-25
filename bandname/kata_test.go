package bandname_test

import (
	"testing"

	"github.com/fernandoocampo/katas/bandname"
)

type data struct {
	p string
	r string
}

var tables = []data{
	{"knife", "The Knife"},
	{"tart", "Tartart"},
	{"sandles", "Sandlesandles"},
	{"bed", "The Bed"},
	{"The Step-daughter", "The Step-Daughter"},
}

func TestBandNameGenerator(t *testing.T) {
	for _, item := range tables {
		result := bandname.BandNameGenerator(item.p)
		if item.r != result {
			t.Errorf("given parameters %s, it expects '%s', but it got '%s'", item.p, item.r, result)
		}
	}
}
