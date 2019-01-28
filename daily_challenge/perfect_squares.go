package daily_challenge

/**
 * daily-challenge-279
 * PUZZLE: Perfect Squares
 */
func numSquares(n int) int {
	min := n

	for index := n / 2; index > 0; index-- {
		left := n
		count := 0
		start := index

		for left > 0 {
			square := start * start
			if square > left {
				start--
			} else {
				count++
				left -= square
			}
		}

		if min > count {
			min = count
		}
	}

	return min
}
