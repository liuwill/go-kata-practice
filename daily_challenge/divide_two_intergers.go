package daily_challenge

import "math"

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

	for dividend >= divisor {
		dividend -= divisor

		if mark {
			result--
		} else {
			result++
		}

		if result >= max || result <= min {
			break
		}
	}

	return result
}
