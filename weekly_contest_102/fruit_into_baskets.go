package weekly_contest_102

func totalFruit(tree []int) int {
	// fmt.Printf("%v\n", tree)
	max := 0
	length := len(tree)

	for i := 0; i < length; {
		cursor := tree[i]
		left := length - i
		if left <= max {
			break
		}
		// fmt.Printf("%v ", i)

		current := 1
		step := 1
		diff := false
		second := -1
		for _, next := range tree[i+1:] {
			if next == cursor {
				current++
				if !diff {
					step++
				}
			} else if second != -1 && next == second {
				current++
				diff = true
			} else if second == -1 {
				current++
				second = next
				diff = true
			} else {
				break
			}
		}

		if max < current {
			max = current
		}

		i += step
	}
	// println("End")
	return max
}
