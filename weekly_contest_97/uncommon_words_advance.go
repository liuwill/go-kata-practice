package weekly_contest_97

func checkWordInDict(word string, result []string) []string {
	for index, item := range result {
		if item == word {
			result = append(result[:index], result[index+1:]...)
			return result
		}
	}
	result = append(result[:], word)
	return result
}

func scanSentence(sentence string, result []string) []string {
	word := ""
	for _, v := range sentence {
		letter := string(v)

		if letter != " " {
			word += letter
		} else {
			result = checkWordInDict(word, result)
			word = ""
		}
	}

	if len(word) > 0 {
		result = checkWordInDict(word, result)
	}

	return result
}

func uncommonFromSentencesAdvance(A string, B string) []string {
	result := []string{}
	result = scanSentence(A, result)
	result = scanSentence(B, result)

	return result
}
