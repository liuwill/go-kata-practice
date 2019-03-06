package weekly_contest

/**
 * weekly-contest-120
 * PUZZLE: squares-of-a-sorted-array
 */
func sortedSquares(A []int) []int {
	min := -1
	currentMin := 10001

	for i, v := range A {
		current := v
		if currentMin < current {
			continue
		}

		if v < 0 {
			current = -v
		}

		if current < currentMin {
			min = i
			currentMin = current
		}
	}

	result := make([]int, len(A))
	result[0] = A[min] * A[min]
	pos := 1

	left := 0
	right := 0

	length := len(A) - 1
	topVal := 0
	endVal := 0
	for pos <= length {
		if min-(left+1) < 0 {
			for min+right+1 < len(A) {
				right++
				endVal := A[min+right] * A[min+right]
				result[pos] = endVal
				pos++
			}
			continue
		} else {
			topVal = A[min-(left+1)] * A[min-(left+1)]
		}

		if min+(right+1) > length {
			for min-(left+1) >= 0 {
				left++
				topVal = A[min-left] * A[min-left]
				result[pos] = topVal
				pos++
			}
			continue
		} else {
			endVal = A[min+right+1] * A[min+right+1]
		}

		if topVal > endVal {
			right++
			result[pos] = endVal
			pos++
		} else {
			left++
			result[pos] = topVal
			pos++
		}
	}
	return result
}

func shortSortedSquares(A []int) []int {
	result := make([]int, len(A))
	top := 0
	end := len(A) - 1

	for i := len(A) - 1; i >= 0; i-- {
		topVal := A[top] * A[top]
		endVal := A[end] * A[end]

		if topVal > endVal {
			result[i] = topVal
			top++
		} else {
			result[i] = endVal
			end--
		}
	}

	return result
}
