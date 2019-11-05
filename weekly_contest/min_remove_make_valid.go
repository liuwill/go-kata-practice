package weekly_contest

func minRemoveToMakeValid(s string) string {
	ll := len(s)
	targetMap := make([]int, ll)
	stack := make([]int, ll)
	top := -1

	tll := 0
	for i, v := range s {
		if v != '(' && v != ')' {
			targetMap[i] = 1
			tll++
			continue
		}

		if v == '(' {
			top++
			stack[top] = i
		} else if v == ')' {
			if top < 0 {
				continue
			}

			tt := stack[top]
			targetMap[i] = 1
			targetMap[tt] = 1
			tll += 2
			top--
		}
	}

	runeStr := make([]rune, tll)
	rll := 0
	for i, v := range s {
		if targetMap[i] == 1 {
			runeStr[rll] = v
			rll++
		}
	}

	return string(runeStr)
}
