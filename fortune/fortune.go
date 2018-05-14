package fortune

import (
	"fmt"
	"math"
)

// Fortune returns true if John can make a living until the nth year and false if it is not possible.
// f0 deposit at the beginning of year 1
// p bank's interest rate. It is constant over the years
// c john's withdraw
// n Year until John can make a living
// i percent per year of inflation. It is supposed to stay constant over the years.
func Fortune(f0 int, p float64, c0 int, n int, i float64) bool {

	percent := 100.0
	fmt.Printf("%d, ", f0)
	if n == 1 {
		fmt.Printf("\n")
		return f0 >= 0
	}

	interest := p / percent
	inflation := i / percent

	incoming := int((float64(f0) + interest*float64(f0)) - float64(c0))
	outcoming := int(float64(c0) + inflation*float64(c0))

	return Fortune(incoming, p, outcoming, n-1, i)
}

// FortuneTwo returns true if John can make a living until the nth year and false if it is not possible.
func FortuneTwo(f0 int, p float64, c0 int, n int, i float64) bool {
	fmt.Printf("%d, ", f0)
	if n == 1 {
		fmt.Printf("\n")
		return f0 >= 0
	}
	f1 := f0 + int(math.Trunc(float64(f0)*p/100.0)) - c0
	c1 := c0 + int(math.Trunc(float64(c0)*i/100.0))
	return Fortune(f1, p, c1, n-1, i)
}
