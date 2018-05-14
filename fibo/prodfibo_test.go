package fibo_test

import (
	"testing"

	"github.com/fernandoocampo/katas/fibo"
)

type data struct {
	p uint64
	r [3]uint64
}

var tables = []data{
	{714, [3]uint64{21, 34, 1}},
	{800, [3]uint64{34, 55, 0}},
	{4895, [3]uint64{55, 89, 1}},
	{5895, [3]uint64{89, 144, 0}},
	{74049690, [3]uint64{6765, 10946, 1}},
	{84049690, [3]uint64{10946, 17711, 0}},
}

func TestProductFib(t *testing.T) {
	for _, item := range tables {
		result := fibo.ProductFib(item.p)
		if result[0] != item.r[0] || result[1] != item.r[1] || result[2] != item.r[2] {
			t.Errorf("given parameters %d, it expects %v, but it got %v", item.p, item.r, result)
		}
	}
}
