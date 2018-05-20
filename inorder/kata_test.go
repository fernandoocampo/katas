package inorder_test

import (
	"testing"

	"github.com/fernandoocampo/katas/inorder"
)

type data struct {
	p []int
	r bool
}

var tables = []data{
	{[]int{1, 2, 4, 7, 19}, true},
	{[]int{1, 2, 3, 4, 5}, true},
	{[]int{1, 6, 10, 18, 2, 4, 20}, false},
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, false},
}

func TestIsAscOrder(t *testing.T) {
	for _, item := range tables {
		result := inorder.InAscOrder(item.p)

		if result != item.r {
			t.Errorf("Given parameter %v, it expects %t, but it got %t", item.p, item.r, result)
		}
	}
}

func TestIsAscOrderTwo(t *testing.T) {
	for _, item := range tables {
		result := inorder.InAscOrderTwo(item.p)

		if result != item.r {
			t.Errorf("Given parameter %v, it expects %t, but it got %t", item.p, item.r, result)
		}
	}
}
