package multiples

// Multiple3And5  returns the sum of all the multiples of 3 or 5 below the number passed in.
func Multiple3And5(number int) int {
	var t, f, sum int
	values := make(map[int]int)

	// multiples of 3
	for i := 1; t < number || f < number; i++ {
		t = 3 * i
		f = 5 * i
		if t < number {
			values[t] = t
		}
		if f < number {
			values[f] = f
		}
	}

	for _, value := range values {
		sum += value
	}

	return sum
}

// Multiple3And5Two as Multiple3And5 but efficiently
func Multiple3And5Two(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	return sum
}
