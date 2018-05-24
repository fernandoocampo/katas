package mumbling_test

import (
	"testing"

	"github.com/fernandoocampo/katas/mumbling"
)

type data struct {
	p string
	r string
}

var tables = []data{
	{"ZpglnRxqenU", "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu"},
	{"NyffsGeyylB", "N-Yy-Fff-Ffff-Sssss-Gggggg-Eeeeeee-Yyyyyyyy-Yyyyyyyyy-Llllllllll-Bbbbbbbbbbb"},
	{"MjtkuBovqrU", "M-Jj-Ttt-Kkkk-Uuuuu-Bbbbbb-Ooooooo-Vvvvvvvv-Qqqqqqqqq-Rrrrrrrrrr-Uuuuuuuuuuu"},
	{"EvidjUnokmM", "E-Vv-Iii-Dddd-Jjjjj-Uuuuuu-Nnnnnnn-Oooooooo-Kkkkkkkkk-Mmmmmmmmmm-Mmmmmmmmmmm"},
	{"HbideVbxncC", "H-Bb-Iii-Dddd-Eeeee-Vvvvvv-Bbbbbbb-Xxxxxxxx-Nnnnnnnnn-Cccccccccc-Ccccccccccc"},
}

func TestAccum(t *testing.T) {
	for _, item := range tables {
		result := mumbling.Accum(item.p)

		if result != item.r {
			t.Errorf("Given parameter %s it expects %s, but it got %s", item.p, item.r, result)
		}
	}
}

func TestAccumTwo(t *testing.T) {
	for _, item := range tables {
		result := mumbling.AccumTwo(item.p)

		if result != item.r {
			t.Errorf("Given parameter %s it expects %s, but it got %s", item.p, item.r, result)
		}
	}
}
