package daily_challenge

func hasAlternatingBits(n int) bool {
	result := true
	if n < 2 {
		return result
	}

	bit := n % 2
	n = n / 2

	current := bit
	for n > 0 {
		bit := n % 2
		if bit == current {
			result = false
			break
		}

		current = bit
	}

	return result
}
