package weekly_contest

import "strings"

func doMove(source []int, target []int) string {
	vLetter := ""
	pLetter := ""

	vDistance := target[0] - source[0]
	pDistance := target[1] - source[1]

	vNum := 0
	pNum := 0

	if vDistance > 0 {
		vLetter = "D"
		vNum = vDistance
	} else if vDistance < 0 {
		vLetter = "U"
		vNum = vDistance * -1
	}

	if pDistance > 0 {
		pLetter = "R"
		pNum = pDistance
	} else if pDistance < 0 {
		pLetter = "L"
		pNum = pDistance * -1
	}

	zMark := false
	if target[0] == 5 && pNum > 0 && vNum > 0 {
		zMark = true
		vNum--
	}

	result := ""
	if vNum > 0 {
		result += strings.Repeat(vLetter, vNum)
	}
	if pNum > 0 {
		result += strings.Repeat(pLetter, pNum)
	}
	if zMark {
		result += vLetter
	}

	return result + "!"
}

/**
 * daily-challenge-1138
 * PUZZLE: Alphabet Board Path
 */
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
