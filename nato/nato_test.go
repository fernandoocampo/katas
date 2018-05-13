package nato_test

import (
	"testing"

	"github.com/fernandoocampo/katas/nato"
)

func TestToNato(t *testing.T) {
	tables := []struct {
		x string
		y string
	}{
		{"If you can read", "India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"},
		{"Did not see that coming", "Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf"},
		{"go for it!", "Golf Oscar Foxtrot Oscar Romeo India Tango !"},
		{"Miloso", "Mike India Lima Oscar Sierra Oscar"},
	}

	for _, item := range tables {
		result := nato.ToNato(item.x)

		if result != item.y {
			t.Errorf("for %s, we expect %s, but we got %s", item.x, item.y, result)
		}
	}
}

func TestToNatoTwo(t *testing.T) {
	tables := []struct {
		x string
		y string
	}{
		{"If you can read", "India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"},
		{"Did not see that coming", "Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf"},
		{"go for it!", "Golf Oscar Foxtrot Oscar Romeo India Tango !"},
		{"Miloso", "Mike India Lima Oscar Sierra Oscar"},
	}

	for _, item := range tables {
		result := nato.ToNatoTwo(item.x)

		if result != item.y {
			t.Errorf("for %s, we expect %s, but we got %s", item.x, item.y, result)
		}
	}
}
