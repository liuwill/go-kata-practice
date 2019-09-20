package weekly_contest

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	total := 0
	min := 0

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
