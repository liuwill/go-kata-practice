package weekly_contest_83

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

		if end-start >= 3 {
			target = append(target, []int{start, end})
		}
		current = v
		start = i
		end = i
	}

	if end-start >= 3 {
		target = append(target, []int{start, end})
	}
	return target
}
