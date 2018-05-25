package bandname

import (
	"strings"
)

func BandNameGenerator(word string) string {
	if word == "" {
		return word
	}

	the := "the "

	if len(word) > 4 && strings.EqualFold(word[0:4], the) {
		return strings.Title(word)
	}

	if word[0] == word[len(word)-1] {
		return strings.Title(word + word[1:])
	}

	return strings.Title(the + word)
}

func BandNameGeneratorTwo(word string) string {
	first := word[:1]
	last := word[len(word)-1:]

	if first != last {
		return "The " + strings.Title(word)
	}

	return strings.Title(word) + word[1:]

}
