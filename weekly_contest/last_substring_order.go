package weekly_contest

func lastSubstring(s string) string {
	big := 'a'
	positions := []int{}
	ll := len(s)

	for i, v := range s {
		if v == big {
			positions = append(positions, i)
		} else if v > big {
			positions = []int{i}
			big = v
		}
	}

	result := ""
	for _, pos := range positions {
		current := s[pos:ll]
		if current > result {
			result = current
		}
	}

	return result
}
