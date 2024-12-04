package balance

const openparens rune = '('
const opensquares rune = '['
const opencurlies rune = '{'
const closeparens rune = ')'
const closesquares rune = ']'
const closecurlies rune = '}'

var keys = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}
var values = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func Balance(data string) bool {
	openbracket := make([]rune, 0)

	for _, value := range data {
		if value == openparens || value == opensquares || value == opencurlies {
			openbracket = append(openbracket, value)
			continue
		}

		switch value {
		case closeparens:
			openbracket = closeBracket(openparens, openbracket)
		case closesquares:
			openbracket = closeBracket(opensquares, openbracket)
		case closecurlies:
			openbracket = closeBracket(opencurlies, openbracket)
		}
	}
	if len(openbracket) == 0 {
		return true
	}

	return false
}

func closeBracket(expectedvalue rune, brackets []rune) []rune {
	if len(brackets) > 0 && brackets[len(brackets)-1] == expectedvalue {
		return brackets[:len(brackets)-1]
	}
	return brackets
}

func BalanceParentheses(value string) bool {
	if value == "" {
		return true
	}

	if value[0] == ')' {
		return false
	}

	balance := 0
	for _, v := range value {
		if v == ')' {
			balance--
		} else {
			balance++
		}
		if balance < 0 {
			return false
		}
	}

	return balance == 0
}

func BalanceTwo(data string) bool {
	if data == "" {
		return true
	}

	var stack []rune

	for _, v := range data {
		if _, ok := keys[v]; ok {
			stack = append(stack, v)
			continue
		}

		key, ok := values[v]
		if !ok {
			continue
		}

		if len(stack) == 0 && ok {
			return false
		}

		if len(stack) > 0 && stack[len(stack)-1] != key {
			return false
		}

		if len(stack) > 0 && stack[len(stack)-1] == key {
			stack = stack[:len(stack)-1]
			continue
		}
	}

	return len(stack) == 0
}
