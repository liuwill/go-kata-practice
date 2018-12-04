package daily_challenge

import (
	"sort"
)

func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	count := 0
	for i, v := range A {
		if i == 0 {
			continue
		}

		if v <= A[i-1] {
			count += A[i-1] - v + 1
			A[i] += A[i-1] - v + 1
		}
	}
	return count
}
