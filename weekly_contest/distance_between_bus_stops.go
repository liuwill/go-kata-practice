package weekly_contest

/**
 * daily-challenge-1184
 * PUZZLE: Distance Between Bus Stops
 */
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	total := 0
	min := 0

	big := destination
	if start > big {
		destination = start
		start = big
	}

	for i, v := range distance {
		total += v
		if i >= start && i < destination {
			min += v
		}
	}

	guess := total - min
	if min > guess {
		min = guess
	}

	return min
}
