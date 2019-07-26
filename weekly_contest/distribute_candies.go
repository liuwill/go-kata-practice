package weekly_contest

/**
 * daily-challenge-1103
 * PUZZLE: Distribute Candies to People
 */
func distributeCandies(candies int, num_people int) []int {
	sum := func(n int) int { return n * (n + 1) / 2 }
	// base := 0
	current := 0
	last := 0
	round := 0
	for current <= candies {
		last = current
		current = sum(num_people * round)
		round++
	}
	left := candies - last
	round--

	result := make([]int, num_people)

	top := sum(round) * num_people
	for i, _ := range result {
		result[i] += top - round*(num_people-1)
	}
	for i := 0; i < len(result) && left > 0; i++ {
		remain := last + 1
		if remain > left {
			remain = left
		}
		result[i] += remain

		last++
		left -= remain
	}
	return result
}
