package weekly_contest

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
