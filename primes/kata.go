package primes

import (
	"math"
)

// Step returns the first primes couple that have a gap of g in the range between m and n.
// 4,100,110, []int{103, 107}
func Step(g, m, n int) []int {

	// checks that g and m must be always greater than 2
	if g < 2 && m < 2 && n < m {
		return nil
	}

	var result = make([]int, 2)

	for j := m; j <= n; j++ {
		if isPrime(j) {
			result[0] = j
			for i := j + 1; i <= n; i++ {
				if isPrime(i) {
					if (i - result[0]) == g {
						result[1] = i
						return result
					}
					if (i - result[0]) > g {
						break
					}
				}
			}
		}
	}

	return nil
}

// Step2 simple way from codewars
func Step2(g, m, n int) []int {
	for i := m; i <= n; i++ {
		if isPrime(i) && i+g <= n && isPrime(i+g) {
			return []int{i, i + g}
		}
	}

	return nil
}

// isPrime calculates if the given value is a prime number
func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
