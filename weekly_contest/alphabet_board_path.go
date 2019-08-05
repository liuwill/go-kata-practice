package weekly_contest

import "strings"

func doMove(source []int, target []int) string {
	vLetter := ""
	pLetter := ""

	vDistance := target[0] - source[0]
	pDistance := target[1] - source[1]

	if vDistance > 0 {
		vLetter = strings.Repeat("D", vDistance)
	} else if vDistance < 0 {
		vLetter = strings.Repeat("U", vDistance*-1)
	}

	if pDistance > 0 {
		pLetter = strings.Repeat("R", pDistance)
	} else if pDistance < 0 {
		pLetter = strings.Repeat("L", pDistance*-1)
	}

	return vLetter + pLetter + "!"
}

func alphabetBoardPath(target string) string {
	start := 'a'
	board := []string{"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"}

	sourceMap := [][]int{}

	for i, line := range board {
		for j, _ := range line {
			sourceMap = append(sourceMap, []int{i, j})
		}
	}

	action := ""
	current := []int{0, 0}
	for _, letter := range target {
		index := letter - start
		position := sourceMap[index]

		moves := doMove(current, position)

		current = position
		action += moves
	}
	return action
}
