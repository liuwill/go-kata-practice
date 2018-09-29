package weekly_contest_102

import (
	"math"
)

const MAX_SUBARRAY_LENGTH = 30000

func sumSubarrayMinsFast(A []int) int {
	maxSum := int(math.Pow10(9)) + 7
	stack := make([][]int, len(A))
	top := 0

	sum := 0
	for _, val := range A {
		sum += val
	}

	stack[top] = []int{0, len(A) - 1, len(A)}
	top += 1
	for top > 0 {
		top--

		start := stack[top][0]
		end := stack[top][1]
		// fmt.Printf(">> %v %v\n", stack[top], list)
		min := MAX_SUBARRAY_LENGTH
		pos := 0

		for i := start; i <= end; i++ {
			val := A[i]
			if val < min {
				min = val
				pos = i
			}
		}

		// println("min", min, pos, start, end)
		front := pos - start
		backend := end - pos
		sum += min * (front + backend + front*backend)

		if front > 1 {
			stack[top] = []int{start, pos - 1, front}
			// println("front", start, pos-1, front)
			top++
		}

		if backend > 1 {
			stack[top] = []int{pos + 1, end, backend}
			// println("back", pos+1, end, backend)
			top++
		}
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
