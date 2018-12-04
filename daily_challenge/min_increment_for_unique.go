package daily_challenge

import (
	"sort"
)

func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	count := 0
	for i := 1; i < len(A); i++ {
		v := A[i]

		if v <= A[i-1] {
			step := A[i-1] - v + 1
			count += step
			A[i] += step
		}
	}
	return count
}
