package weekly_contest_102

func sortArrayByParity(A []int) []int {
	i, j := 0, len(A)-1

	for i < j {
		for A[i]%2 == 0 {
			i++
		}

		for A[j]%2 > 0 && j > i {
			j--
		}

		middle := A[i]
		A[i] = A[j]
		A[j] = middle
	}
	return A
}
