package nato

import (
	"bytes"
	"strings"
	"unicode"
)

// ToNato translate a string to Pilot's alphabet (NATO phonetic alphabet)
func ToNato(words string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India",
		"Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra",
		"Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}

	var buffer bytes.Buffer

	for _, word := range words {
		if string(word) != "" && string(word) != " " {
			buffer.WriteString(findWord(nato, string(word)))
			buffer.WriteString(" ")
		}
	}
	result := buffer.String()
	result = result[0:(len(result) - 1)]

	return result
}

func findWord(words []string, key string) string {
	for _, word := range words {
		if strings.EqualFold(string(word[0]), key) {
			return word
		}
	}
	return key
}

// ToNatoTwo translate a string to Pilot's alphabet (NATO phonetic alphabet) in a best way
func ToNatoTwo(words string) string {
	nato := []string{
		"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot",
		"Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November",
		"Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor",
		"Whiskey", "Xray", "Yankee", "Zulu",
	}
	charToCharlie := map[rune]string{}
	for _, value := range nato {
		charToCharlie[rune(value[0])] = value
	}

	result := ""
	for _, letter := range words {
		if unicode.IsLetter(letter) {
			result += charToCharlie[unicode.ToUpper(letter)] + " "
		} else if unicode.IsPunct(letter) {
			result += string(letter)
		}
	}
	return strings.TrimSpace(result)
}
