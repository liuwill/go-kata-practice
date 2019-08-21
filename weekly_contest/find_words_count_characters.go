package weekly_contest

func countCharacters(words []string, chars string) int {
	letterMap := make(map[rune]int)

	for _, c := range chars {
		if _, ok := letterMap[c]; !ok {
			letterMap[c] = 0
		}
		letterMap[c]++
	}

	count := 0
	for _, word := range words {
		match := 0
		for _, l := range word {
			if _, ok := letterMap[l]; !ok || letterMap[l] <= 0 {
				match = 0
				break
			}

			match++
			letterMap[l]--
		}

		count += match
	}

	return count
}
