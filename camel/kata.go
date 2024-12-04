package camel

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const sepdash = "-"
const sepunderscore = "_"

func ToCamelCase(s string) string {
	if !strings.Contains(s, sepdash) && !strings.Contains(s, sepunderscore) {
		return s
	}

	var split []string

	if strings.Contains(s, sepdash) {
		split = strings.Split(s, sepdash)
	}

	if strings.Contains(s, sepunderscore) {
		split = strings.Split(s, sepunderscore)
	}

	var result = split[0]
	for i := 1; i < len(split); i++ {
		result += strings.ToUpper(string(split[i][0])) + split[i][1:]
		fmt.Println(result)
	}

	return result
}

func ToCamelCaseTwo(s string) string {
	if s == "" {
		return ""
	}
	caser := cases.Title(language.English)
	result := caser.String(strings.Replace(strings.Replace(s, "-", " ", -1), "_", " ", -1))
	result = s[:1] + result[1:]
	result = strings.Replace(result, " ", "", -1)
	return result
}
