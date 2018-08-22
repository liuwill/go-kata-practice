package weekly_contest_98

import "strings"

func comparePattern(content string, pattern string) bool {
	return strings.Compare(content, pattern) == 0
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
