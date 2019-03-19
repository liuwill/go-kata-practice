package weekly_contest

import "math"

/**
 * daily-challenge-1012
 * PUZZLE: Complement of Base 10 Integer
 */
func bitwiseComplement(N int) int {
	// target := ""
	// source := ""

	copy := N
	if copy == 0 {
		return 1
	}

	i := 0
	result := 0
	for copy > 0 {
		bit := copy % 2
		// source = string(bit) + source
		// target += string(bit ^ 1)
		reverse := bit ^ 1
		result += int(reverse) * int(math.Pow(float64(2), float64(i)))

		copy = copy / 2
		i++
	}

	return result
}
