package ages_test

import (
	"testing"

	"github.com/fernandoocampo/katas/ages"
)

func TestTwoOldestAges(t *testing.T) {
	tables := []struct {
		a []int
		b [2]int
	}{
		{[]int{6, 5, 83, 5, 3, 18}, [2]int{18, 83}},
		{[]int{1, 5, 87, 45, 8, 8}, [2]int{45, 87}},
		{[]int{39, 53, 83, 51, 59, 61, 95, 23, 99, 49}, [2]int{95, 99}},
		{[]int{43, 25, 83, 11, 31, 91, 85, 25, 95, 65}, [2]int{91, 95}},
		{[]int{19, 5, 43, 13, 75, 89, 43, 89, 25, 49}, [2]int{89, 89}},
	}

	for _, item := range tables {
		result := ages.TwoOldestAges(item.a)

		if item.b[0] != result[0] || item.b[1] != result[1] {
			t.Errorf("for %v, was expected: %v, but it got %v", item.a, item.b, result)
		}
	}
}

func TestTwoOldestAgesTwo(t *testing.T) {
	tables := []struct {
		a []int
		b [2]int
	}{
		{[]int{6, 5, 83, 5, 3, 18}, [2]int{18, 83}},
		{[]int{1, 5, 87, 45, 8, 8}, [2]int{45, 87}},
		{[]int{39, 53, 83, 51, 59, 61, 95, 23, 99, 49}, [2]int{95, 99}},
		{[]int{43, 25, 83, 11, 31, 91, 85, 25, 95, 65}, [2]int{91, 95}},
		{[]int{19, 5, 43, 13, 75, 89, 43, 89, 25, 49}, [2]int{89, 89}},
	}

	for _, item := range tables {
		result := ages.TwoOldestAgesTwo(item.a)

		if item.b[0] != result[0] || item.b[1] != result[1] {
			t.Errorf("for %v, was expected: %v, but it got %v", item.a, item.b, result)
		}
	}
}

func TestTwoOldestAgesThree(t *testing.T) {
	tables := []struct {
		a []int
		b [2]int
	}{
		{[]int{6, 5, 83, 5, 3, 18}, [2]int{18, 83}},
		{[]int{1, 5, 87, 45, 8, 8}, [2]int{45, 87}},
		{[]int{39, 53, 83, 51, 59, 61, 95, 23, 99, 49}, [2]int{95, 99}},
		{[]int{43, 25, 83, 11, 31, 91, 85, 25, 95, 65}, [2]int{91, 95}},
		{[]int{19, 5, 43, 13, 75, 89, 43, 89, 25, 49}, [2]int{89, 89}},
	}

	for _, item := range tables {
		result := ages.TwoOldestAgesThree(item.a)

		if item.b[0] != result[0] || item.b[1] != result[1] {
			t.Errorf("for %v, was expected: %v, but it got %v", item.a, item.b, result)
		}
	}
}

func TestTwoOldestAgesFour(t *testing.T) {
	tables := []struct {
		a []int
		b [2]int
	}{
		{[]int{6, 5, 83, 5, 3, 18}, [2]int{18, 83}},
		{[]int{1, 5, 87, 45, 8, 8}, [2]int{45, 87}},
		{[]int{39, 53, 83, 51, 59, 61, 95, 23, 99, 49}, [2]int{95, 99}},
		{[]int{43, 25, 83, 11, 31, 91, 85, 25, 95, 65}, [2]int{91, 95}},
		{[]int{19, 5, 43, 13, 75, 89, 43, 89, 25, 49}, [2]int{89, 89}},
	}

	for _, item := range tables {
		result := ages.TwoOldestAgesFour(item.a)

		if item.b[0] != result[0] || item.b[1] != result[1] {
			t.Errorf("for %v, was expected: %v, but it got %v", item.a, item.b, result)
		}
	}
}
