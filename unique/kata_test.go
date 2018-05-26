package unique_test

import (
	"testing"

	"github.com/fernandoocampo/katas/unique"
)

type data struct {
	p []float32
	r float32
}

var tables = []data{
	{[]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}, float32(2)},
	{[]float32{0, 0, 0.55, 0, 0}, float32(0.55)},
}

func TestFindUniq(t *testing.T) {
	for _, item := range tables {
		result := unique.FindUniq(item.p)

		if result != item.r {
			t.Errorf("Given parameter %v, it expects %f, but it got %f", item.p, item.r, result)
		}
	}
}

func TestFindUniqTwo(t *testing.T) {
	for _, item := range tables {
		result := unique.FindUniqTwo(item.p)

		if result != item.r {
			t.Errorf("Given parameter %v, it expects %f, but it got %f", item.p, item.r, result)
		}
	}
}

func TestFindUniqThree(t *testing.T) {
	for _, item := range tables {
		result := unique.FindUniqThree(item.p)

		if result != item.r {
			t.Errorf("Given parameter %v, it expects %f, but it got %f", item.p, item.r, result)
		}
	}
}
