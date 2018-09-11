package weekly_contest_100

func isMonotonic(A []int) bool {
	increasing := true
	decreasing := true
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			decreasing = false
		}

		if A[i] < A[i-1] {
			increasing = false
		}

		if !increasing && !decreasing {
			break
		}
	}
	return increasing || decreasing
}
