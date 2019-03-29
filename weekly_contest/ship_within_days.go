package weekly_contest

/**
 * daily-challenge-1014
 * PUZZLE: Capacity To Ship Packages Within D Days
 */
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

	for target = max; target < total; target++ {
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
