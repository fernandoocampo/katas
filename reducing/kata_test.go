package reducing_test

import (
	"testing"

	"github.com/fernandoocampo/katas/reducing"
)

type data struct {
	init  int
	f     reducing.FParam
	fname string
	dta   []int
	sol   []int
}

var tables = []data{
	{18, reducing.Gcdi, "Gcdi", []int{18, 69, -90, -78, 65, 40}, []int{18, 3, 3, 3, 1, 1}},
	{0, reducing.Som, "Som", []int{357, 112, 28, -52, 644, 119}, []int{357, 469, 497, 445, 1089, 1208}},
	{10, reducing.Maxi, "Maxi", []int{10, -32, 190, 300, -42, -38, 50, 405, -46, 225, -31}, []int{10, 10, 190, 300, 300, 300, 300, 405, 405, 405, 405}},
	{6, reducing.Lcmu, "Lcmu", []int{6, -72, -62, -22, -23, 80}, []int{6, 72, 2232, 24552, 564696, 5646960}},
	{64, reducing.Mini, "Mini", []int{64, -67, -43, 12, -15, 108, 12, 104, -36}, []int{64, -67, -67, -67, -67, -67, -67, -67, -67}},
}

func TestOpenArray(t *testing.T) {
	for _, item := range tables {
		result := reducing.OperArray(item.f, item.dta, item.init)

		for i, ritem := range item.sol {
			if result[i] != ritem {
				t.Errorf("given parameters [%d, %s, %v], it expects %v, but it got %v", item.init, item.fname, item.dta, item.sol, result)
				continue
			}
		}
	}
}
