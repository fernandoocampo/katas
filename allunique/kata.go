package allunique

import "strings"

func HasUniqueChar(str string) bool {
	if len(str) == 1 {
		return true
	}

	if !strings.Contains(str[1:], string(str[0])) {
		return HasUniqueChar(str[1:])
	} else {
		return false
	}
}

func HasUniqueCharTwo(str string) bool {
	var m = map[rune]bool{}
	for _, w := range str {
		m[w] = true
	}
	return len(m) == len(str)
}
