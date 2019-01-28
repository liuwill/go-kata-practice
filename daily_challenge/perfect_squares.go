package daily_challenge

/**
 * daily-challenge-279
 * PUZZLE: Perfect Squares
 */
func numSquares(n int) int {
	start := n / 2
	left := n
	count := 0

	for left > 0 {
		square := start * start
		if square > left {
			start--
		} else {
			count++
			println(start)
			left -= square
		}
	}
	return count
}
