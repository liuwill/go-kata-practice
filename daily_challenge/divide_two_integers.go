package daily_challenge

import "math"

/**
 * daily-challenge-29
 * PUZZLE: Divide Two Integers
 */
func divide(dividend int, divisor int) int {
	result := 0
	max := int(math.Pow(2, 31) - 1)
	min := int(math.Pow(2, 31)) * -1

	mark := false
	if dividend < 0 {
		mark = !mark
		dividend *= -1
	}
	if divisor < 0 {
		mark = !mark
		divisor *= -1
	}

	step := 1
	currentDiv := divisor
	for dividend >= currentDiv {
		result += step
		dividend -= currentDiv
		currentDiv += currentDiv
		step += step

		if dividend < currentDiv {
			currentDiv = divisor
			step = 1
		}
	}

	if mark {
		result = -result
	}

	if result > max {
		return max
	} else if result < min {
		return min
	}
	return result
}
