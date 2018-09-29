package weekly_contest_102

import (
	"math"
)

func sumSubarrayMinsSimple(A []int) int {
	maxSum := int(math.Pow10(9)) + 7
	sum := 0

	for _, val := range A {
		sum += val
	}

	for len(A) > 1 {
		min := 30001
		pos := 0
		// fmt.Printf("++++ %v\n", A)
		for i, val := range A {
			if val < min {
				min = val
				pos = i

				sum += val * i
				// println("==>", val, i, val*i)
			} else {
				sum += min * (i - pos)
				// println("=====>", val, i, i-pos, val*(i-pos))
			}
			// println(sum, "===", i, val, pos, min)
		}

		if pos > 0 {
			sum += min
		}
		A = A[pos+1:]
	}

	return sum % maxSum
}

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
