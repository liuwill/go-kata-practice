package weekly_contest

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

	i := 1
	result := 0
	for copy > 0 {
		bit := copy % 2
		// source = string(bit) + source
		// target += string(bit ^ 1)
		reverse := bit ^ 1
		result += int(reverse) * i

		copy = copy / 2
		i = i * 2
	}

	return result
}
