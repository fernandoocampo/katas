package multiline

import (
	"bufio"
	"regexp"
	"strings"
)

type book struct {
	title       string
	author      string
	description string
}

var db map[string]book

func Index(library *string) {
	db = make(map[string]book)
	newstr := strings.NewReader(*library)
	b := bufio.NewReader(newstr)

	newbook := new(book)

	for {
		line, err := b.ReadString('\n')
		if err != nil {
			break
		}

		if line == "\n" || line == "" {
			db[newbook.title] = *newbook
			newbook = new(book)
			continue
		}

		if len(line) >= 7 && line[:7] == "TITLE: " {
			newbook.title = line[6:]
			continue
		}
		if len(line) >= 7 && line[:8] == "AUTHOR: " {
			newbook.author = line[8:]
			continue
		}
		if len(line) >= 7 && line[:13] == "DESCRIPTION: " {
			newbook.description = line[12:]
			continue
		}

		newbook.description += " " + line

	}
}

func Search(word string) []string {
	result := []string{}
	r := regexp.MustCompile(`(?i)\b` + word + `\b`)
	for key, value := range db {
		if r.MatchString(key) || r.MatchString(value.author) || r.MatchString(value.description) {
			result = append(result, key)
			continue
		}

	}
	return result
}
