package weekly_contest_98

func comparePattern(content string, pattern string) bool {
	dict := make(map[string]string)
	reverseDice := make(map[string]string)
	for i, v := range content {
		letter := string(v)
		patternLetter := string(pattern[i])

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
