package weekly_contest

/**
 * daily-challenge-1002
 * PUZZLE: Find Common Characters
 */
func commonChars(A []string) []string {
	countList := make([]int, 26)
	countMap := make([][]int, len(A))
	FIRST_LETTER := 'a'
	for i, list := range A {
		countMap[i] = make([]int, 26)
		for _, letter := range list {
			pos := letter - FIRST_LETTER
			countMap[i][pos] += 1
		}
	}

	for i, _ := range countList {
		countList[i] = countMap[0][i]
		for j := 1; j < len(countMap); j++ {
			if countList[i] > countMap[j][i] {
				countList[i] = countMap[j][i]
			}

			if countList[i] == 0 {
				break
			}
		}
	}

	target := []string{}
	for pos, count := range countList {
		letter := rune(pos) + FIRST_LETTER

		left := count
		for left > 0 {
			target = append(target, string(letter))
			left -= 1
		}
	}

	return target
}
