package weekly_contest

import "sort"

/**
 * weekly-contest-1033
 * PUZZLE: Moving Stones Until Consecutive
 */
func numMovesStones(a int, b int, c int) []int {
	min, max := 0, 0
	sortList := []int{a, b, c}
	sort.Ints(sortList)
	if sortList[0] == sortList[1]-1 && sortList[1] == sortList[2]-1 {
		min = 0
	} else if sortList[1]-sortList[0] <= 2 || sortList[2]-sortList[1] <= 2 {
		min = 1
	} else {
		min = 2
	}

	max = (sortList[2] - sortList[1] - 1) + (sortList[1] - sortList[0] - 1)

	return []int{min, max}
}
