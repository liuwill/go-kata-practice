package weekly_contest

/**
 * daily-challenge-1103
 * PUZZLE: Distribute Candies to People
 */
func distributeCandies(candies int, num_people int) []int {
	result := make([]int, num_people)

	pos := 0
	for start := 1; candies > 0; start++ {
		current := start
		if current > candies {
			current = candies
		}
		result[pos] += current

		pos = (pos + 1) % num_people
		candies -= current
	}

	return result
}

/*
func distributeCandies(candies int, num_people int) []int {
	sum := func(n int) int { return n * (n + 1) / 2 }
	current := 0
	last := 0
	round := 1
	for current <= candies {
		last = current
		current = sum(num_people * round)
		round++
	}
	left := candies - last
	round -= 2

	result := make([]int, num_people)

	top := sum(round) * num_people
	for i, _ := range result {
		result[i] += top - round*(num_people-i-1)
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
*/
