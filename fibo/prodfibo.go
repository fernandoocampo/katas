package fibo

// ProductFib takes an integer (prod) and returns an array [F(n), F(n+1), true] or {F(n), F(n+1), 1} or (F(n), F(n+1), True)
func ProductFibTwo(prod uint64) [3]uint64 {
	if prod == 0 {
		return [3]uint64{0, 0, 1}
	}

	if prod == 1 {
		return [3]uint64{1, 1, 1}
	}

	var left, right uint64

	for i := 0; ; i++ {
		left = uint64(FibonacciRecursion(i))
		right = uint64(FibonacciRecursion(i + 1))

		if prod == (left * right) {
			return [3]uint64{left, right, 1}
		}

		if (left * right) > prod {
			return [3]uint64{left, right, 0}
		}

	}
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func ProductFib(prod uint64) [3]uint64 {
	f1 := uint64(0)
	f2 := uint64(1)
	is := uint64(0)
	for f1*f2 < prod {
		f1, f2 = f2, f1+f2
	}
	if prod == f1*f2 {
		is = 1
	}
	return [3]uint64{f1, f2, is}
}
