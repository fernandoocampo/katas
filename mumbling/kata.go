package mumbling

import (
	"strings"
	"unicode"
)

func Accum(s string) string {
	// "ZpglnRxqenU", "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu"
	if s == "" {
		return s
	}
	result := ""
	for j, letter := range s {
		for i := 0; i <= j; i++ {
			if i == 0 {
				result += string(unicode.ToUpper(letter))
				continue
			}
			result += string(unicode.ToLower(letter))
		}
		result += "-"
	}

	return result[:len(result)-1]
}

func AccumTwo(s string) string {
	parts := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		parts[i] = strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i)
	}

	return strings.Join(parts, "-")
}
