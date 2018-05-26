package reducing

import "math"

func Gcdi(x, y int) int {
	x = int(math.Abs(float64(x)))
	y = int(math.Abs(float64(y)))
	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}

func Som(x, y int) int {
	return x + y
}
func Maxi(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}
func Mini(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}
func Lcmu(x, y int) int {
	x = int(math.Abs(float64(x)))
	y = int(math.Abs(float64(y)))
	result := x * y / Gcdi(x, y)

	return result
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {
	result := make([]int, len(arr))
	result[0] = f(init, arr[0])

	for i := 1; i < len(arr); i++ {
		result[i] = f(result[i-1], arr[i])
	}

	return result
}
