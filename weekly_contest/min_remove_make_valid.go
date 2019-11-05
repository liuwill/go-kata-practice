package weekly_contest

/**
 * daily-challenge-1249
 * PUZZLE: Minimum Remove to Make Valid Parentheses
 */
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

func minRemoveToMakeValidFast(s string) string {
	ll := len(s)
	stack := make([]int, ll)
	top := -1

	for i, v := range s {
		if v == '(' {
			top++
			stack[top] = i
		} else if v == ')' {
			if top > -1 && s[stack[top]] == '(' {
				top--
			} else {
				top++
				stack[top] = i
			}
		}
	}

	target := make([]byte, 0, len(s)-(top+1))
	start, end := 0, len(s)
	sll := top + 1
	for i := 0; i < sll; i++ {
		end = stack[i]
		target = append(target, ([]byte(s[start:end]))...)

		start = end + 1
		end = len(s)
	}

	target = append(target, ([]byte(s[start:end]))...)
	return string(target)
}
