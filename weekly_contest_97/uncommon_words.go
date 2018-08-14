package weekly_contest_97

func detectSentence(sentence string, dict map[string]int) {
	word := ""
	for _, v := range sentence {
		letter := string(v)

		if letter != " " {
			word += letter
		} else {
			if _, ok := dict[word]; ok {
				dict[word]++
			} else {
				dict[word] = 1
			}
			word = ""
		}
	}

	if len(word) > 0 {
		if _, ok := dict[word]; ok {
			dict[word]++
		} else {
			dict[word] = 1
		}
		word = ""
	}
}

func uncommonFromSentences(A string, B string) []string {
	dict := make(map[string]int)
	detectSentence(A, dict)
	detectSentence(B, dict)

	result := []string{}
	for k, v := range dict {
		if v == 1 {
			result = append(result[:], k)
		}
	}

	return result
}
