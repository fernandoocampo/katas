package divisibility

import "fmt"

// Thirt Multiply the right most digit of the number with the left most
// number in the sequence shown above, the second right most digit to
// the second left most digit of the number in the sequence.
func Thirt(n int) int {
	sum := 0
	number := n

	for i := 1; number > 0; i *= 10 {
		fmt.Println(number % 10)
		fmt.Println(i % 13)
		sum += (i % 13) * (number % 10)
		number /= 10
	}

	if n == sum {
		return n
	}

	return Thirt(sum)
}
