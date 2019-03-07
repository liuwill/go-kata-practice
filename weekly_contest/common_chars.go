package weekly_contest

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

		if count == len(A) {
			target = append(target, string(letter))
		}
	}

	return target
}
