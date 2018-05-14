package fortune_test

import (
	"testing"

	"github.com/fernandoocampo/katas/fortune"
)

type fortuneData struct {
	f0 int
	p  float64
	c0 int
	n  int
	i  float64
	r  bool
}

var tables []fortuneData = []fortuneData{
	{100000, 1.0, 2000, 15, 1.0, true},
	{100000, 1.0, 9185, 12, 1.0, false},
	{100000000, 1.0, 100000, 50, 1.0, true},
	{100000000, 1.5, 10000000, 50, 1.0, false},
	{100000000, 5.0, 1000000, 50, 1.0, true},
	{100000, 1, 10000, 10, 1, true},
}

func TestFortune(t *testing.T) {

	for _, item := range tables {
		result := fortune.Fortune(item.f0, item.p, item.c0, item.n, item.i)

		if result != item.r {
			t.Errorf("For parameters %d, %f, %d, %d, %f it expects: %t, but it got %t", item.f0, item.p, item.c0, item.n, item.i, item.r, result)
		}
	}
}

func TestFortuneTwo(t *testing.T) {

	for _, item := range tables {
		result := fortune.FortuneTwo(item.f0, item.p, item.c0, item.n, item.i)

		if result != item.r {
			t.Errorf("For parameters %d, %f, %d, %d, %f it expects: %t, but it got %t", item.f0, item.p, item.c0, item.n, item.i, item.r, result)
		}
	}
}
