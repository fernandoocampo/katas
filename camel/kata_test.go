package camel_test

import (
	"testing"

	"github.com/fernandoocampo/katas/camel"
)

type data struct {
	p string
	r string
}

var tables = []data{
	{"", ""},
	{"The_Stealth_Warrior", "TheStealthWarrior"},
	{"the-Stealth-Warrior", "theStealthWarrior"},
}

func TestToCamelCase(t *testing.T) {
	for _, item := range tables {
		result := camel.ToCamelCase(item.p)

		if result != item.r {
			t.Errorf("Given parameter %s, it expects %s, but it got %s", item.p, item.r, result)
		}
	}
}

func TestToCamelCaseTwo(t *testing.T) {
	for _, item := range tables {
		result := camel.ToCamelCaseTwo(item.p)

		if result != item.r {
			t.Errorf("Given parameter %s, it expects %s, but it got %s", item.p, item.r, result)
		}
	}
}
