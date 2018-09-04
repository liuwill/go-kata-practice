package weekly_contest_98

func sumSubseqWidths(A []int) int {
	sum := 0

	for i := 0; i < len(A); i++ {
		max := A[i]
		min := A[i]
		for j := i + 1; j < len(A); j++ {
			if min > A[j] {
				min = A[j]
			}
			if max < A[j] {
				max = A[j]
			}

			sum += max - min
		}
	}

	return sum
}
