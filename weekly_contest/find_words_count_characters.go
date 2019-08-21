package weekly_contest

/**
 * daily-challenge-1160
 * PUZZLE: Find Words That Can Be Formed by Characters
 */
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

		currentMap := make(map[rune]int)
		for k, v := range letterMap {
			currentMap[k] = v
		}

		for _, l := range word {
			if _, ok := currentMap[l]; !ok || currentMap[l] <= 0 {
				match = 0
				break
			}

			match++
			currentMap[l]--
		}

		count += match
	}

	return count
}

func parseCharactersWord(word string) []int {
	meta := make([]int, 26)
	for _, l := range word {
		meta[l-'a']++
	}
	return meta
}

func countWordMatchs(meta []int, word string) int {
	count := 0
	current := make([]int, 26)
	for _, l := range word {
		pos := l - 'a'

		current[pos]++
		if current[pos] > meta[pos] {
			return 0
		}
		count++
	}
	return count
}

/**
 * Smart Answer
 */
func countCharactersSmart(words []string, chars string) int {
	meta := parseCharactersWord(chars)

	count := 0
	for _, word := range words {
		count += countWordMatchs(meta, word)
	}
	return count
}
