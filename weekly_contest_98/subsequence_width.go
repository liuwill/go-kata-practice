package weekly_contest_98

import (
	"math"
	"sort"
)

func sumSubseqWidths(A []int) int {
	maxSum := int(math.Pow10(9)) + 7
	sum := 0
	length := len(A)
	sort.Ints(A)

	distancePows := make([]int, len(A))
	distancePows[0] = 1
	for i := 1; i < length; i++ {
		distancePows[i] = distancePows[i-1] * 2 % maxSum
	}

	for i := 0; i < length; i++ {
		sum = (sum + (distancePows[i]-distancePows[length-1-i])*A[i]) % maxSum
	}

	return sum
}
