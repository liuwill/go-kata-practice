package weekly_contest

/**
 * weekly-contest-1029
 * PUZZLE: Two City Scheduling
 */
func twoCitySchedCost(costs [][]int) int {
	target := 0
	for i := 0; i < len(costs); i++ {
		item := costs[i]
		current := item[0]
		if item[1] < item[0] {
			current = item[1]
		}

		target += current
	}
	return target
}
