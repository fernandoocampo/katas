package ages

import "slices"

// TwoOldestAges should take an array of numbers as its argument and return the two highest numbers within the array.
func TwoOldestAges(ages []int) [2]int {
	if ages == nil || len(ages) < 2 {
		return [2]int{0, 0}
	}

	var result [2]int
	highestIndex := 0

	for i, val := range ages {
		if val > result[1] {
			result[1] = val
			highestIndex = i
		}
	}

	for i, val := range ages {
		if val > result[0] && i != highestIndex {
			result[0] = val
		}
	}

	return result
}

func TwoOldestAgesTwo(ages []int) [2]int {
	a, b := 0, 0
	for _, v := range ages {
		if v > b {
			a, b = b, v
		} else if v > a {
			a = v
		}
	}
	return [2]int{a, b}
}

func TwoOldestAgesThree(ages []int) [2]int {
	a, b := 0, 0
	for _, v := range ages {
		if v > b {
			a, b = b, v
			continue
		}
		if v > a {
			a = v
		}
	}
	return [2]int{a, b}
}

func TwoOldestAgesFour(ages []int) [2]int {
	var result [2]int
	if len(ages) == 0 {
		return result
	}
	if len(ages) == 1 {
		result[0] = ages[0]
		return result
	}

	slices.Sort(ages)

	return [2]int{ages[len(ages)-2], ages[len(ages)-1]}
}
