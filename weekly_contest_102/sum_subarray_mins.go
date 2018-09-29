package weekly_contest_102

import (
	"math"
	"sort"
)

func sumSubarrayMins(A []int) int {
	maxSum := int(math.Pow10(9)) + 7
	sum := 0
	sort.Ints(A)

	count := 0
	times := 0
	for i := len(A) - 1; i >= 0; i-- {
		val := A[i]

		times += int(math.Pow(2, float64(count)))
		sum += val * (times) % maxSum
		println(val, i)
		count++
	}

	return sum % maxSum
}
