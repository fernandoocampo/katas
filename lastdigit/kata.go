package lastdigit

import (
	"fmt"
	"math"
)

func LastDigit(as []int) int {
	result := pow(as)

	fmt.Printf("%f\n", result)

	if result < 0 {
		result = 1
	}

	return int(int64(result) % int64(10))

}

func pow(as []int) float64 {
	if len(as) == 0 {
		return 1
	}

	if len(as) == 1 {
		return float64(as[0])
	}

	if len(as) == 2 {
		return math.Pow(float64(as[0]), float64(as[1]))
	}
	return math.Pow(float64(as[0]), float64(pow(as[1:])))
}
