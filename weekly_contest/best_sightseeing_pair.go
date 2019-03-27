package weekly_contest

/**
 * daily-challenge-1021
 * PUZZLE: Best Sightseeing Pair
 */
func maxScoreSightseeingPair(A []int) int {
	score := 0
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			current := A[i] + A[j] + i - j
			if current > score {
				score = current
			}
		}
	}
	return score
}

func maxScoreSightseeingPairFast(A []int) int {
	score := 0
	pos := 0
	for i := 1; i < len(A); i++ {
		current := A[pos] + A[i] + pos - i
		if current > score {
			score = current
		}

		if A[i]+i > A[pos]+pos {
			pos = i
		}
	}
	return score
}
