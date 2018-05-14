package constructing

// StringConstructing finds the least number of operations needed to construct the given string s.
// a -> string to be appended. Contains only lowercase English letters. 1 <= a.length <= 20
// s -> desired string containing only lowercase English letters.
func StringConstructing(a, s string) int {
	var i, ops int
	var curr string

	for curr != s {
		for i < len(curr) && i < len(s) && s[i] == curr[i] {
			i++
		}

		if i >= len(curr) {
			curr += a
		} else {
			curr = curr[:i] + curr[i+1:len(curr)]
		}

		ops++
	}

	return ops
}
