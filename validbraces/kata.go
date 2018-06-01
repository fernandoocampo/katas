package validbraces

func ValidBraces(str string) bool {
	if len(str) == 0 {
		return true
	}

	stack := []rune{}

	for _, item := range str {
		switch item {
		case '(', '[', '{':
			stack = append(stack, item)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		default:
			continue
		}
	}
	return len(stack) == 0
}

var m = map[string]string{
	"{": "}",
	"(": ")",
	"[": "]",
}

func ValidBracesTwo(str string) bool {
	s := make([]string, 0)
	for _, r := range str {
		r := string(r)
		if len(s) > 0 && m[s[len(s)-1]] == r {
			s = s[:len(s)-1]
		} else {
			s = append(s, r)
		}
	}
	return len(s) == 0
}
