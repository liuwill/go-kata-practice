package weekly_contest

import "math"

/**
 * daily-challenge-1012
 * PUZZLE: Complement of Base 10 Integer
 */
func bitwiseComplement(N int) int {
	target := ""
	source := ""

	copy := N
	for copy > 0 {
		bit := copy % 2
		source = string(bit) + source
		target += string(bit ^ 1)

		copy = copy / 2
	}

	result := 0
	for i, bit := range target {
		result += int(bit) * int(math.Pow(float64(2), float64(i)))
	}

	return result
}
