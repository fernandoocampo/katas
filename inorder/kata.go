package inorder

import "sort"

// InAscOrder recursively iterate the given array comparing the first
// two indexes, if the array contains less than two items, it returns
// true
func InAscOrder(numbers []int) bool {

	if len(numbers) <= 1 {
		return true
	}

	if numbers[0] < numbers[1] {
		return InAscOrder(numbers[1:])
	}
	return false
}

func InAscOrderTwo(numbers []int) bool {
	return sort.IntsAreSorted(numbers)
}
