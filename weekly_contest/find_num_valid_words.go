package weekly_contest

func initPuzzle(puzzle string) []int {
	list := make([]int, 26)
	for _, l := range puzzle {
		list[l-97] = 1
	}
	return list
}

func findNumOfValidWords(words []string, puzzles []string) []int {
	target := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		pMap := initPuzzle(puzzle)
		first := rune(puzzle[0])

		for _, word := range words {
			head := false
			match := true
			for _, l := range word {
				if l == first {
					head = true
				}

				if pMap[l-97] != 1 {
					match = false
					break
				}
			}

			if head && match {
				target[i]++
			}
		}
	}
	return target
}
