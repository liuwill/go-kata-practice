package weekly_contest

/**
 * daily-challenge-1014
 * PUZZLE: Capacity To Ship Packages Within D Days
 */
func shipWithinDays(weights []int, D int) int {
	target := 0
	min := 0
	max := 0
	for _, val := range weights {
		if val > min {
			min = val
		}
		max += val
	}

	for target = min; target < max; target++ {
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

func shipWithinDaysBinary(weights []int, D int) int {
	min := 0
	max := 0
	for _, val := range weights {
		if val > min {
			min = val
		}
		max += val
	}
	if max/len(weights) > min {
		min = max / len(weights)
	}

	for min < max {
		mid := (min + max) / 2
		days := 1
		weight := 0

		for _, val := range weights {
			if val+weight > mid {
				weight = 0
				days++
			}
			weight += val
		}

		if days <= D {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return min
}
