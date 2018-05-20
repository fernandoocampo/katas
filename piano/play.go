package piano

// BlackOrWhiteKey will receive an integer between 1 and 10000
// and return the string "black" or "white"
func BlackOrWhiteKey(keyPressCount int) string {
	var result string
	//var color string
	counter := 0

	cycles := keyPressCount / 88
	if cycles%88 != 0 || cycles == 0 {
		cycles++
	}

	for m := 1; m <= int(cycles); m++ {

		for n := 1; n <= 3; n++ {
			if (n + counter) == keyPressCount {
				if n%2 == 0 {
					return "black"
				}
				return "white"
			}
		}

		counter += 3

		for k := 1; k <= 7; k++ {

			iterationResult := iteratePattern(keyPressCount, counter, 7)

			if iterationResult != "" {
				return iterationResult
			}

			counter += 5

			iterationResult = iteratePattern(keyPressCount, counter, 9)
			if iterationResult != "" {
				return iterationResult
			}
			counter += 7
		}
		counter++
		if counter == keyPressCount {
			return "white"
		}

	}
	return result
}

func iteratePattern(keyPressCount int, counter int, keys int) string {
	for i, j := 1, 2; i < (keys - 1); i, j = i+2, j+2 {
		if (i + counter) == keyPressCount {
			return "white"
		}

		if (j + 2) < keys {
			if (j + counter) == keyPressCount {
				return "black"
			}
		}
	}

	return ""
}

const (
	w = "white"
	b = "black"
)

// BlackOrWhiteKeyTwo will receive an integer between 1 and 10000
// and return the string "black" or "white"
func BlackOrWhiteKeyTwo(key int) string {
	keys := []string{w, b, w, w, b, w, b, w, w, b, w, b}
	return keys[((key-1)%88)%len(keys)]
}
