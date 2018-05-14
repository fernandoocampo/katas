package brackets

import (
	"fmt"
)

// WhichIs returns the index that is the ']' of the given index that is a '['
func WhichIs(text string, index int) int {
	result := -1

	if text[index] != '[' {
		return -1
	}

	findcounter := -1
	for i, letter := range text {
		if letter == '[' {
			if i > index {
				findcounter++
			} else {
				findcounter = 0
			}
		}
		if letter == ']' {
			if i > index && findcounter == 0 {
				result = i
				break
			} else {
				findcounter--
			}
		}

	}

	return result
}

// WhichIsTwo returns the index that is the ']' of the given index that is a '['
func WhichIsTwo(text string, index int) int {
	result := -1

	if text[index] != '[' {
		return -1
	}

	findcounter := 0
	for i, letter := range text[index:] {
		fmt.Println("hola", i)
		if letter == '[' {
			findcounter++
		}
		if letter == ']' {
			findcounter--
		}

		if findcounter == 0 {
			result = i + index
			break
		}

	}

	return result
}
