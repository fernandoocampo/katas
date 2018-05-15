package diagonal_test

import (
	"testing"

	"github.com/fernandoocampo/katas/diagonal"
)

type data struct {
	p1 int
	p2 int
	r  int
}

var tables = []data{
	{20, 3, 5985},
	{20, 4, 20349},
	{20, 5, 54264},
	{20, 15, 20349},
	{100, 0, 101},
	{100, 10, 158940114100040},
}

func TestDiagonal(t *testing.T) {
	for _, item := range tables {
		result := diagonal.Diagonal(item.p1, item.p2)

		if result != item.r {
			t.Errorf("Given the parameters %d, %d it expects %d, but it got %d",
				item.p1, item.p2, item.r, result)
		}
	}
}

func TestDiagonalTwo(t *testing.T) {
	for _, item := range tables {
		result := diagonal.DiagonalTwo(item.p1, item.p2)

		if result != item.r {
			t.Errorf("Given the parameters %d, %d it expects %d, but it got %d",
				item.p1, item.p2, item.r, result)
		}
	}
}

func TestDiagonalThree(t *testing.T) {
	for _, item := range tables {
		result := diagonal.DiagonalThree(item.p1, item.p2)

		if result != item.r {
			t.Errorf("Given the parameters %d, %d it expects %d, but it got %d",
				item.p1, item.p2, item.r, result)
		}
	}
}
