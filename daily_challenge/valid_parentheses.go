package daily_challenge

func isLeftParenthese(letter rune) bool {
	if letter == '{' || letter == '[' || letter == '(' {
		return true
	}

	return false
}

func isValid(s string) bool {
	pos := 0
	temp := make([]rune, len(s))
	parentheseMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	answer := true
	for _, v := range s {
		if isLeftParenthese(v) {
			temp[pos] = v
			pos++
		} else {
			left := parentheseMap[v]
			if pos < 1 || left != temp[pos-1] {
				answer = false
				break
			}

			pos--
		}
	}

	if pos > 0 {
		answer = false
	}
	return answer
}
