package constructing_test

import (
	"testing"

	"github.com/fernandoocampo/katas/constructing"
)

type data struct {
	str1 string
	str2 string
	r    int
}

var tables = []data{
	{"aba", "abbabba", 9},
	{"aba", "abaa", 4},
	{"aba", "a", 3},
	{"a", "a", 1},
	{"a", "aaa", 3},
	{"abcdefgh", "hgfedcba", 64},
	{"sxdfcgdgxdfgdxxf", "xgdfsxgdxgfsgdfxgfdgfgdfgdgsdfxgfdxgdfxgdgdfgdfxgsdxgdfxgfgsxfgdfgsdfxgdfxgsgsfgxsgsdxgsdfxgsgsdfxgsdfx", 288},
}

func TestStringConstructing(t *testing.T) {
	for _, item := range tables {
		result := constructing.StringConstructing(item.str1, item.str2)
		if result != item.r {
			t.Errorf("Given parameters %s, %s, it is expected got %d, but it got %d", item.str1, item.str2, item.r, result)
		}
	}
}
