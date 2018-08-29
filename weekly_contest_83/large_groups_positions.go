package weekly_contest_83

func isMatchLarge(start int, end int) bool {
	return end-start >= 2
}

func largeGroupPositions(S string) [][]int {
	target := [][]int{}
	current := '#'
	start := 0
	end := 0
	for i, v := range S {
		if v == current {
			end = i
			continue
		}

		if isMatchLarge(start, end) {
			target = append(target, []int{start, end})
		}

		current = v
		start = i
		end = i
	}

	if isMatchLarge(start, end) {
		target = append(target, []int{start, end})
	}
	return target
}
