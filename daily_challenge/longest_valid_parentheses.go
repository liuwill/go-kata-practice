package daily_challenge

func longestValidParentheses(s string) int {
	// println(s, "===========")
	max := 0
	left := '('
	// right := ')'

	match := 0
	validPos := 0
	matchList := make([]int, len(s))
	validList := make([][]int, len(s)/2+1)
	for i, v := range s {
		// fmt.Printf("matchList=%v validList=%v", matchList, validList)
		// println(" match=", match, " i=", i, " v=", string(v))
		if v == left {
			matchList[match] = i
			match++
			continue
		}

		if match <= 0 {
			match = 0
			continue
		}

		match--
		current := i - matchList[match] + 1
		isMerge := false
		for j := 0; j < validPos; j++ {
			last := validList[j]
			if last[1]+1 == matchList[match] {
				current = i - last[0] + 1
				validList[j] = []int{last[0], i}
				validPos = j + 1
				isMerge = true
				break
			}
		}
		if !isMerge {
			validList[validPos] = []int{matchList[match], i}
			validPos++
		}

		if current > max {
			max = current
		}
	}

	return max
}
