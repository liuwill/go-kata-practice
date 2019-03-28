package weekly_contest

func shipWithinDays(weights []int, D int) int {
	target := 0
	max := 0
	total := 0
	for _, val := range weights {
		if val > max {
			max = val
		}
		total += val
	}

	for target = max; target < len(weights)*D; target++ {
		w := target
		count := 1

		for i := 0; i < len(weights); i++ {
			current := w - weights[i]
			if current <= 0 {
				w = target
				if current < 0 {
					i--
				}
				count++
			} else {
				w -= weights[i]
			}
		}

		if count <= D {
			break
		}
	}
	return target
}
