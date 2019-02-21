package daily_challenge

func findPoisonedDuration(timeSeries []int, duration int) int {
	lastAttack := timeSeries[0]

	attack := 0
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
