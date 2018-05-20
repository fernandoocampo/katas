package fight_test

import (
	"testing"

	"github.com/fernandoocampo/katas/fight"
)

type data struct {
	p1 float64
	p2 float64
	r  float64
}

var tables = []data{
	{100.0, 50.0, 50.0},
	{1.0, 100.0, 0.0},
}

func TestCombat(t *testing.T) {
	for _, item := range tables {
		result := fight.Combat(item.p1, item.p2)
		if result != item.r {
			t.Errorf("given health: %g and damage: %g, it expects %g, but it got %g", item.p1, item.p2, item.r, result)
		}
	}
}

func TestCombatTwo(t *testing.T) {
	for _, item := range tables {
		result := fight.CombatTwo(item.p1, item.p2)
		if result != item.r {
			t.Errorf("given health: %g and damage: %g, it expects %g, but it got %g", item.p1, item.p2, item.r, result)
		}
	}
}
