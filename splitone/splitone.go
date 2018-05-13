package splitone

import "regexp"

// Solution receives a string and returns the string into pairs of two characters.
func Solution(str string) []string {

	if str == "" {
		return nil
	}

	isodd := !((len(str) % 2) == 0)

	var pair string
	var result = make([]string, 0)
	charcounter := 0
	for i, charitem := range str {
		pair += string(charitem)
		charcounter++
		if charcounter > 1 {
			charcounter = 0
			result = append(result, pair)
			pair = ""
		}

		if isodd && i == len(str)-1 {
			result = append(result, pair)
		}
	}

	if isodd {
		result[len(result)-1] += "_"
	}

	return result
}

func SolutionTwo(str string) []string {
	var res []string
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 {
		res = append(res, str[i:i+2])
	}
	return res
}

func SolutionThree(str string) []string {
	return regexp.MustCompile("..?").FindAllString(str+"_"[:len(str)%2], -1)
}
