package divisibility_test

import (
	"testing"

	"github.com/fernandoocampo/katas/divisibility"
)

type data struct {
	n int
	r int
}

var tables = []data{
	{8529, 79},
	{85299258, 31},
	{5634, 57},
}

func TestThirt(t *testing.T) {
	for _, item := range tables {
		result := divisibility.Thirt(item.n)

		if result != item.r {
			t.Errorf("with given param %d, it expects result %d, but it got %d", item.n, item.r, result)
		}
	}
}
