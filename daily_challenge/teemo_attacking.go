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

	lastAttack := timeSeries[0]
	for i := 1; i < len(timeSeries); i++ {
		current := timeSeries[i]
		distance := current - lastAttack
		if distance > duration {
			attack += duration
		} else {
			attack += distance
		}

		lastAttack = timeSeries[i]
	}
	return attack + duration
}
