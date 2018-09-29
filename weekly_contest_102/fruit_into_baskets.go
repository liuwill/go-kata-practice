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
		second := -1
		for _, next := range tree[i+1:] {
			if next == cursor {
				current++
			} else if second != -1 && next == second {
				current++
			} else if second == -1 {
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
