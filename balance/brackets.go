package balance

const openparens rune = '('
const opensquares rune = '['
const opencurlies rune = '{'
const closeparens rune = ')'
const closesquares rune = ']'
const closecurlies rune = '}'

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
