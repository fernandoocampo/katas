package cycles

import (
	"fmt"
)

// Cycle returns the length of the cycle if n and 10 are coprimes, otherwise returns -1.
func Cycle(n int) int {
	length := 1
	r := 1

	for i := 1; i < n; i++ {
		r = (10 * r) % n
		fmt.Printf("%d\n", r)

		switch r {
		case 0:
			return -1
		case 1:
			return length
		default:
			length++
		}
	}

	return -1
}
