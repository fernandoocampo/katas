package balance_test

import (
	"testing"

	"github.com/fernandoocampo/katas/balance"
)

type data struct {
	param  string
	result bool
}

var tables = []data{
	{"(a[0]+b[2c[6]]) {24 + 53}", true},
	{"f(e(d))", true},
	{"[()]{}([])", true},
	{"((b)", false},
	{"(c]", false},
	{"{(a[])", false},
	{"([)]", false},
	{"[({{}}])", false},
	{")(", false},
}

func TestBalance(t *testing.T) {
	for _, item := range tables {
		result := balance.Balance(item.param)
		if result != item.result {
			t.Errorf("Given parameter %s, it expects %t, but it got %t", item.param, item.result, result)
		}
	}
}

func TestBalanceTwo(t *testing.T) {
	for _, item := range tables {
		result := balance.BalanceTwo(item.param)
		if result != item.result {
			t.Errorf("Given parameter %s, it expects %t, but it got %t", item.param, item.result, result)
		}
	}
}

func TestOnlyParentheses(t *testing.T) {
	testData := map[string]struct {
		value string
		want  bool
	}{
		"()": {
			value: "()",
			want:  true,
		},
		")(": {
			value: ")(",
			want:  false,
		},
		"())))()()": {
			value: "())))()()",
			want:  false,
		},
		"())))((((": {
			value: "())))((((",
			want:  false,
		},
	}

	for testName, testData := range testData {
		// given
		t.Run(testName, func(t *testing.T) {
			got := balance.BalanceParentheses(testData.value)
			if got != testData.want {
				t.Errorf("want: %t, but got: %t", testData.want, got)
			}
		})
	}
}
