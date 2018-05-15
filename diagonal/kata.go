package diagonal

import "math/big"

// Diagonal calculates the sum of the binomial coefficients on a given diagonal.
// where n is the line where we start and p is the number of the diagonal
func Diagonal(n int, p int) int {
	if n < p || p < 0 {
		return -1
	}

	if p == 0 {
		return n + 1
	}

	matrix := make([][]int, n+1)
	matrix[0] = make([]int, n+1)
	matrix[1] = make([]int, n+1)
	matrix[0][0] = 1
	matrix[1][0] = 1
	matrix[1][1] = 1

	sum := matrix[0][p] + matrix[1][p]

	for i := 2; i < n+1; i++ {
		matrix[i] = make([]int, n+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == n {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = matrix[i-1][j-1] + matrix[i-1][j]
				if j == p {
					sum += matrix[i][j]
				}
			}
		}
	}
	return sum
}

// DiagonalTwo calculates the sum of the binomial coefficients on a given diagonal.
// where n is the line where we start and p is the number of the diagonal
func DiagonalTwo(n int, p int) int {
	var z big.Int
	return int(z.Binomial(int64(n+1), int64(p+1)).Int64())
}

// DiagonalThree calculates the sum of the binomial coefficients on a given diagonal.
// where n is the line where we start and p is the number of the diagonal
func DiagonalThree(n, p int) int {
	sum := 1
	num := 1
	for i := 1; i < n-p+1; i++ {
		num = num * (p + i) / (i)
		sum += num
	}
	return sum
}
