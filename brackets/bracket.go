package brackets

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

// WhichIsThree returns the index that is the ']' of the given index that is a '['
func WhichIsThree(text string, index int) int {
	result := -1

	if text[index] != '[' {
		return -1
	}

	findcounter := 0
	for i := index; i < len(text); i++ {
		if text[i] == '[' {
			findcounter++
			continue
		}
		if text[i] == ']' {
			findcounter--
		}
		if findcounter == 0 {
			result = i
			break
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
