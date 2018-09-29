package weekly_contest_102

import (
	"math"
)

func sumSubarrayMins(A []int) int {
	maxSum := int(math.Pow10(9)) + 7
	sum := 0

	for i, val := range A {
		for d := 1; d < len(A)-i+1; d++ {
			min := val
			// fmt.Printf("%v\n", A[i:i+d])
			for _, cursor := range A[i : i+d] {
				if cursor < min {
					min = cursor
				}
			}
			sum += min
		}
	}

	return sum % maxSum
}
