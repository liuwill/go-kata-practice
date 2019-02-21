package daily_challenge

/**
 * daily-challenge-495
 * PUZZLE: Teemo Attacking
 */
func findPoisonedDuration(timeSeries []int, duration int) int {
	attack := 0
	if len(timeSeries) < 1 {
		return attack
	}

	for i := 1; i < len(timeSeries); i++ {
		distance := timeSeries[i] - timeSeries[i-1]
		if distance > duration {
			attack += duration
		} else {
			attack += distance
		}
	}
	return attack + duration
}
