package weekly_contest

/**
 * daily-challenge-1023
 * PUZZLE: Camelcase Matching
 */
func camelMatch(queries []string, pattern string) []bool {
	target := make([]bool, len(queries))
	// patternList := []rune(pattern)
	for pos := 0; pos < len(queries); pos++ {
		item := queries[pos]

		p := 0
		answer := false
		i := 0
		for ; i < len(item) && p < len(pattern); i++ {
			letter := item[i]
			current := pattern[p]

			if letter == current {
				p++
			} else if item[i] < 'a' || item[i] > 'z' {
				break
			}
		}

		if p == len(pattern) {
			answer = true
		}

		if i < len(item) && answer {
			for ; i < len(item); i++ {
				if item[i] < 'a' || item[i] > 'z' {
					answer = false
				}
			}
		}
		target[pos] = answer

	}
	return target
}
