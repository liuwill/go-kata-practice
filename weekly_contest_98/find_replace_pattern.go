package weekly_contest_98

func comparePattern(content string, pattern string) bool {
	dict := make(map[uint8]uint8)
	reverseDice := make(map[uint8]uint8)
	for i, v := range content {
		letter := uint8(v)
		patternLetter := uint8(pattern[i])

		if _, ok := dict[letter]; !ok {
			dict[letter] = patternLetter
		}

		if _, ok := reverseDice[patternLetter]; !ok {
			reverseDice[patternLetter] = letter
		}

		if reverseDice[patternLetter] != letter || dict[letter] != patternLetter {
			return false
		}
	}
	return true
}

func findAndReplacePattern(words []string, pattern string) []string {
	result := []string{}

	for _, value := range words {
		if comparePattern(value, pattern) {
			result = append(result, value)
		}
	}
	return result
}
