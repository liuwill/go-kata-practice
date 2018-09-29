package weekly_contest_102

import (
	"math"
	"sort"
)

func sumSubarrayMins(A []int) int {
	maxSum := int(math.Pow10(9)) + 7
	sum := 0
	sort.Ints(A)

	pow := 0
	for i := len(A) - 1; i >= 0; i-- {
		val := A[i]

		times := int(math.Pow(2, float64(pow))) - 1
		sum += val * (times + 1) % maxSum
		println(val, times+1)
		pow++
	}

	return sum % maxSum
}
