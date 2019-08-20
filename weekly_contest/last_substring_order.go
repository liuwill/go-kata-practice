package weekly_contest

func lastSubstring(s string) string {
	big := 'a'
	positions := []int{}

	for i, v := range s {
		if v == big {
			positions = append(positions, i)
		} else if v > big {
			positions = []int{i}
		}
	}

	result := ""
	for _, pos := range positions {
		current := s[pos:]
		if current > result {
			result = current
		}
	}

	return result
}
