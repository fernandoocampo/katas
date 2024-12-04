package bandname

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func BandNameGenerator(word string) string {
	if word == "" {
		return word
	}

	the := "the "

	caser := cases.Title(language.English)

	if len(word) > 4 && strings.EqualFold(word[0:4], the) {
		return caser.String(word)
		// return strings.Title(word)
	}

	if word[0] == word[len(word)-1] {
		return caser.String(word + word[1:])
	}

	return caser.String(the + word)
}

func BandNameGeneratorTwo(word string) string {
	first := word[:1]
	last := word[len(word)-1:]

	caser := cases.Title(language.English)

	if first != last {
		return "The " + caser.String(word)
	}

	return caser.String(word) + word[1:]

}
