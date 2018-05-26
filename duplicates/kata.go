package duplicates

import (
	"strings"
)

func Duplicate_count(s1 string) int {
	counter := 0
	recorder := ""

	s2 := strings.ToLower(s1)

	for _, item := range s2 {
		if strings.ContainsRune(recorder, item) {
			continue
		}
		if strings.Count(s2, string(item)) > 1 {
			counter++
		}
		recorder += string(item)
	}

	return counter
}

func Duplicate_count_two(s string) (c int) {
	h := map[rune]int{}
	for _, r := range strings.ToLower(s) {
		if h[r]++; h[r] == 2 {
			c++
		}
	}
	return
}
