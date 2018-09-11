package daily_challenge

func scoreOfParentheses(S string) int {
	left := '('

	count := 0
	result := 0
	for i, v := range S {
		if v == left {
			count++
			continue
		}

		if int(S[i-1]) == int(left) {
			result += 1 << uint16(count-1)
		}

		count--
	}
	return result
}
