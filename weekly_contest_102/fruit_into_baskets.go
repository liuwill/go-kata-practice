package weekly_contest_102

func totalFruit(tree []int) int {
	max := 0
	length := len(tree)

	for i, cursor := range tree {
		left := length - i
		if left < max {
			break
		}

		current := 1
		second := 0
		for _, next := range tree[i+1:] {
			if next == cursor {
				current++
			} else if second != 0 && next == second {
				current++
			} else if second == 0 {
				current++
				second = next
			} else {
				break
			}
		}

		if max < current {
			max = current
		}
	}
	return max
}
