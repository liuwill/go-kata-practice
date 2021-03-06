package weekly_contest

/**
 * daily-challenge-1221
 * PUZZLE: Split a String in Balanced Strings
 */
func balancedStringSplit(s string) int {
	llen := 0
	rlen := 0

	count := 0
	for _, v := range s {
		if v == 'L' {
			llen++
		} else if v == 'R' {
			rlen++
		}

		if llen == rlen {
			llen = 0
			rlen = 0
			count++
		}
	}
	return count
}
