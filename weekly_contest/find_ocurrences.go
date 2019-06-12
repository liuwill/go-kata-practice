package weekly_contest

import "strings"

/**
 * weekly-contest-1078
 * PUZZLE: Occurrences After Bigram
 */
func findOcurrences(text string, first string, second string) []string {
	result := []string{}
	wordList := strings.Split(text, " ")
	for i := 0; i < len(wordList)-1; i++ {
		if wordList[i] != first || wordList[i+1] != second {
			continue
		} else if i+2 > len(wordList)-1 {
			continue
		}

		result = append(result, wordList[i+2])
	}

	return result
}
