package daily_challenge

func longestValidParentheses(s string) int {
	max := 0
	left := '('
	right := ')'

	current := 0
	match := 0

	for _, v := range s {
		if v == left {
			match++
		}

		if v == right {
			if match > 0 {
				match--
				current += 2
			} else {
				current = 0
				match = 0
			}
		}

		if current > max {
			max = current
		}
	}
	return max
}
