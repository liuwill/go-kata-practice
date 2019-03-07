package weekly_contest

/**
 * daily-challenge-1002
 * PUZZLE: Find Common Characters
 */
func commonChars(A []string) []string {
	countList := make([]int, 26)
	FIRST_LETTER := 'a'
	for _, list := range A {
		for _, letter := range list {
			pos := letter - FIRST_LETTER
			countList[pos] += 1
		}
	}

	target := []string{}
	for pos, count := range countList {
		letter := rune(pos) + FIRST_LETTER

		left := count
		for left >= len(A) {
			target = append(target, string(letter))
			left -= len(A)
		}
	}

	return target
}
